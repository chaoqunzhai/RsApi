package apis

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/jobs/watch"
	"go-admin/cmd/migrate/migration/models"
	"go-admin/common/prometheus"
	"go-admin/common/qiniu"
	"go-admin/common/utils"
	"go-admin/config"
	"go-admin/costAlg"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Crontab struct {
	api.Api
}

// 写入数据到文件
func writeHostsToFile(filename string, hosts []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(strings.Join(hosts,","))
	return writer.Flush()
}

// 读取文件内容
func readHostsFromFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	hosts := strings.Split(string(content), ",")
	// 移除空字符串
	var result []string
	for _, host := range hosts {
		if host != "" {
			result = append(result, strings.TrimSpace(host))
		}
	}
	return result, nil
}

// 记录不同的数据（这里简单打印，实际应用中可以写入日志或数据库）
func recordDifferentHosts(currentHosts, yesterdayHosts []string) []string {
	currentSet := make(map[string]struct{})
	yesterdaySet := make(map[string]struct{})

	for _, host := range currentHosts {
		currentSet[host] = struct{}{}
	}
	for _, host := range yesterdayHosts {
		yesterdaySet[host] = struct{}{}
	}

	differentHosts :=make([]string,0)
	for host := range currentSet {
		if _, found := yesterdaySet[host]; !found {
			differentHosts = append(differentHosts, host)
		}
	}
	for host := range yesterdaySet {
		if _, found := currentSet[host]; !found {
			differentHosts = append(differentHosts, host)
		}
	}

	if len(differentHosts) > 0 {
		fmt.Println("Different hosts today compared to yesterday:", differentHosts)
	} else {
		fmt.Println("No different hosts today compared to yesterday.")
	}

	return differentHosts
}

func (e Crontab) ComputeMonth(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	currentTime := time.Now()


	var hostIds []int64
	e.Orm.Model(&models.Host{}).Select("id").Scan(&hostIds)


	backtrackStr:=c.Query("backtrack")
	backtrackInt :=1
	if backtrackStr != ""{
		backtrackInt,_=strconv.Atoi(backtrackStr)
	}

	for _,row:=range hostIds{

		for i := 0; i < backtrackInt; i++ {
			month := currentTime.AddDate(0, 0, -i).Format("2006-01")
			var hostRow []models.HostIncome
			e.Orm.Model(&models.HostIncome{}).Where("host_id = ? and alg_day like ? and record_m = 0",row,month+"%").Find(&hostRow)

			monthIncome :=0.0
			monthCost :=0.0
			for _,incomeRow:=range hostRow{
				monthCost +=incomeRow.DayCost
				monthIncome+=incomeRow.Income

				e.Orm.Model(&models.HostIncome{}).Where("id = ?",incomeRow.Id).Updates(map[string]interface{}{
					"record_m":1,
				})
			}
			var HostIncomeMonth models.HostIncomeMonth
			e.Orm.Model(&models.HostIncomeMonth{}).Where("host_id = ? and month = ?",row,month).Limit(1).Find(&HostIncomeMonth)
			HostIncomeMonth.HostId = row
			HostIncomeMonth.Month = month
			HostIncomeMonth.Income +=utils.RoundDecimal(monthIncome)
			HostIncomeMonth.Cost +=utils.RoundDecimal(monthCost)
			if HostIncomeMonth.Id == 0 {
				e.Orm.Create(&HostIncomeMonth)
			}else {
				e.Orm.Save(&HostIncomeMonth)
			}

		}
	}
	e.OK("","记录成功")
	return
}
func (e Crontab) DataBurning(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//保存今天的主机统计
	onlineHealthySql := "healthy_at >= DATE_SUB(NOW(), INTERVAL 30 MINUTE)"
	offlineHealthySql := "healthy_at <= DATE_SUB(NOW(), INTERVAL 30 MINUTE) OR healthy_at IS NULL"
	//在线-----
	var (
		onlineCount    int64
		offlineCount int64
		totalBandwidth int64
		waitCount int64
		todoCount int64
	)
	onlineHostIds:=make([]string,0)
	e.Orm.Model(&models.Host{}).Select("IFNULL(SUM(balance), 0) as totalBandwidth").Where(onlineHealthySql).Scan(&totalBandwidth)
	e.Orm.Model(&models.Host{}).Select("id").Where(onlineHealthySql).Find(&onlineHostIds).Count(&onlineCount)
	notStatus := []int{3, 4}
	e.Orm.Model(&models.Host{}).Where("status not in ?", notStatus).Where(offlineHealthySql).Count(&offlineCount)

	e.Orm.Model(&models.Host{}).Where("status = 3").Count(&waitCount)
	e.Orm.Model(&models.Host{}).Where("status = 4").Count(&todoCount)


	modelDat:=models.DataBurningHost{
		Online:onlineCount,
		Offline: offlineCount,
		TotalBandwidth: utils.RoundDecimal(totalBandwidth / 1000),
		Wait: waitCount,
		Todo: todoCount,
	}

	//判断当前/tmp/data_burning_host.txt 文件是否存在, 如果不存在 创建 写入当前在线的主机ID,  如果存在 读取文件内容,和当前在线的主机做比对

	filename := "/tmp/data_burning_host.txt"


	// 检查文件是否存在
	if _, err1 := os.Stat(filename); os.IsNotExist(err1) {
		// 文件不存在，写入当前在线的主机ID
		err1 := writeHostsToFile(filename, onlineHostIds)
		if err1 != nil {
			fmt.Println("Error writing current hosts to file:", err1)
			return
		}
		fmt.Println("File created and current hosts written.")
	} else {
		// 文件存在，读取文件内容
		yesterdayHosts, err1 := readHostsFromFile(filename)
		if err1 != nil {
			fmt.Println("Error reading yesterday's hosts from file:", err1)
			return
		}
		// 比对今天和昨天的主机ID
		differentHosts :=recordDifferentHosts(onlineHostIds, yesterdayHosts)
		modelDat.DiffHost = strings.Join(differentHosts,",")
		// 更新文件内容为当前在线的主机ID（这里模拟每天更新）
		err = writeHostsToFile(filename, onlineHostIds)
		if err != nil {
			fmt.Println("Error updating file with current hosts:", err)
			return
		}
		fmt.Println("File updated with current hosts.")
	}

	e.Orm.Create(&modelDat)
	e.OK("","记录成功")
	return
}


