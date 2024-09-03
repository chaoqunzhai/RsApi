package models

import (
	"time"

	"go-admin/common/models"
)

type AssetInbound struct {
	models.Model

	InboundCode string    `json:"inboundCode" gorm:"type:varchar(100);comment:入库单号"`
	WarehouseId int       `json:"warehouseId" gorm:"type:int;comment:库房编码"`
	InboundFrom int8      `json:"inboundFrom" gorm:"type:tinyint(1);comment:来源(1=直接入库、2=采购入库)"`
	FromCode    string    `json:"fromCode" gorm:"type:varchar(100);comment:来源凭证编码(采购编码)"`
	InboundBy   int       `json:"inboundBy" gorm:"type:int;comment:入库人编码"`
	InboundAt   time.Time `json:"inboundAt" gorm:"type:timestamp;comment:入库时间"`
	Remark      string    `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetInbound) TableName() string {
	return "asset_inbound"
}

func (e *AssetInbound) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetInbound) GetId() interface{} {
	return e.Id
}
