package watch

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/cmd/migrate/migration/models"
	"go-admin/common/prometheus"
	"go-admin/common/utils"
	"go-admin/config"
	"go-admin/global"
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
	PercentG     float64 `json:"percentG"`     //日95带宽G 今天计算的日95带宽(G)
	IspDayPrice  float64 `json:"ispPrice"`     //计算今天 运营商的收益
	Usage        float64 `json:"usage"`        //利用率
}

func WatchOnlineUsage() {
	dbList := sdk.Runtime.GetDb()

	for _, d := range dbList {
		var hostList []models.Host
		d.Model(&models.Host{}).Select("host_name,balance").Where("healthy_at >= DATE_SUB(NOW(), INTERVAL 30 MINUTE)").Find(&hostList)

		nowTime := time.Now()
		tenMinutesAgo := fmt.Sprintf("%v", nowTime.Add(-10*time.Minute).Unix())
		endTime := fmt.Sprintf("%v", nowTime.Unix())
		for _, host := range hostList {

			transmitQuery := fmt.Sprintf("sum(rate(phy_nic_network_transmit_bytes_total{instance=\"%v\"}[5m]))*8", host.HostName)

			BandwidthIncome := host.Balance / 1000 //换算成G

			compute := RequestPromResult(tenMinutesAgo, endTime, "60", transmitQuery)
			updateMap := make(map[string]interface{})
			if compute.Empty {
				updateMap["status"] = global.HostOffline
			} else {
				updateMap["status"] = global.HostSuccess
				updateMap["usage"] = utils.RoundDecimalFlot64(3, compute.PercentG/BandwidthIncome)
			}
			d.Model(&models.Host{}).Where("id = ?", host.Id).Updates(updateMap)

		}
	}
}

func RequestPromResult(start, end, setup, query string) *MonitorCompute {

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
		result.PercentBytes = Percent
		result.PercentG = utils.RoundDecimalFlot64(3, Percent/(1024*1024*1024))

	} else {
		result.Empty = true
	}

	return result
}
