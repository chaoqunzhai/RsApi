package prometheus

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-admin/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type MetricType struct {
	NodeName string `json:"node_name"`
}

type ResultRangeType struct {
	Metric MetricType      `json:"metric"`
	Value  [][]interface{} `json:"values"`
}

type QueryRangeData struct {
	ResultType string            `json:"resultType"`
	Result     []ResultRangeType `json:"result"`
}

type QueryRangeInfo struct {
	Status string         `json:"status"`
	Data   QueryRangeData `json:"data"`
}

type QueryNodeInfo struct {
	Status string         `json:"status"`
	Data   QueryNodeInfoData `json:"data"`
}
type QueryData struct {
	ResultType string       `json:"resultType"`
	Result     []ResultType `json:"result"`
}
type ResultType struct {
	Metric struct {
	} `json:"metric"`
	Value []interface{} `json:"value"`
}

type QueryNodeInfoData struct {
	ResultType string       `json:"resultType"`
	Result     []ResultNodeInfo `json:"result"`
}
type ResultNodeInfo struct {
	Metric NodeInfoMetric `json:"metric"`
	Value []interface{} `json:"value"`
}
type NodeInfoMetric struct {
	Name           string `json:"__name__"`
	Business       string `json:"business"`
	DeviceCity     string `json:"device_city"`
	DeviceIsp      string `json:"device_isp"`
	DeviceProvince string `json:"device_province"`
	Domainname     string `json:"domainname"`
	Instance       string `json:"instance"`
	Machine        string `json:"machine"`
	Nodename       string `json:"nodename"`
	Release        string `json:"release"`
	Remark         string `json:"remark"`
	Sn string `json:"sn"`
	Sysname        string `json:"sysname"`
	Version        string `json:"version"`
}
type QueryInfo struct {
	Status string    `json:"status"`
	Data   QueryData `json:"data"`
}


func GetPromNodeInfoResult(u *url.URL) (result *QueryNodeInfo, err error) {

	u.RawQuery = u.Query().Encode()

	urlPath := u.String()
	result = &QueryNodeInfo{}
	username := config.ExtConfig.Prometheus.Username
	password := config.ExtConfig.Prometheus.Password

	// 设置 Authorization 头
	auth := fmt.Sprintf("%s:%s", username, password)
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	// 创建一个新的请求
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", authEncoded))
	httpClient := &http.Client{Timeout: 10 * time.Second}

	r, err := httpClient.Do(req)
	if err != nil {
		return result, err
	}

	defer func() {
		_ = r.Body.Close()
	}()

	resultBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("NewDecoder err!", err)
		return result, err
	}

	if UnmarshalErr := json.Unmarshal(resultBody, &result); UnmarshalErr != nil {

		fmt.Println("数据序列化失败", UnmarshalErr)
		return
	}

	return result, err
}
func GetPromRangeResult(u *url.URL) (result *QueryRangeInfo, err error) {

	u.RawQuery = u.Query().Encode()

	urlPath := u.String()
	result = &QueryRangeInfo{}
	username := config.ExtConfig.Prometheus.Username
	password := config.ExtConfig.Prometheus.Password

	// 设置 Authorization 头
	auth := fmt.Sprintf("%s:%s", username, password)
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	// 创建一个新的请求
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", authEncoded))
	httpClient := &http.Client{Timeout: 10 * time.Second}

	r, err := httpClient.Do(req)
	if err != nil {
		return result, err
	}

	defer func() {
		_ = r.Body.Close()
	}()

	resultBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("NewDecoder err!", err)
		return result, err
	}

	if UnmarshalErr := json.Unmarshal(resultBody, &result); UnmarshalErr != nil {

		fmt.Println("数据序列化失败", UnmarshalErr)
		return
	}

	return result, err
}

func GetPromQueryResult(u *url.URL) (result *QueryInfo, err error) {

	u.RawQuery = u.Query().Encode()

	urlPath := u.String()
	result = &QueryInfo{}
	username := config.ExtConfig.Prometheus.Username
	password := config.ExtConfig.Prometheus.Password

	// 设置 Authorization 头
	auth := fmt.Sprintf("%s:%s", username, password)
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	// 创建一个新的请求
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", authEncoded))
	httpClient := &http.Client{Timeout: 10 * time.Second}

	r, err := httpClient.Do(req)
	if err != nil {
		return result, err
	}

	defer func() {
		_ = r.Body.Close()
	}()

	resultBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("NewDecoder err!", err)
		return result, err
	}

	if UnmarshalErr := json.Unmarshal(resultBody, &result); UnmarshalErr != nil {

		fmt.Println("数据序列化失败", UnmarshalErr)
		return
	}

	return result, err
}
