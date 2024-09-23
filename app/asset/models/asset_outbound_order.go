package models

import (
	"go-admin/common/models"
)

type AssetOutboundOrder struct {
	models.Model

	Desc           string                 `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Code           string                 `json:"code" gorm:"type:varchar(50);comment:出库编码"`
	CustomId       string                 `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	PhoneNumber    string                 `json:"phoneNumber" gorm:"type:varchar(20);comment:PhoneNumber"`
	Region         string                 `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	Ems            string                 `json:"ems" gorm:"type:varchar(20);comment:物流公司"`
	TrackingNumber string                 `json:"trackingNumber" gorm:"type:varchar(30);comment:物流单号"`
	Address        string                 `json:"address" gorm:"type:varchar(255);comment:联系地址"`
	User           string                 `json:"user" gorm:"type:varchar(50);comment:联系人"`
	IdcId          string                 `json:"idcId" gorm:"type:bigint;comment:idcId"`
	Count          string                 `json:"count" gorm:"type:bigint;comment:出库数量"`
	Asset          []AdditionsWarehousing `json:"asset" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (AssetOutboundOrder) TableName() string {
	return "asset_outbound_order"
}

func (e *AssetOutboundOrder) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetOutboundOrder) GetId() interface{} {
	return e.Id
}
