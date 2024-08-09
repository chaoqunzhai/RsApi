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

type ResultType struct {
	Metric MetricType      `json:"metric"`
	Value  [][]interface{} `json:"values"`
}

type QueryData struct {
	ResultType string       `json:"resultType"`
	Result     []ResultType `json:"result"`
}

type QueryInfo struct {
	Status string    `json:"status"`
	Data   QueryData `json:"data"`
}

func GetPromResult(u *url.URL) (result *QueryInfo, err error) {

	u.RawQuery = u.Query().Encode()

	urlPath := u.String()
	result = &QueryInfo{}
	username := config.ExtConfig.Prometheus.Username
	password := config.ExtConfig.Prometheus.Password

	// 设置 Authorization 头
	auth := fmt.Sprintf("%s:%s", username, password)
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	//fmt.Println("URL", urlPath)
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
