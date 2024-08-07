package models

import (
	"go-admin/common/models"
)

type RsDial struct {
	models.Model

	Desc            string      `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	CustomUser      int64       `json:"customUser" gorm:"type:bigint;comment:所属客户"`
	Isp             int64       `json:"isp" gorm:"type:int;comment:运营商"`
	Up              string      `json:"up" gorm:"type:varchar(191);comment:上行带宽"`
	Down            string      `json:"down" gorm:"type:varchar(191);comment:下行带宽"`
	Charging        int64       `json:"charging" gorm:"type:int;comment:计费方式"`
	Price           float64     `json:"price" gorm:"type:double;comment:业务线单价"`
	ManagerLineCost float64     `json:"managerLineCost" gorm:"type:double;comment:管理线价格"`
	BandwidthType   int64       `json:"bandwidthType" gorm:"type:bigint;comment:宽带类型"`
	TransProvince   int64       `json:"transProvince" gorm:"type:tinyint(1);comment:是否跨省"`
	MoreDialing     int64       `json:"moreDialing" gorm:"type:tinyint(1);comment:是否支持多拨"`
	IsManager       int64       `json:"isManager" gorm:"type:tinyint(1);comment:是否管理线"`
	Account         string      `json:"account" gorm:"type:varchar(25);comment:账号"`
	Pass            string      `json:"pass" gorm:"type:varchar(20);comment:密码"`
	DialName        string      `json:"dialName" gorm:"type:varchar(20);comment:线路名称"`
	Status          int64       `json:"status" gorm:"type:int;comment:拨号状态,1:已拨通 0:待使用 -1:拨号异常"`
	Source          int64       `json:"source" gorm:"type:int;comment:拨号状态,0:录入 1:自动上报"`
	IdcId           int64       `json:"idcId" gorm:"type:bigint;comment:关联的IDC"`
	HostId          int64       `json:"hostId" gorm:"type:bigint;comment:关联主机ID"`
	DeviceId        int64       `json:"deviceId" gorm:"type:bigint;comment:关联网卡ID"`
	RunTime         string      `json:"runTime"  gorm:"type:varchar(35);comment:启用时间"`
	IdcInfo         interface{} `json:"idcInfo" gorm:"-"`
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
