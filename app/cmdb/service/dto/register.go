/*
*
@Author: chaoqun
* @Date: 2024/7/25 22:49
*/
package dto

type RegisterMetrics struct {
	CPU            int                    `json:"CPU"`
	Memory         string                 `json:"memory"`
	Disk           string                 `json:"disk"`
	Sn             string                 `json:"sn"`
	Hostname       string                 `json:"hostname"`
	Ip             string                 `json:"ip"`
	Business       string                 `json:"business"`
	Kernel         string                 `json:"kernel"`
	BusinessSn     map[string]string      `json:"business_sn"`
	Remark         string                 `json:"remark"`
	Province       string                 `json:"province"`
	City           string                 `json:"city"`
	Isp            string                 `json:"isp"`
	NetDevice      string                 `json:"netDevice"`
	Balance        float64                `json:"balance"`
	TransmitNumber float64                `json:"transmitNumber"`
	ReceiveNumber  float64                `json:"receiveNumber"`
	MemoryMap      map[string]interface{} `json:"memoryMap"`
	ExtendMap      []SoftwareRow          `json:"extendMap"`
}

type SoftwareRow struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}
