package prometheus

import (
	"fmt"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/utils"
	"go-admin/config"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type MonitorResult struct {
	Data    []interface{}  `json:"data"`
	Compute MonitorCompute `json:"compute"`
}
type MonitorCompute struct {
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
	Avg     float64 `json:"avg"`
	Percent float64 `json:"percent"`
}

func requestPromResult(query string, req *dto.RsHostMonitorFlow) interface{} {

	result := MonitorResult{}
	//查询普罗米修斯数据
	queryUrl, err := url.Parse(func() string {
		vv, _ := url.JoinPath(config.ExtConfig.Prometheus.Endpoint, "/api/v1/query_range")
		return vv
	}())

	parameters := url.Values{}
	parameters.Add("start", req.Start)
	parameters.Add("end", req.End)
	parameters.Add("step", fmt.Sprintf("%v", req.Setup))
	parameters.Add("query", query)

	queryUrl.RawQuery = parameters.Encode()

	ProResult, err := GetPromResult(queryUrl)

	if err != nil {
		return result
	}

	XData := make([]interface{}, 0)
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
						XData = append(XData, []interface{}{addUnixTime.Format(time.DateTime), 0})
						//进行计算
						XValue = append(XValue, 0)

					}
				}

			}
			parentUnix = unixFloat

			valueStr := row[1].(string)
			valueFloat, _ := strconv.ParseFloat(valueStr, 64)

			unixTime := time.Unix(int64(unixFloat), 0)

			XData = append(XData, []interface{}{unixTime.Format(time.DateTime), valueStr})
			//进行计算
			XValue = append(XValue, valueFloat)

		}
	}
	sort.Float64s(XValue)

	result.Data = XData
	//计算95值
	result.Compute.Percent = utils.Percentile(XValue, 0.95)
	//计算max
	result.Compute.Max = utils.Max(XValue)
	//计算最小
	result.Compute.Min = utils.Min(XValue)

	//计算平均
	result.Compute.Avg = utils.Avg(XValue)

	return result
}

func Transmit(hostname string, req *dto.RsHostMonitorFlow) map[string]interface{} {
	//主机监控内容

	transmitQuery := fmt.Sprintf("sum(rate(phy_nic_network_transmit_bytes_total{instance=\"%v\"}[5m]))*8", hostname)

	transmitData := requestPromResult(transmitQuery, req)

	receiveQuery := fmt.Sprintf("sum(rate(phy_nic_network_receive_bytes_total{instance=\"%v\"}[5m]))*8", hostname)

	receiveData := requestPromResult(receiveQuery, req)

	provinceOutQuery := fmt.Sprintf("avg(province_out_percent{instance=\"%v\"}) by (instance) * sum(rate(phy_nic_network_transmit_bytes_total{instance=\"%v\"}[5m]))by(instance)*8",
		hostname, hostname,
	)

	provinceOutData := requestPromResult(provinceOutQuery, req)

	response := map[string]interface{}{
		"transmit":     transmitData,
		"receive":      receiveData,
		"province_out": provinceOutData,
	}

	return response
}
