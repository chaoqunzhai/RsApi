package models

import (
	"go-admin/common/models"
)

type RsContacts struct {
	models.Model

	Desc       string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	UserName   string `json:"userName" gorm:"type:varchar(30);comment:用户名"`
	CustomerId int64  `json:"customerId" gorm:"type:int;comment:客户ID"`
	BuUser     int64  `json:"buUser" gorm:"type:int;comment:商务人员"`
	Phone      string `json:"phone" gorm:"type:varchar(20);comment:电话号码"`
	Landline   string `json:"landline" gorm:"type:varchar(20);comment:座机号"`
	Region     string `json:"region" gorm:"type:varchar(80);comment:管理区域,也是城市ID"`
	Email      string `json:"email" gorm:"type:varchar(20);comment:电话号码"`
	Address    string `json:"address" gorm:"type:varchar(100);comment:地址"`
	Department string `json:"department" gorm:"type:varchar(30);comment:部门"`
	Duties     string `json:"duties" gorm:"type:varchar(30);comment:职务"`
	models.ModelTime
	models.ControlBy
}

func (RsContacts) TableName() string {
	return "rs_contacts"
}

func (e *RsContacts) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsContacts) GetId() interface{} {
	return e.Id
}
