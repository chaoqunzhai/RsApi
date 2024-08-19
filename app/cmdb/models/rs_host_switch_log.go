package models

import (
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsHostSwitchLog struct {
	models.Model

	JobId      string `json:"jobId" gorm:"type:varchar(50);comment:任务ID"`
	HostId     int    `json:"hostId" gorm:"type:bigint;comment:切换的主机ID"`
	BuTargetId int    `json:"bu_target_id" gorm:"index;comment:切换的新业务ID"`
	BuSource   string `json:"bu_source"  gorm:"type:varchar(30);index;comment:原来的业务中文名称"`
	BuEnSource string `json:"bu_en_source"  gorm:"type:varchar(30);index;comment:原来的业务英文名称"`

	BusinessInfo interface{}  `json:"businessInfo" gorm:"-"`
	Desc         string       `json:"desc" gorm:"type:varchar(30);comment:切换业务备注"`
	CreatedAt    models.XTime `json:"createdAt" gorm:"comment:创建时间"`
	CreateBy     int          `json:"createBy" gorm:"index;comment:创建者"`
	models.ExtendUserBy
}

func (e *RsHostSwitchLog) AfterFind(tx *gorm.DB) (err error) {
	var user models2.SysUser
	userId := e.CreateBy
	tx.Model(&user).Select("user_id,username").Where("user_id = ?", userId).Limit(1).Find(&user)

	if user.UserId > 0 {
		e.UpdatedUser = user.Username

	}
	return
}
func (RsHostSwitchLog) TableName() string {
	return "rs_host_switch_log"
}

func (e *RsHostSwitchLog) GetId() interface{} {
	return e.Id
}
