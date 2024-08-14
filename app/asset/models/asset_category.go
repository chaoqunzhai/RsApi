package models

import (
	"go-admin/common/models"
)

type AssetCategory struct {
	models.Model

	CategoryName string `json:"categoryName" gorm:"type:varchar(100);comment:类别名称"`
	Remark       string `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetCategory) TableName() string {
	return "asset_category"
}

func (e *AssetCategory) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetCategory) GetId() interface{} {
	return e.Id
}
