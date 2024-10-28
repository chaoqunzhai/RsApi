package models

import (
	"time"

	"go-admin/common/models"
)

type RsHostIncome struct {
	models.Model

	HostId            string       `json:"hostId" gorm:"type:bigint;comment:主机ID"`
	Isp               string       `json:"isp" gorm:"type:bigint;comment:运营商ID"`
	IdcId             string       `json:"idcId" gorm:"type:bigint;comment:IDC ID"`
	BuId              string       `json:"buId" gorm:"type:bigint;comment:业务ID"`
	Income            string       `json:"income" gorm:"type:double;comment:Income"`
	Usage             string       `json:"usage" gorm:"type:double;comment:Usage"`
	Bandwidth95       string       `json:"bandwidth95" gorm:"type:double;comment:Bandwidth95"`
	BandwidthIncome   string       `json:"bandwidthIncome" gorm:"type:double;comment:BandwidthIncome"`
	Estimate          string       `json:"estimate" gorm:"type:double;comment:Estimate"`
	Actual            string       `json:"actual" gorm:"type:double;comment:Actual"`
	SlaPrice          string       `json:"slaPrice" gorm:"type:double;comment:SlaPrice"`
	SlaInfo           string       `json:"slaInfo" gorm:"type:longtext;comment:SlaInfo"`
	SettleStatus      string       `json:"settleStatus" gorm:"type:bigint;comment:SettleStatus"`
	SettleTime        time.Time    `json:"settleTime" gorm:"type:datetime(3);comment:SettleTime"`
	SettleBandwidth   string       `json:"settleBandwidth" gorm:"type:double;comment:SettleBandwidth"`
	TotalBandwidth    string       `json:"totalBandwidth" gorm:"type:double;comment:TotalBandwidth"`
	HeartbeatNum      string       `json:"heartbeatNum" gorm:"type:bigint;comment:HeartbeatNum"`
	NightHeartbeatNum string       `json:"nightHeartbeatNum" gorm:"type:bigint;comment:NightHeartbeatNum"`
	CreatedAt         models.XTime `json:"createdAt" gorm:"comment:创建时间"`
}

func (RsHostIncome) TableName() string {
	return "rs_host_income"
}

func (e *RsHostIncome) GetId() interface{} {
	return e.Id
}
