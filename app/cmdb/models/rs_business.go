package models

import (
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsBusiness struct {
	models.Model
	Status    int    `json:"status" gorm:"type:tinyint;comment:排序" comment:"状态"`
	Layer     int    `json:"layer" gorm:"type:tinyint;comment:排序"`
	Desc      string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Name      string `json:"name" gorm:"type:varchar(50);comment:业务云名称"`
	Algorithm string `json:"algorithm" gorm:"type:varchar(30);comment:计费算法标记"`
	EnName    string `json:"enName" gorm:"index;type:varchar(30);comment:业务英文名字"`
	models.ExtendUserBy
	models.ModelTime
	models.ControlBy
}

func (RsBusiness) TableName() string {
	return "rs_business"
}
func (e *RsBusiness) AfterFind(tx *gorm.DB) (err error) {
	var user models2.SysUser
	userId := e.CreateBy
	if e.UpdateBy != 0 {
		userId = e.UpdateBy
	}
	tx.Model(&user).Select("user_id,username").Where("user_id = ?", userId).Limit(1).Find(&user)

	if user.UserId > 0 {
		e.UpdatedUser = user.Username

	}
	return
}

func (e *RsBusiness) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsBusiness) GetId() interface{} {
	return e.Id
}
