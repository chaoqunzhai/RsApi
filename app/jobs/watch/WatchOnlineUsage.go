package watch

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/cmd/migrate/migration/models"
	"go-admin/common/prometheus"
	"go-admin/common/utils"
	"go-admin/config"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type MonitorCompute struct {
	Empty        bool    `json:"empty"`
	Max          float64 `json:"max"`
	Min          float64 `json:"min"`
	Avg          float64 `json:"avg"`
	PercentBytes float64 `json:"percentBytes"` //今天计算的95带宽(bytes)
	PercentValue float64 `json:"percentG"`     //日95带宽G 今天计算的日95带宽(G)
	IspDayPrice  float64 `json:"ispPrice"`     //计算今天 运营商的收益
	Usage        float64 `json:"usage"`        //利用率
}

func WatchOnlineUsage() {
	dbList := sdk.Runtime.GetDb()
	now := time.Now()
	// 计算昨天的日期
	yesterday := now.AddDate(0, 0, -1)

	// 计算昨天的 0 点
	startOfYesterday := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())

	// 计算昨天的 23:59:59
	endOfYesterday := startOfYesterday.Add(24*time.Hour - time.Second)

	for _, d := range dbList {
		var hostList []models.Host
		d.Model(&models.Host{}).Select("host_name,balance,id").Where("healthy_at >= DATE_SUB(NOW(), INTERVAL 30 MINUTE)").Find(&hostList)

		tenMinutesAgo := fmt.Sprintf("%v", startOfYesterday.Unix())
		endTime := fmt.Sprintf("%v", endOfYesterday.Unix())
		for _, host := range hostList {

			transmitQuery := fmt.Sprintf("sum(rate(phy_nic_network_transmit_bytes_total{instance=\"%v\"}[5m]))*8", host.HostName)

			BandwidthIncome := host.Balance / 1000 //MB换算成G

			compute := RequestPromResult(tenMinutesAgo, endTime, "60", transmitQuery, false)
			updateMap := make(map[string]interface{})
			updateTag := false
			if !compute.Empty { //有数据
				updateTag = true
				updateMap["usage"] = utils.RoundDecimalFlot64(3, compute.PercentValue/BandwidthIncome)
				updateMap["percent_value"] = compute.PercentValue

			} else { //空的监控数据,那就尝试在 另外一个指标中请求

				computeTwo := WatchFlowDownloadBandwidth(tenMinutesAgo, endTime, host)
				fmt.Printf("在其他的参数中请求 %v,%+v,%v\n", host.HostName, computeTwo, BandwidthIncome)
				if !computeTwo.Empty { //有数据
					updateTag = true
					updateMap["usage"] = utils.RoundDecimalFlot64(3, computeTwo.PercentValue/BandwidthIncome)
					updateMap["percent_value"] = computeTwo.PercentValue
				}
			}

			if updateTag {
				d.Model(&models.Host{}).Where("id = ?", host.Id).Updates(updateMap)
			}

		}
	}
}

// 点心的接口定义是 flow_download_bandwidth_by_minute
func WatchFlowDownloadBandwidth(tenMinutesAgo, endTime string, host models.Host) *MonitorCompute {
	//这里面的单位是MB
	transmitQuery := fmt.Sprintf("sum(flow_bandwidth_by_minute{instance=\"%v\"})", host.HostName)

	compute := RequestPromResult(tenMinutesAgo, endTime, "60", transmitQuery, true)

	return compute
}

func RequestPromResult(start, end, setup, query string, isMb bool) *MonitorCompute {

	result := &MonitorCompute{}
	//查询普罗米修斯数据
	queryUrl, err := url.Parse(func() string {
		vv, _ := url.JoinPath(config.ExtConfig.Prometheus.Endpoint, "/api/v1/query_range")
		return vv
	}())

	parameters := url.Values{}
	parameters.Add("start", start)
	parameters.Add("end", end)
	parameters.Add("step", setup)
	parameters.Add("query", query)
	queryUrl.RawQuery = parameters.Encode()

	ProResult, err := prometheus.GetPromResult(queryUrl)

	if err != nil {
		return result
	}
	XData := make([][]interface{}, 0)
	XValue := make([]float64, 0)

	if len(ProResult.Data.Result) > 0 {

		parentUnix := 0.0 //上一个时间戳

		setupNumber := 0.0 //时间戳的区间
		for _, row := range ProResult.Data.Result[0].Value {

			if len(row) != 2 {
				continue
			}
			unixFloat := row[0].(float64)
			if setupNumber == 0 { //只需要计算一次区间即可
				setupNumber = unixFloat - parentUnix //计算区间
			}
			if parentUnix > 0 { //如果有上一个值
				//时间戳的差大于区间,那就少点了,补多少点,那就算一个次数循环写入
				ac := unixFloat - parentUnix
				if ac > setupNumber {
					for i := 0; i < int(ac/setupNumber); i++ {
						addUnix := unixFloat + setupNumber*float64(i+1)

						addUnixTime := time.Unix(int64(addUnix), 0)
						//进行计算
						XData = append(XData, []interface{}{addUnixTime, 0})
						XValue = append(XValue, 0)

					}
				}

			}
			parentUnix = unixFloat

			valueStr := row[1].(string)
			valueFloat, _ := strconv.ParseFloat(valueStr, 64)
			unixTime := time.Unix(int64(unixFloat), 0)
			XData = append(XData, []interface{}{unixTime, valueStr})

			//进行计算
			XValue = append(XValue, valueFloat)

		}
	}
	sort.Float64s(XValue)

	//计算95值
	if len(XValue) > 1 {

		Percent := utils.Percentile(XValue, 0.95)
		//fmt.Println("Percent", Percent)
		result.PercentBytes = Percent
		if isMb { //是MB的单位
			result.PercentValue = utils.RoundDecimalFlot64(3, Percent/1000)
		} else {
			result.PercentValue = utils.RoundDecimalFlot64(3, Percent/(1024*1024*1024))
		}

	} else {
		result.Empty = true
	}

	return result
}
