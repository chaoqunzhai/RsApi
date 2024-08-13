package models

import (
	"go-admin/common/models"
)

type RsCustom struct {
	models.Model

	Desc        string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Name        string `json:"name" gorm:"type:varchar(20);comment:客户名称"`
	Type        int    `json:"type" gorm:"type:bigint;comment:客户类型,customer_type"`
	Cooperation int    `json:"cooperation" gorm:"type:bigint;comment:合作状态,work_status"`
	Region      string `json:"region" gorm:"type:varchar(50);comment:所在地区"`
	Address     string `json:"address" gorm:"type:varchar(120);comment:地址"`
	models.ModelTime
	models.ControlBy
}

func (RsCustom) TableName() string {
	return "rs_custom"
}

func (e *RsCustom) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsCustom) GetId() interface{} {
	return e.Id
}
