package models

import (
	"go-admin/common/models"
	"time"
)

type HostIncome struct {
	Model
	CreatedAt         time.Time    `json:"createdAt" gorm:"comment:计算时间"`
	HostId            int          `json:"hostId" gorm:"index;comment:主机ID"`
	Isp               int          `json:"isp" gorm:"index;comment:运营商ID"`
	IdcId             int          `json:"idcId" gorm:"index;comment:IDC ID"`
	BuId              int          `json:"buId" gorm:"index;comment:业务ID"`
	Income            float64      `json:"income" gorm:"收益"`
	Usage             float64      `json:"usage" gorm:"利用率,单位是%"`
	Bandwidth95       float64      `json:"bandwidth95" gorm:"95带宽,单位是G"`
	BandwidthIncome   float64      `json:"bandwidthIncome" gorm:"计费带宽,单位是G"`
	Estimate          float64      `json:"estimate" gorm:"预估收益"`
	Actual            float64      `json:"actual" gorm:"实际收益"`
	SlaPrice          float64      `json:"slaPrice" gorm:"SLA扣款费用"`
	SlaInfo           string       `json:"slaInfo" gorm:"varchar(50);触发SLA原因"`
	SettleStatus      int          `json:"settleStatus" gorm:"default:1;是否已经结算 1:未结算 2:已结算"`
	SettleTime        models.XTime `json:"settleTime" gorm:"结算时间"`
	SettleBandwidth   float64      `json:"settleBandwidth" gorm:"结算带宽"`
	TotalBandwidth    float64      `json:"totalBandwidth" gorm:"总跑量"`
	HeartbeatNum      int          `json:"heartbeatNum" gorm:"总打点数,通常来说是288个点,5分钟一个点"`
	NightHeartbeatNum int          `json:"nightHeartbeatNum" gorm:"晚高峰打点"`
}

func (HostIncome) TableName() string {
	return "rs_host_income"
}
