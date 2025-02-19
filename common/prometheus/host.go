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
	Title   string         `json:"title"`
	Data    []interface{}  `json:"data"`
	Compute MonitorCompute `json:"compute"`
}
type MonitorCompute struct {
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
	Avg     float64 `json:"avg"`
	Percent float64 `json:"percent"`
}

//	func funUnit(value float64) string  {
//		var unit string
//		if (value > 999999999) {
//			unit = fmt.Sprintf("%.2f",value / 1000000000)
//		} else if (Number(value) > 999999) {
//			unit = fmt.Sprintf("%.2f",value / 1000000000)
//			n = (Number(value / 1000000)).toFixed(2) + " M";
//		} else if (Number(value) > 999) {
//			unit = fmt.Sprintf("%.2f",value / 1000)
//			n = (Number(value / 1000)).toFixed(2) + " K";
//		} else {
//			n = Number(value).toFixed(2) + " b";
//			unit = fmt.Sprintf("%.2f",value)
//		}
//	}

func RequestQueryPromResult(title, query string, req *dto.RsHostMonitorFlow, isMb bool) string {

	//查询普罗米修斯数据
	queryUrl, err := url.Parse(func() string {
		vv, _ := url.JoinPath(config.ExtConfig.Prometheus.Endpoint, "/api/v1/query")
		return vv
	}())

	parameters := url.Values{}
	parameters.Add("time", req.Start)

	parameters.Add("query", query)
	queryUrl.RawQuery = parameters.Encode()

	ProResult, err := GetPromQueryResult(queryUrl)

	if err != nil {

		return ""
	}
	if len(ProResult.Data.Result) > 0 {
		if len(ProResult.Data.Result[0].Value) > 1 {
			return fmt.Sprintf("%v", ProResult.Data.Result[0].Value[1])
		}
	}
	return ""
}
func RequestPromResult(title, query string, req *dto.RsHostMonitorFlow, isMb bool) interface{} {

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

	ProResult, err := GetPromRangeResult(queryUrl)

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

			if isMb {
				valueFloat = valueFloat * 1000000
			}
			XData = append(XData, []interface{}{unixTime.Format(time.DateTime), valueFloat})
			//进行计算
			XValue = append(XValue, valueFloat)

		}
	}
	sort.Float64s(XValue)

	result.Title = title
	result.Data = XData
	//计算95值
	if len(XValue) > 1 {
		result.Compute.Percent = utils.Percentile(XValue, 0.95)
		//计算max
		result.Compute.Max = utils.Max(XValue)
		//计算最小
		result.Compute.Min = utils.Min(XValue)

		//计算平均
		result.Compute.Avg = utils.Avg(XValue)
	}

	return result
}

func Transmit(hostname string, req *dto.RsHostMonitorFlow) map[string]interface{} {
	//主机监控内容

	transmitQuery := fmt.Sprintf("sum(rate(phy_nic_network_transmit_bytes_total{instance=\"%v\"}[5m]))*8", hostname)

	//fmt.Println("transmitQuery!", transmitQuery)
	transmitData := RequestPromResult("上行", transmitQuery, req, false)

	receiveQuery := fmt.Sprintf("sum(rate(phy_nic_network_receive_bytes_total{instance=\"%v\"}[5m]))*8", hostname)

	receiveData := RequestPromResult("下行", receiveQuery, req, false)

	provinceOutQuery := fmt.Sprintf("v4_province_out{instance=\"%v\"}", hostname)

	provinceOutData := RequestPromResult("出省流量", provinceOutQuery, req, false)

	response := map[string]interface{}{
		"transmit":     transmitData,
		"receive":      receiveData,
		"province_out": provinceOutData,
	}

	return response
}

func DianXinTransmit(hostname string, req *dto.RsHostMonitorFlow) map[string]interface{} {
	//主机监控内容

	transmitQuery := fmt.Sprintf("sum(flow_bandwidth_by_minute{instance=\"%v\"})", hostname)

	transmitData := RequestPromResult("上行", transmitQuery, req, true)
	receiveQuery := fmt.Sprintf("sum(flow_download_bandwidth_by_minute{instance=\"%v\"})", hostname)

	receiveData := RequestPromResult("下行", receiveQuery, req, true)

	response := map[string]interface{}{
		"transmit": transmitData,
		"receive":  receiveData,
	}

	return response
}
