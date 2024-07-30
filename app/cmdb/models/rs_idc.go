package models

import (
	"go-admin/common/models"
)

type RsIdc struct {
	models.Model
	Layer           int     `json:"layer" gorm:"type:tinyint;comment:排序"`
	Enable          bool    `json:"enable" gorm:"type:tinyint(1);comment:开关"`
	Desc            string  `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Number          int     `json:"number" gorm:"type:bigint;comment:机房编号"`
	Name            string  `json:"name" gorm:"type:varchar(100);comment:机房名称"`
	CustomUser      int     `json:"customUser" gorm:"type:bigint;comment:所属客户"`
	Region          string  `json:"region" gorm:"type:varchar(50);comment:所在地区"`
	Address         string  `json:"address" gorm:"type:varchar(255);comment:详细地址"`
	IpV6            int     `json:"ipV6" gorm:"type:tinyint(1);comment:是否IPV6"`
	TypeId          int     `json:"typeId" gorm:"type:int;comment:机房类型"`
	BusinessUser    int     `json:"businessUser" gorm:"type:bigint;comment:商务人员"`
	WechatName      string  `json:"wechatName" gorm:"type:varchar(100);comment:企业微信群名称"`
	WebHookUrl      string  `json:"webHookUrl" gorm:"type:varchar(100);comment:企业微信webhookUrl"`
	Status          int     `json:"status" gorm:"type:int;comment:机房状态"`
	Belong          int     `json:"belong" gorm:"type:int;comment:机房归属"`
	Isp             int     `json:"isp" gorm:"type:int;comment:运营商"`
	AllBandwidth    string  `json:"allBandwidth" gorm:"type:varchar(35);comment:机房总带宽"`
	AllLine         int     `json:"allLine" gorm:"type:int;comment:机房总线路"`
	Up              string  `json:"up" gorm:"type:varchar(191);comment:上行带宽"`
	Down            string  `json:"down" gorm:"type:varchar(191);comment:下行带宽"`
	Price           float64 `json:"price" gorm:"type:double;comment:单价"`
	ManageLine      int     `json:"manageLine" gorm:"type:int;comment: 管理线路数"`
	ManagerLineCost float64 `json:"managerLineCost" gorm:"type:double;comment:管理线价格"`
	BandwidthType   int     `json:"bandwidthType" gorm:"type:bigint;comment:宽带类型"`
	Charging        int     `json:"charging" gorm:"type:int;comment:计费方式"`
	TransProvince   int     `json:"transProvince" gorm:"type:tinyint(1);comment:是否跨省"`
	MoreDialing     int     `json:"moreDialing" gorm:"type:tinyint(1);comment:是否支持多拨"`
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
