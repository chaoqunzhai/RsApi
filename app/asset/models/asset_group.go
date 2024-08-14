package models

import (
	"go-admin/common/models"
)

type AssetGroup struct {
	models.Model

	GroupName   string `json:"groupName" gorm:"type:varchar(128);comment:资产组合名称"`
	MainAssetId int    `json:"mainAssetId" gorm:"type:int;comment:主资产编码"`
	Remark      string `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetGroup) TableName() string {
	return "asset_group"
}

func (e *AssetGroup) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetGroup) GetId() interface{} {
	return e.Id
}
