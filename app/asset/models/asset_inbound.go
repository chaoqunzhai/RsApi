package models

import (
	"time"

	"go-admin/common/models"
)

type AssetInbound struct {
	models.Model

	AssetId     string    `json:"assetId" gorm:"type:int;comment:资产编码"`
	WarehouseId string    `json:"warehouseId" gorm:"type:int;comment:库房编码"`
	InboundFrom string    `json:"inboundFrom" gorm:"type:enum('Purchased','SelfMade','Rented','Other');comment:来源(采购、自产、租赁、其它)"`
	FromCode    string    `json:"fromCode" gorm:"type:varchar(100);comment:来源凭证编码"`
	InboundBy   string    `json:"inboundBy" gorm:"type:int;comment:入库人编码"`
	InboundAt   time.Time `json:"inboundAt" gorm:"type:timestamp;comment:入库时间"`
	Attachment  string    `json:"attachment" gorm:"type:varchar(255);comment:附件"`
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
