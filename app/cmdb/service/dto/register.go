/**
@Author: chaoqun
* @Date: 2024/7/25 22:49
*/
package dto

type RegisterHost struct {
	Sn string `json:"sn" form:"sn"`
	HostName string `json:"host_name" form:"hostname"`
	Ip string `json:"ip" form:"ip"`
	CPU string `json:"cpu" form:"cpu"`
	Disk string `json:"disk" form:"disk"`
	Memory string `json:"memory" form:"memory" `
	Remark string `json:"remark" form:"remark" `
	Software []*SoftwareRow `json:"software"`
	Static   []map[string]interface{} `json:"static"`
}
type SoftwareRow struct {
	Key  string `json:"key"`
	Value string `json:"value"`
	Desc string `json:"desc"`
}
