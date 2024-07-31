package models

import (
	"go-admin/common/models"
)

type RsBusiness struct {
	models.Model

	Layer     int    `json:"layer" gorm:"type:tinyint;comment:排序"`
	Enable    bool   `json:"enable" gorm:"type:tinyint(1);comment:开关"`
	Desc      string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Name      string `json:"name" gorm:"type:varchar(50);comment:业务云名称"`
	Algorithm string `json:"algorithm" gorm:"type:varchar(120);comment:算法备注"`

	models.ModelTime
	models.ControlBy
}

func (RsBusiness) TableName() string {
	return "rs_business"
}

func (e *RsBusiness) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsBusiness) GetId() interface{} {
	return e.Id
}
