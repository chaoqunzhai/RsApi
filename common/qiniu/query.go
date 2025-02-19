package qiniu

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)


var (
	token = "ad977442ad064efd86887f0d2807545f"
	encryptKey = "5e74cb2ad4e2c9e2d3f59a1e6c8a5d4999df48e5dd69871d2798e0a146b91ee9"
	host = "https://ecscm-openapi.ksyun.com"
)

type Result struct {
	Data []ResultData `json:"data"`
	Status int `json:"status"`
}
type ResultData struct {
	Date        string  `json:"date"`
	DeviceId    string  `json:"device_id"`
	Isp         string  `json:"isp"`
	SrmCode     string  `json:"srm_code"`
	ChargeDay95 float64 `json:"charge_day95"`
	SrmChannel  string  `json:"srm_channel"`
}
func getHMAC(plainText string, encryptKey string) string {
	key := []byte(encryptKey)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(plainText))
	return hex.EncodeToString(h.Sum(nil))
}
func GetQueryQiNiu(url string,params map[string]interface{})  (dat *Result,err error) {
	timeStamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	pan := token + timeStamp + url
	sk := getHMAC(pan, encryptKey)


	headers := map[string]string{
		"token":      token,
		"timestamps": timeStamp,
		"secret-key": sk,
	}

	req, err := http.NewRequest("GET", host+url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, fmt.Sprintf("%v",value))
	}
	req.URL.RawQuery = query.Encode()

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)


	result :=&Result{}
	if resp.StatusCode != 200 {
		msg :=fmt.Sprint("Request failed with status code:", resp.StatusCode,resp)
		return nil, errors.New(msg)
	}
	err = json.Unmarshal(responseBody,&result)
	if err!=nil{
		return nil, err
	}
	return result, nil
}