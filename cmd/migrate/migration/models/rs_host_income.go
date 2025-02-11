package models

import (
	"go-admin/common/utils"
	"gorm.io/gorm"
	"time"
)

type HostIncome struct {
	Model
	AlgDay            string       `json:"algDay"  gorm:"varchar(15);index;comment:计算时间"`
	CreatedAt         time.Time    `json:"createdAt" gorm:"comment:计算时间"`
	HostId            int          `json:"hostId" gorm:"index;comment:主机ID"`
	Isp               int          `json:"isp" gorm:"index;comment:运营商ID"`
	IdcId             int          `json:"idcId" gorm:"index;comment:IDC ID"`
	BuId              int          `json:"buId" gorm:"index;comment:业务名称"`
	Income            float64      `json:"income" gorm:"预估收益"`
	Usage             float64      `json:"usage" gorm:"利用率,单位是%"`
	Bandwidth95       float64      `json:"bandwidth95" gorm:"95计费带宽"`
	BandwidthIncome   float64      `json:"bandwidthIncome" gorm:"计费带宽,单位是G"`
	AvgBuDayPrice       float64      `json:"avgBuDayPrice" gorm:"计算每天业务的价格=运营商费用/当月天数"`
	RetryId           int          `json:"retryId" gorm:"重算任务ID"`
	RetryPrice        float64      `json:"retryPrice" gorm:"重算价格"`
	SlaPrice          float64      `json:"slaPrice" gorm:"SLA扣款费用"`
	SlaInfo           string       `json:"slaInfo" gorm:"varchar(50);触发SLA原因"`
	SettleStatus      int          `json:"settleStatus" gorm:"default:1;是否已经结算 1:未结算 2:已结算"`
	SettleTime        string `json:"settleTime" gorm:"varchar(15);结算时间"`
	SettleBandwidth   float64      `json:"settleBandwidth" gorm:"结算带宽"`
	SettlePrice       float64      `json:"settlePrice" gorm:"结算收益"`
	TotalBandwidth    float64      `json:"totalBandwidth" gorm:"总跑量"`
	HeartbeatNum      int          `json:"heartbeatNum" gorm:"总打点数,通常来说是288个点,5分钟一个点"`
	NightHeartbeatNum int          `json:"nightHeartbeatNum" gorm:"晚高峰打点"`
	DayCost float64 `json:"day_cost" gorm:"每天成本"`
	CostAlgorithm string `json:"cost_algorithm" gorm:"type:varchar(100);成本算法"`
	RecordM bool `json:"record_m"  gorm:"default:0;comment:是否已经记录月记录中"`
}

func (HostIncome) TableName() string {
	return "rs_host_income"
}
type HostIncomeMonth struct {
	Model
	Month string `json:"month"`
	HostId            int64         `json:"hostId" gorm:"index;comment:主机ID"`
	Income            float64      `json:"income" gorm:"预估收益"`
	Cost float64 `json:"cost" gorm:"成本"`
	GrossProfit interface{} `json:"gross_profit" gorm:"-"`
}
func (e *HostIncomeMonth) AfterFind(tx *gorm.DB) (err error) {

	e.Income = utils.RoundDecimal(e.Income)
	e.Cost = utils.RoundDecimal(e.Cost)

	return nil
}
func (HostIncomeMonth) TableName() string {
	return "rs_host_income_month"
}