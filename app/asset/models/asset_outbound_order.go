package models

import (
	"go-admin/common/models"
	"go-admin/global"
	"gorm.io/gorm"
)

type AssetOutboundOrder struct {
	models.Model
	CreateUser     string                 `json:"createUser" gorm:"-"`
	Desc           string                 `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Code           string                 `json:"code" gorm:"type:varchar(50);comment:出库编码"`
	CustomId       int                    `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	PhoneNumber    string                 `json:"phoneNumber" gorm:"type:varchar(20);comment:PhoneNumber"`
	Region         string                 `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	Ems            string                 `json:"ems" gorm:"type:varchar(20);comment:物流公司"`
	TrackingNumber string                 `json:"trackingNumber" gorm:"type:varchar(30);comment:物流单号"`
	Address        string                 `json:"address" gorm:"type:varchar(255);comment:联系地址"`
	UserId         int                    `json:"userId" gorm:"comment:联系人"`
	IdcId          int                    `json:"idcId" gorm:"type:bigint;comment:idcId"`
	Asset          []AdditionsWarehousing `json:"asset" gorm:"-"`
	RegionInfo     interface{}            `json:"regionInfo" gorm:"-"`
	CustomInfo     interface{}            `json:"customInfo" gorm:"-"`
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

func (e *AssetOutboundOrder) AfterFind(tx *gorm.DB) (err error) {
	if row, _ := global.UserDatMap.Get(e.CreateBy); row != nil {

		e.CreateUser = row.Username
	}
	return nil
}
