package models

import (
	"go-admin/common/models"
)

type RsIdc struct {
	models.Model

	Desc     string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Number   int    `json:"number" gorm:"type:bigint;comment:机房编号"`
	Name     string `json:"name" gorm:"type:varchar(100);comment:机房名称"`
	BuId     int    `json:"buId" gorm:"type:bigint;comment:商务人员"`
	CustomId int    `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	Region   string `json:"region" gorm:"type:varchar(50);comment:所在地区"`
	Address  string `json:"address" gorm:"type:varchar(255);comment:详细地址"`
	IpV6     int    `json:"ipV6" gorm:"type:tinyint(1);comment:是否IPV6"`
	TypeId   int    `json:"typeId" gorm:"type:int;comment:机房类型"`

	WechatName    string `json:"wechatName" gorm:"type:varchar(100);comment:企业微信群名称"`
	WebHookUrl    string `json:"webHookUrl" gorm:"type:varchar(100);comment:企业微信webhookUrl"`
	Status        int    `json:"status" gorm:"type:int;comment:机房状态"`
	Belong        int    `json:"belong" gorm:"type:int;comment:机房归属"`
	TransProvince int    `json:"transProd" gorm:"default:1;comment:是否跨省"`
	MoreDialing   int    `json:"moreDialing" gorm:"default:1;comment:是否支持多拨"`
	models.ModelTime
	models.ControlBy
}

func (RsIdc) TableName() string {
	return "rs_idc"
}

func (e *RsIdc) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsIdc) GetId() interface{} {
	return e.Id
}
