package models

import (
	"go-admin/common/models"
)

type AssetRecording struct {
	models.Model

	AssetId   int          `json:"assetId" gorm:"type:bigint;comment:关联资产ID"`
	User      string       `json:"user" gorm:"type:varchar(30);comment:操作人"`
	Type      int          `json:"type" gorm:"comment:操作类型"`
	Info      string       `json:"info" gorm:"type:varchar(100);comment:处理内容"`
	BindOrder string       `json:"bindOrder" gorm:"type:varchar(50);comment:关联单据"`
	CreatedAt models.XTime `json:"createdAt" gorm:"comment:创建时间"`
}

func (AssetRecording) TableName() string {
	return "asset_recording"
}

func (e *AssetRecording) GetId() interface{} {
	return e.Id
}