func (e Crontab) QiNiuAmount(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//读取prom 中 node_uname_info{business=~"jinshan"} 的数据
	query := fmt.Sprintf("node_uname_info{business=~\"jinshan\"}")

	queryUrl, err := url.Parse(func() string {
		vv, _ := url.JoinPath(config.ExtConfig.Prometheus.Endpoint, "/api/v1/query")
		return vv
	}())

	parameters := url.Values{}
	parameters.Add("time", fmt.Sprintf("%v",time.Now().Unix()))

	parameters.Add("query", query)
	queryUrl.RawQuery = parameters.Encode()

	ProResult, err := prometheus.GetPromNodeInfoResult(queryUrl)

	if err != nil {

		e.Error(-1,errors.New("数据不存在"),"")
		return
	}
	if len(ProResult.Data.Result) == 0 {
		e.Error(-1,errors.New("数据不存在"),"")
		return
	}

	algDay:=time.Now().AddDate(0,0,-1).Format(time.DateOnly)
	qiNiuUrl:="/supplier/outer/sirius/supply/external/out_api/getDataAndReturn"
	for _,row:=range ProResult.Data.Result {
		//主机名
		hostName :=row.Metric.Instance
		deviceId :=row.Metric.Sn

		//hostName 是CMDB里面的数据
		var host models.Host
		e.Orm.Model(&models.Host{}).Select("id").Where("host_name = ?",hostName).Limit(1).Find(&host)
		if host.Id == 0 {continue}

		//拿deviceId 去 七牛的API里面查询到费用

		params :=map[string]interface{}{
			"device_id":deviceId,
			"srm_channel":"1000122751",
			"start_date":algDay,
			"end_date":algDay,
		}
		dat,getErr :=qiniu.GetQueryQiNiu(qiNiuUrl,params)
		if getErr!=nil{

			fmt.Println("getErr",getErr.Error())
			continue
		}

		for _, miniRow :=range dat.Data{
			var hostIncome models.HostIncome
			e.Orm.Model(&models.HostIncome{}).Select("id").Where("host_id = ? and alg_day = ?",host.Id, miniRow.Date).Limit(1).Find(&hostIncome)
			if hostIncome.Id == 0 {continue}
			e.Orm.Model(&models.HostIncome{}).Where("id = ?",hostIncome.Id).Updates(map[string]interface{}{
				"settle_bandwidth":miniRow.ChargeDay95,
			})

		}


	}
	e.OK("","successful")
	return

}
func (e Crontab) OpenApiAmount(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	parentDayStr := c.Query("day")
	parentDay :=0
	if parentDayStr != ""{
		parentDay,_ = strconv.Atoi(parentDayStr)
	}

	//开始采集第三方openApi结算的收益， 都是晚上执行 然后白天的数据
	costAlgorithm := costAlg.OpenApiLinWu{}
	costAlgorithm.SetupDb(sdk.Runtime.GetDb())
	costAlgorithm.LoopData(parentDay)

	e.OK("", "successful")
	return
}

func (e Crontab) WatchOnlineUsage(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	startTime := time.Now()
	watch.WatchOnlineUsage()
	result := map[string]interface{}{
		"runTime": time.Now().Sub(startTime).Seconds(),
	}
	e.OK(result, "successful")
	return
}
func (e Crontab) Algorithm(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	backtrackStr:=c.Query("backtrack")
	backtrackInt :=1
	if backtrackStr != ""{
		backtrackInt,_=strconv.Atoi(backtrackStr)
	}


	startTime := time.Now()
	costAlgorithm := costAlg.CostAlgorithm{
		BacktrackInt:backtrackInt,
	}
	costAlgorithm.SetupDb(sdk.Runtime.GetDb())
	costAlgorithm.StartHostCompute()

	result := map[string]interface{}{
		"runTime": time.Now().Sub(startTime).Seconds(),
	}
	e.OK(result, "successful")
	return
}

