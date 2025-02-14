package costAlg

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Result  struct {
		CurrentPage int `json:"currentPage"`
		PageSize    int `json:"pageSize"`
		Total       int `json:"total"`
		TotalPage   int `json:"totalPage"`
		Records     []struct {
			SettleStatus       string  `json:"settleStatus"`
			WorkDate           string  `json:"workDate"`
			Bandwidth          float64 `json:"bandwidth"`
			IspCode            string  `json:"ispCode"`
			IspName            string  `json:"ispName"`
			Province           string  `json:"province"`
			City               string  `json:"city"`
			OnlineCode         string  `json:"onlineCode"`
			OnlineStatus       string  `json:"onlineStatus"`
			Sn                 string  `json:"sn"`
			TotalBandwidth     float64 `json:"totalBandwidth"`
			ResendBandwidth    float64 `json:"resendBandwidth"`
			HeartbeatNum       int     `json:"heartbeatNum"`
			NightHeartbeatNum  int     `json:"nightHeartbeatNum"`
			Amount             float64 `json:"amount"`
			Price              int     `json:"price"`
			IsA6               string  `json:"isA6"`
			IsProvince         string  `json:"isProvince"`
			IsP6               string  `json:"isP6"`
		} `json:"records"`
	} `json:"result"`
	Timestamp int64 `json:"timestamp"`
}
type OpenApiLinWu struct {
	Orm         *gorm.DB
	RunTime     map[string]string
	workDate,url string

}
// genHeaders 生成请求头
func (c *OpenApiLinWu)genHeaders(payload string) map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"appid":        "10052",
		"timeStamp":    strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	// 生成签名
	var body string
	if payload != "" {
		body = "10052" + "cafb75dbefdb4ed1afa35f0483cf6594" + payload
	} else {
		body = "10052" + "cafb75dbefdb4ed1afa35f0483cf6594"
	}

	baseRes := base64.StdEncoding.EncodeToString([]byte(body))
	hash := sha256.Sum256([]byte(baseRes))
	sign := fmt.Sprintf("%x", hash)

	headers["sign"] = sign
	return headers
}

func (c *OpenApiLinWu) SetupDb(dbs map[string]*gorm.DB) {
	for _, db := range dbs {
		c.Orm = db
	}
	c.RunTime = make(map[string]string)
	c.workDate = time.Now().AddDate(0,0,-1).Format(time.DateOnly)
	c.url = "https://openapi.linkfog.cn/openapi/v2/queryDeviceRevenueByPage"
}

func (c *OpenApiLinWu) LoopData(parentDay int) {

	for i := 0; i < parentDay; i++ {
		// 计算日期：当前时间 + i 天
		futureDate := time.Now().AddDate(0,0,-i).Format(time.DateOnly)
		// 格式化日期为 YYYY-MM-DD
		c.StartHostAmount(futureDate)
	}

}
func (c *OpenApiLinWu) StartHostAmount(workDate string) {
	//请求


	jsonData := map[string]interface{}{
		"workDate": workDate,
		"pageNo":   1,
		"pageSize": 10000,
	}
	jsonValue, _ := json.Marshal(jsonData)

	payloadStr := string(jsonValue)

	// 生成请求头
	headers := c.genHeaders(payloadStr)

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// 设置请求头
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

	body, _ := ioutil.ReadAll(resp.Body)


	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	for _,row :=range response.Result.Records {


		var hostRow models.Host
		c.Orm.Model(&models.Host{}).Select("id").Where("sn = ?",row.Sn).Limit(1).Find(&hostRow)
		if hostRow.Id == 0 {
			fmt.Printf("领雾 sn:%v cmdb not found\n",row.Sn)
			continue}

		if row.SettleStatus == "待结算"{
			fmt.Printf("领雾 sn:%v 待结算 cmdb.id:%v\n",row.Sn,hostRow.Id)
			continue}

		var hostIncome models.HostIncome

		c.Orm.Model(&models.HostIncome{}).Where("host_id = ? and alg_day = ?",hostRow.Id,c.workDate).Limit(1).Find(&hostIncome)
		if hostIncome.Id == 0 {
			fmt.Printf("领雾 sn:%v  cmdb.id:%v HostIncome.alg_day:%v 不存在\n",row.Sn,hostRow.Id,c.workDate)
			continue

		}
		hostIncome.SettleStatus = 2
		hostIncome.SettleTime = c.workDate
		hostIncome.SettlePrice = row.Amount
		// 备注：结算带宽 = 总带宽 - 重传带宽
		hostIncome.SettleBandwidth = row.Bandwidth
		c.Orm.Save(&hostIncome)
	}
}