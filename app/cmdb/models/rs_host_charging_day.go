package models

import (
	"go-admin/common/models"
)

type RsHostChargingDay struct {
	models.Model

	BusinessId string       `json:"businessId" gorm:"type:bigint;comment:切换的业务ID"`
	HostId     string       `json:"hostId" gorm:"type:bigint;comment:关联的主机ID"`
	Cost       string       `json:"cost" gorm:"type:double;comment:计算的费用"`
	Banlance95 string       `json:"banlance95" gorm:"type:double;comment:95带宽值"`
	Sla        string       `json:"sla" gorm:"type:varchar(120);comment:触发SLA原因"`
	Desc       string       `json:"desc" gorm:"type:varchar(120);comment:计费备注"`
	CreatedAt  models.XTime `json:"createdAt" gorm:"comment:创建时间"`
	CreateBy   int          `json:"createBy" gorm:"index;comment:创建者"`
}

func (RsHostChargingDay) TableName() string {
	return "rs_host_charging_day"
}

func (e *RsHostChargingDay) GetId() interface{} {
	return e.Id
}
