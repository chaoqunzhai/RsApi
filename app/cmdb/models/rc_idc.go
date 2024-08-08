package models

import (
	"go-admin/common/models"
)

type RcIdc struct {
	models.Model

	Layer         string `json:"layer" gorm:"type:tinyint;comment:排序"`
	Enable        string `json:"enable" gorm:"type:tinyint(1);comment:开关"`
	Desc          string `json:"desc" gorm:"type:varchar(255);comment:备注"`
	Name          string `json:"name" gorm:"type:varchar(100);comment:机房名称"`
	Status        string `json:"status" gorm:"type:int;comment:机房状态"`
	Belong        string `json:"belong" gorm:"type:int;comment:机房归属"`
	TypeId        string `json:"typeId" gorm:"type:int;comment:机房类型"`
	BusinessUser  string `json:"businessUser" gorm:"type:bigint;comment:商务人员"`
	CustomUser    string `json:"customUser" gorm:"type:bigint;comment:所属客户"`
	Region        string `json:"region" gorm:"type:bigint;comment:所在区域"`
	Charging      string `json:"charging" gorm:"type:int;comment:计费方式"`
	Price         string `json:"price" gorm:"type:double;comment:单价"`
	WeChatName    string `json:"weChatName" gorm:"type:varchar(255);comment:企业微信"`
	IpV6          string `json:"ipV6" gorm:"type:tinyint(1);comment:是否IPV6"`
	TransProvince string `json:"transProvince" gorm:"type:tinyint(1);comment:是否跨省"`
	Address       string `json:"address" gorm:"type:varchar(255);comment:详细地址"`
	models.ModelTime
	models.ControlBy
}

func (RcIdc) TableName() string {
	return "rc_idc"
}

func (e *RcIdc) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcIdc) GetId() interface{} {
	return e.Id
}
