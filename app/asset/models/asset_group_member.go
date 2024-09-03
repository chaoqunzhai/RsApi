package models

import (
	"go-admin/common/models"
)

type AssetGroupMember struct {
	models.Model

	AssetGroupId int  `json:"assetGroupId" gorm:"type:int;comment:资产组合编码"`
	AssetId      int  `json:"assetId" gorm:"type:int;comment:资产编码"`
	IsMain       int8 `json:"isMain" gorm:"type:tinyint(1);comment:是否为主资产(1=是,2=否)"`
	models.ModelTime
	models.ControlBy
}

func (AssetGroupMember) TableName() string {
	return "asset_group_member"
}

func (e *AssetGroupMember) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetGroupMember) GetId() interface{} {
	return e.Id
}
