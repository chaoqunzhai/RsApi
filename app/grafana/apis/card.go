package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/cmdb/service/dto"
	models2 "go-admin/cmd/migrate/migration/models"
	cDto "go-admin/common/dto"
	"go-admin/common/prometheus"
	"go-admin/common/utils"
	"math"
	"sort"
	"strconv"
	"sync"
)

type Card struct {
	api.Api
}

type HostBindBu struct {
	HostId     int `json:"host_id"`
	BusinessId int `json:"business_id"`
}
type BusinessInfo struct {
	Id      int     `json:"id"`
	Online  int     `json:"online"`
	Offline int     `json:"offline"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
type HostInfo struct {
	HostId  int     `json:"host_id"`
	Status  int     `json:"status"`
	Balance float64 `json:"balance"`
}

type MonitorQuery struct {
	BusinessId      string `form:"businessId"`
	Start           string `form:"start"`
	End             string `form:"end"`
	Setup           int    `form:"setup"`
	cDto.Pagination `search:"-"`
}

var (
	IspList = []string{"电信", "联通", "移动"}
)

func (e Card) CardList(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var buList []models2.Business
	e.Orm.Model(&models2.Business{}).Find(&buList)

	buMap := make(map[int]BusinessInfo, 0)
	for _, bu := range buList {

		buMap[bu.Id] = BusinessInfo{
			Id:      bu.Id,
			Online:  0,
			Offline: 0,
			Name:    bu.Name,
		}
	}
	//result := make([]interface{}, 0)

	HostBindBuList := make([]HostBindBu, 0)
	e.Orm.Raw("select * from host_bind_business").Scan(&HostBindBuList)

	hostIds := make([]int, 0)
	hostBindBu := make(map[int]int, 0)
	for _, row := range HostBindBuList {
		hostIds = append(hostIds, row.HostId)
		hostBindBu[row.HostId] = row.BusinessId
	}

	hostList := make([]models2.Host, 0)
	e.Orm.Model(&models2.Host{}).Select("id,status,balance").Find(&hostList)
	hostStatusMap := make(map[int]HostInfo, 0)
	for _, host := range hostList {
		hostStatusMap[host.Id] = HostInfo{
			Balance: host.Balance,
			Status:  host.Status,
		}
	}
	for _, row := range HostBindBuList {

		buInfo, ok := buMap[row.BusinessId]
		if ok {

			if hostInfo, hostStatusOk := hostStatusMap[row.HostId]; hostStatusOk {
				switch hostInfo.Status {
				case -1:
					buInfo.Offline += 1
				case 1:
					buInfo.Online += 1

				}
				buInfo.Balance += hostInfo.Balance / 1000
			}
		}

		buInfo.Balance = utils.RoundDecimal(buInfo.Balance)

		buMap[row.BusinessId] = buInfo
	}
	e.OK(buMap, "")
	return
}

func (e Card) Monitor(c *gin.Context) {
	req := MonitorQuery{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//如果有业务ID 那就查询 业务关联的主机
	HostBindBuList := make([]int, 0)
	orm := e.Orm.Model(&models2.Host{})

	queryMap := make(map[int]int, 0)
	if req.BusinessId != "" {
		var buModel models2.Business
		e.Orm.Model(&models2.Business{}).Where("id = ?", req.BusinessId).Limit(1).Find(&buModel)

		e.Orm.Raw(fmt.Sprintf("select host_id from host_bind_business where business_id = %v", req.BusinessId)).Scan(&HostBindBuList)

		if len(HostBindBuList) > 0 {

			orm = orm.Where("id in ?", HostBindBuList)

			for _, hostId := range HostBindBuList {

				switch buModel.Name {
				case "点心":
					queryMap[hostId] = 1
				default:
					queryMap[hostId] = 0
				}

			}
		}
	}
	hostList := make([]models2.Host, 0)

	var count int64
	orm.Select("id,host_name,remark").Where("status = 1").Scopes(
		cDto.Paginate(req.GetPageSize(), req.GetPageIndex()),
	).Find(&hostList).Order(fmt.Sprintf("`usage` asc")).Limit(-1).Offset(-1).Count(&count)

	data := make([]interface{}, 0)

	promReq := dto.RsHostMonitorFlow{
		Start: req.Start,
		End:   req.End,
		Setup: req.Setup,
	}
	for _, row := range hostList {

		var monitorData interface{}

		queryNumber := queryMap[row.Id]
		//fmt.Println("Row!!", row.HostName, queryNumber)
		switch queryNumber {
		case 1:
			monitorData = prometheus.DianXinTransmit(row.HostName, &promReq)
		default:
			monitorData = prometheus.Transmit(row.HostName, &promReq)

		}
		hostRow := map[string]interface{}{
			"host":        row.HostName,
			"remark":      row.Remark,
			"monitorData": monitorData,
		}
		data = append(data, hostRow)

	}

	response := map[string]interface{}{
		"total": count,
		"data":  data,
	}
	e.OK(response, "")
	return
}

func (e Card) IspMonitor(c *gin.Context) {
	req := MonitorQuery{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var response utils.ByMinToMaxMap

	var buName string
	if req.BusinessId != "" {
		var buModel models2.Business
		e.Orm.Model(&models2.Business{}).Where("id = ?", req.BusinessId).Limit(1).Find(&buModel)

		buName = buModel.EnName
	}
	//开启三个协程请求

	wg := sync.WaitGroup{}

	promReq := dto.RsHostMonitorFlow{
		Start: req.Start,
		End:   req.End,
		Setup: req.Setup,
	}
	dataChannel := make(chan utils.Data, 6)
	for index, ispName := range IspList {
		wg.Add(1)
		var QueryReceive string
		var QueryTransmit string
		if buName != "" {
			QueryTransmit = fmt.Sprintf("sum(rate(phy_nic_network_transmit_bytes_total{business=~\".*%v.*\", remark=~\".*%v.*\"}[5m]))*8", buName, ispName)
			QueryReceive = fmt.Sprintf("sum(rate(phy_nic_network_receive_bytes_total{business=~\".*%v.*\", remark=~\".*%v.*\"}[5m]))*8", buName, ispName)

		} else {
			QueryTransmit = fmt.Sprintf("sum(rate(phy_nic_network_transmit_bytes_total{ remark=~\".*%v.*\"}[5m]))*8", ispName)
			QueryReceive = fmt.Sprintf("sum(rate(phy_nic_network_receive_bytes_total{remark=~\".*%v.*\"}[5m]))*8", ispName)
		}

		//fmt.Println("请求！！！！！ api", ispName, promReq)
		go func(index int, ispName, QueryReceive, QueryTransmit string) {

			defer wg.Done()

			transmitData := prometheus.RequestPromResult("上行", QueryTransmit, &promReq, false)
			receiveData := prometheus.RequestPromResult("下行", QueryReceive, &promReq, false)

			dataCa := map[string]interface{}{
				"transmit": transmitData,
				"receive":  receiveData,
			}
			dataChannel <- utils.Data{
				Name:  ispName,
				Data:  dataCa,
				Index: float64(index),
			}

		}(index, ispName, QueryReceive, QueryTransmit)
	}
	wg.Wait()
	close(dataChannel)

	for data := range dataChannel {

		response = append(response, data)
	}
	sort.Sort(response)
	e.OK(response, "")
	return

}

func (e Card) PlanBandWidth(c *gin.Context) {
	req := MonitorQuery{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var buName string
	if req.BusinessId != "" {
		var buModel models2.Business
		e.Orm.Model(&models2.Business{}).Where("id = ?", req.BusinessId).Limit(1).Find(&buModel)

		buName = buModel.EnName
	}
	var response utils.ByMaxToMinMap
	promReq := dto.RsHostMonitorFlow{
		Start: req.Start,
		End:   req.End,
		Setup: req.Setup,
	}
	dataChannel := make(chan utils.Data, 3)
	//查询三网规划带宽
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, ispName := range IspList {
			var Query string
			if buName != "" {
				Query = fmt.Sprintf("sum(plan_bandwidth{remark=~\".*%v.*\", business=~\".*%v.*\"})", ispName, buName)
			} else {
				Query = fmt.Sprintf("sum(plan_bandwidth{remark=~\".*%v.*\"})", ispName)
			}

			resultData := prometheus.RequestQueryPromResult("三网规划带宽", Query, &promReq, false)

			floatNum, err2 := strconv.ParseFloat(fmt.Sprintf("%v", resultData), 64)
			fmt.Println("获取到的数据", floatNum, resultData, err2)
			dataChannel <- utils.Data{
				Name:  ispName,
				Index: utils.RoundDecimalFlot64(2, math.Round(floatNum)/1000),
			}
		}

		defer wg.Done()
	}()
	wg.Wait()
	close(dataChannel)

	for data := range dataChannel {

		response = append(response, data)
	}
	sort.Sort(response)
	e.OK(response, "")
	return

}
