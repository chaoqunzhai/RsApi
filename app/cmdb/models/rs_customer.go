package models

import (

	"go-admin/common/models"

)

type RsCustomer struct {
    models.Model
    
    Desc string `json:"desc" gorm:"type:varchar(35);comment:描述信息"` 
    Name string `json:"name" gorm:"type:varchar(20);comment:客户名称"` 
    Region string `json:"region" gorm:"type:varchar(80);comment:省份城市多ID"` 
    Address string `json:"address" gorm:"type:varchar(80);comment:地址"` 
    Level int64 `json:"level" gorm:"type:int;comment:客户等级"` 
    TypeId int64 `json:"typeId" gorm:"type:int;comment:客户类型"` 
    WorkStatus string `json:"workStatus" gorm:"type:int;comment:合作状态"` 
    models.ModelTime
    models.ControlBy
}

func (RsCustomer) TableName() string {
    return "rs_customer"
}

func (e *RsCustomer) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsCustomer) GetId() interface{} {
	return e.Id
}