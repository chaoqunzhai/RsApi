package models

import (
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/models"
	"gorm.io/gorm"
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
	models.ExtendUserBy
}

func (e *RsHostChargingDay) AfterFind(tx *gorm.DB) (err error) {
	var user models2.SysUser
	userId := e.CreateBy

	tx.Model(&user).Select("user_id,username").Where("user_id = ?", userId).Limit(1).Find(&user)

	if user.UserId > 0 {
		e.UpdatedUser = user.Username

	}
	return
}
func (RsHostChargingDay) TableName() string {
	return "rs_host_charging_day"
}

func (e *RsHostChargingDay) GetId() interface{} {
	return e.Id
}
