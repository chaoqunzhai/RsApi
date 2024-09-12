package models

import (
	"go-admin/common/models"
)

type AssetOutboundMember struct {
	models.Model

	AssetOutboundId   int    `json:"assetOutboundId" gorm:"type:int;comment:资产出库编码"`
	AssetOutboundCode string `json:"assetOutboundCode" gorm:"type:varchar(100);comment:资产出库单号"`
	AssetId           int    `json:"assetId" gorm:"type:int;comment:资产编码"`
	models.ModelTime
	models.ControlBy
}

func (AssetOutboundMember) TableName() string {
	return "asset_outbound_member"
}

func (e *AssetOutboundMember) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetOutboundMember) GetId() interface{} {
	return e.Id
}
