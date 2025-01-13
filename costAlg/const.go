package costAlg

import (
	"fmt"
	"go-admin/global"
	"time"
)

type Business struct {
	Name    string             `json:"name"`
	EnName  string             `json:"enName"`
	Id      int                `json:"id"`
	AlgId   int                `json:"algId"`
	AlgName string             `json:"algName"`
	Price   float64            `json:"price"`
	IspCnf  map[string]*IspCnf `json:"ispCnf"`
	Host    []*Host            `json:"host"`
	SlaConf map[int]*SlaConf   `json:"slaConf"` //触发sla的规则
}

type SlaConf struct {
	Id    int
	Name  string `json:"name"` //规则名称
	Desc  string `json:"desc"`
	Start string `json:"start"`
	End   string `json:"end"`
	Limit int    `json:"limit"` //超过多少分钟
}
type IspCnf struct {
	ConstId     int     //常量Id
	Name        string  //运营商名字
	Price       float64 //当月总价(3000)
	Day         int     //当月天数
	AvgDayPrice float64 //计算的每天价格 = 当月总价(3000) / 自然月(30)
}
type LabelRow struct {
	Label string `json:"label" form:"label" `
	Value string `json:"value" form:"value" `
}
type PromReq struct {
	Setup int
	Start string
	End   string
}
type Host struct {
	HostId          int                        `json:"hostId"`
	HostName        string                     `json:"hostName"`
	HostSn          string                     `json:"hostSn"`
	Balance       float64      `json:"balance" gorm:"type:varchar(50);comment:总带宽"`
	IdcId           int                        `json:"idcId"`
	BuId            int                        `json:"buId"`
	AlgDay          string                     `json:"algDay"`       //计算天数
	PriceCompute    map[string]*MonitorCompute `json:"priceCompute"` //今天的收益 运营商:收益
	IspId           int                        `json:"ispId"`
	BandwidthIncome float64                    `json:"bandwidthIncome"` //计费带宽 就是多少条线 * 单条线路带宽
	CustomId int64 `json:"custom_id"` //客户ID
	ContractId  int64 `json:"contract_id"` //合同ID
	ContractAlg  *ContractAlg //合同费用配置
	AlgNote []string
	BuSn []*LabelRow `json:"buSn"`
}
type ContractAlg struct {
	LinePrice       float64 `json:"LinePrice" gorm:"comment:业务线单价"`
	ManagerLineCost float64 `json:"managerLineCost" gorm:"comment:管理线价格"`
	IspId           int                        `json:"ispId"`

}
type MonitorCompute struct {
	Empty          bool    `json:"empty"`
	Max            float64 `json:"max"`
	Min            float64 `json:"min"`
	Avg            float64 `json:"avg"`
	TotalBandwidth float64 `json:"total"`
	HeartbeatNum   int     `json:"heartbeatNum"`
	SLA            *SlaRow `json:"sla"`
	PercentBytes   float64 `json:"percentBytes"` //今天计算的95带宽(bytes)
	PercentG       float64 `json:"percentG"`     //日95带宽G 今天计算的日95带宽(G)
	IspDayPrice    float64 `json:"ispDayPrice"`  //计算今天 运营商的收益真实收益
	IspCnf         *IspCnf `json:"ispCnf"`       //运营商的计费配置
	Usage          float64 `json:"usage"`        //利用率
	DayCost float64 `json:"day_cost" ` //每天成本
	MonthlyCost float64 `json:"monthly_cost" ` //月成本
	CostAlgorithm string `json:"cost_algorithm" ` //成本算法,如果没有配置客户 合同 是展示文本

}

type SlaRow struct {
	Info    string  `json:"info"`
	Trigger bool    `json:"trigger"`
	Price   float64 `json:"price"`
}
type RsHostMonitorFlow struct {
	Start string `form:"start"`
	End   string `form:"end"`
	Setup int    `form:"setup"`
}

func ParseTime(timeStr string) (t time.Time, err error) {

	layout := fmt.Sprintf("2006-01-02 15:04")
	thisNow := fmt.Sprintf("%v %v", time.Now().Format("2006-01-02"), timeStr)
	return time.ParseInLocation(layout, thisNow, global.LOC)
}
