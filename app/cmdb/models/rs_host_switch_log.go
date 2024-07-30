package models

import (
	"go-admin/common/models"
)

type RsHostSwitchLog struct {
	models.Model

	JobId      string       `json:"jobId" gorm:"type:varchar(50);comment:任务ID"`
	HostId     string       `json:"hostId" gorm:"type:bigint;comment:切换的主机ID"`
	BusinessId string       `json:"businessId" gorm:"type:bigint;comment:切换的新业务ID"`
	BusinessSn string       `json:"businessSn" gorm:"type:varchar(30);comment:原来的业务SN"`
	Desc       string       `json:"desc" gorm:"type:varchar(120);comment:切换业务备注"`
	CreatedAt  models.XTime `json:"createdAt" gorm:"comment:创建时间"`
	CreateBy   int          `json:"createBy" gorm:"index;comment:创建者"`
}

func (RsHostSwitchLog) TableName() string {
	return "rs_host_switch_log"
}

func (e *RsHostSwitchLog) GetId() interface{} {
	return e.Id
}
