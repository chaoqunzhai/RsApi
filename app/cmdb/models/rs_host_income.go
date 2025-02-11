package models

import (
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsHostIncome struct {
	models.Model
	AlgDay            string          `json:"algDay"  gorm:"index;comment:计算时间"`
	HostId            int             `json:"hostId" gorm:"type:bigint;comment:主机ID"`
	Isp               string          `json:"isp" gorm:"type:bigint;comment:运营商ID"`
	IdcId             string          `json:"idcId" gorm:"type:bigint;comment:IDC ID"`
	BuId              int             `json:"buId" gorm:"type:bigint;comment:业务ID"`
	Income            float64         `json:"income" gorm:"type:double;comment:Income"`
	AvgDayPrice       float64         `json:"avgDayPrice" gorm:"计算每天的价格=运营商费用/当月天数"`
	Usage             string          `json:"usage" gorm:"type:double;comment:Usage"`
	Bandwidth95       string          `json:"bandwidth95" gorm:"type:double;comment:Bandwidth95"`
	BandwidthIncome   string          `json:"bandwidthIncome" gorm:"type:double;comment:BandwidthIncome"`
	Estimate          string          `json:"estimate" gorm:"type:double;comment:Estimate"`
	SlaPrice          string          `json:"slaPrice" gorm:"type:double;comment:SlaPrice"`
	SlaInfo           string          `json:"slaInfo" gorm:"type:longtext;comment:SlaInfo"`
	SettleStatus      string          `json:"settleStatus" gorm:"type:bigint;comment:SettleStatus"`
	SettleTime        string `json:"settleTime" gorm:"varchar(15);结算时间"`
	SettleBandwidth   string          `json:"settleBandwidth" gorm:"type:double;comment:SettleBandwidth"`
	SettlePrice       float64         `json:"settlePrice" gorm:"结算收益"`
	TotalBandwidth    string          `json:"totalBandwidth" gorm:"type:double;comment:TotalBandwidth"`
	HeartbeatNum      string          `json:"heartbeatNum" gorm:"type:bigint;comment:HeartbeatNum"`
	NightHeartbeatNum string          `json:"nightHeartbeatNum" gorm:"type:bigint;comment:NightHeartbeatNum"`
	CreatedAt         models.DayXTime `json:"createdAt" gorm:"comment:创建时间"`
	DayCost float64 `json:"day_cost" gorm:"每天成本"`
	MonthlyCost float64 `json:"monthly_cost" gorm:"月成本"`
	CostAlgorithm string `json:"cost_algorithm" gorm:"type:varchar(100);成本算法"`
	HostName          string          `json:"hostName" gorm:"-"`
	Status            int             `json:"status" gorm:"-"`
	BuName            string          `json:"buName" gorm:"-"`
	HostRow interface{} `json:"host_row" gorm:"-"`
}

func (RsHostIncome) TableName() string {
	return "rs_host_income"
}
func (e *RsHostIncome) AfterFind(tx *gorm.DB) (err error) {
	//var hostModel models2.Host
	//
	//if e.HostId == 0 {
	//	return
	//}
	//tx.Model(&hostModel).Select("host_name,id,status").Where("id = ?", e.HostId).Limit(1).Find(&hostModel)
	//
	//if hostModel.Id > 0 {
	//	e.HostName = hostModel.HostName
	//	e.Status = hostModel.Status
	//}
	return
}
func (e *RsHostIncome) GetId() interface{} {
	return e.Id
}
