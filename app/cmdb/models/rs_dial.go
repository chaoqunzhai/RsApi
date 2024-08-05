package models

import (

	"go-admin/common/models"

)

type RsDial struct {
    models.Model
    
    Layer string `json:"layer" gorm:"type:tinyint;comment:排序"` 
    Enable int64 `json:"enable" gorm:"type:tinyint(1);comment:开关"` 
    Desc string `json:"desc" gorm:"type:varchar(35);comment:描述信息"` 
    Number string `json:"number" gorm:"type:varchar(60);comment:账号"` 
    User string `json:"user" gorm:"type:varchar(30);comment:用户名"` 
    Pass string `json:"pass" gorm:"type:varchar(50);comment:密码"` 
    Status int64 `json:"status" gorm:"type:int;comment:拨号状态,1:正常 非1:异常"` 
    IdcId int64 `json:"idcId" gorm:"type:bigint;comment:关联的IDC"` 
    HostId int64 `json:"hostId" gorm:"type:bigint;comment:关联主机ID"` 
    DeviceId int64 `json:"deviceId" gorm:"type:bigint;comment:关联网卡ID"` 
    models.ModelTime
    models.ControlBy
}

func (RsDial) TableName() string {
    return "rs_dial"
}

func (e *RsDial) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsDial) GetId() interface{} {
	return e.Id
}