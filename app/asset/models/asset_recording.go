package models

import (
	"go-admin/common/models"
	"go-admin/global"
	"gorm.io/gorm"
)

type AssetRecording struct {
	models.Model
	CreateUser string       `json:"createUser" gorm:"-"`
	AssetId    int          `json:"assetId" gorm:"type:bigint;comment:关联资产ID"`
	User       string       `json:"user" gorm:"type:varchar(30);comment:操作人"`
	Type       int          `json:"type" gorm:"comment:操作类型"`
	Info       string       `json:"info" gorm:"type:varchar(100);comment:处理内容"`
	BindOrder  string       `json:"bindOrder" gorm:"type:varchar(50);comment:关联单据"`
	CreatedAt  models.XTime `json:"createdAt" gorm:"comment:创建时间"`
	CreateBy   int          `json:"createBy" gorm:"index;comment:创建者"`
}

func (AssetRecording) TableName() string {
	return "asset_recording"
}

func (e *AssetRecording) GetId() interface{} {
	return e.Id
}

func (e *AssetRecording) AfterFind(tx *gorm.DB) (err error) {
	if row, _ := global.UserDatMap.Get(e.CreateBy); row != nil {

		e.CreateUser = row.Username
	}
	return nil
}
