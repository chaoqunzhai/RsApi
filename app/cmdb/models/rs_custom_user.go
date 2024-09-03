package models

import (
	"go-admin/common/models"
)

type RsCustomUser struct {
	models.Model

	Desc     string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	UserName string `json:"userName" gorm:"type:varchar(50);comment:姓名"`
	CustomId int    `json:"customId" gorm:"type:bigint;comment:所属客户"`
	BuId     int    `json:"buId" gorm:"type:bigint;comment:所属商务人员"`
	Phone    string `json:"phone" gorm:"type:varchar(20);comment:联系号码"`
	Email    string `json:"email" gorm:"type:varchar(50);comment:联系邮箱"`
	Region   string `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	Dept     string `json:"dept" gorm:"type:varchar(30);comment:部门"`
	Duties   string `json:"duties" gorm:"type:varchar(30);comment:职务"`
	Address  string `json:"address" gorm:"type:varchar(255);comment:详细地址"`
	models.ModelTime
	models.ControlBy
}

func (RsCustomUser) TableName() string {
	return "rs_custom_user"
}

func (e *RsCustomUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsCustomUser) GetId() interface{} {
	return e.Id
}
