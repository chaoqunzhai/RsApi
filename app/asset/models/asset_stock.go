package models

import (
	"go-admin/common/models"
)

type AssetStock struct {
	models.Model

	WarehouseId string `json:"warehouseId" gorm:"type:int;comment:库房编码"`
	CategoryId  string `json:"categoryId" gorm:"type:int;comment:资产类别编码"`
	Quantity    string `json:"quantity" gorm:"type:int;comment:资产库存数量"`
	Remark      string `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetStock) TableName() string {
	return "asset_stock"
}

func (e *AssetStock) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetStock) GetId() interface{} {
	return e.Id
}
