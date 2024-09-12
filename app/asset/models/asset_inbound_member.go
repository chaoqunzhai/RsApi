package models

import (
	"go-admin/common/models"
)

type AssetInboundMember struct {
	models.Model

	AssetInboundId   int    `json:"assetInboundId" gorm:"type:int;comment:资产入库编码"`
	AssetInboundCode string `json:"assetInboundCode" gorm:"type:varchar(100);comment:资产入库单号"`
	AssetId          int    `json:"assetId" gorm:"type:int;comment:资产编码"`
	models.ModelTime
	models.ControlBy
}

func (AssetInboundMember) TableName() string {
	return "asset_inbound_member"
}

func (e *AssetInboundMember) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetInboundMember) GetId() interface{} {
	return e.Id
}
