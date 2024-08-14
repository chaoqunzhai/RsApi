package models

import (
	"time"

	"go-admin/common/models"
)

type AssetOutbound struct {
	models.Model

	AssetId     int       `json:"assetId" gorm:"type:int;comment:资产编码"`
	WarehouseId int       `json:"warehouseId" gorm:"type:int;comment:库房编码"`
	OutboundTo  int       `json:"outboundTo" gorm:"type:int;comment:出库去向(客户编码)"`
	OutboundBy  int       `json:"outboundBy" gorm:"type:int;comment:出库人编码"`
	OutboundAt  time.Time `json:"outboundAt" gorm:"type:timestamp;comment:出库时间"`
	Remark      string    `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetOutbound) TableName() string {
	return "asset_outbound"
}

func (e *AssetOutbound) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetOutbound) GetId() interface{} {
	return e.Id
}
