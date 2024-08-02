package models

import (
	"go-admin/common/models"
)

type RsTag struct {
	models.Model

	Layer  string `json:"layer" gorm:"type:tinyint;comment:排序"`
	Enable string `json:"enable" gorm:"type:tinyint(1);comment:开关"`
	Desc   string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Name   string `json:"name" gorm:"type:varchar(50);comment:业务云名称"`
	models.ExtendUserBy
	models.ModelTime
	models.ControlBy
}

func (RsTag) TableName() string {
	return "rs_tag"
}

func (e *RsTag) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsTag) GetId() interface{} {
	return e.Id
}
