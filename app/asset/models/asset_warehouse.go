package models

import (
	"go-admin/common/models"
)

type AssetWarehouse struct {
	models.Model

	WarehouseName   string `json:"warehouseName" gorm:"type:varchar(100);comment:库房名称"`
	AdministratorId int    `json:"administratorId" gorm:"type:int;comment:管理员编码"`
	Remark          string `json:"remark" gorm:"type:text;comment:备注"`
	CreateBy        int    `json:"createBy" gorm:"index;comment:创建者"`
	models.ModelTime
	models.ControlBy
}

func (AssetWarehouse) TableName() string {
	return "asset_warehouse"
}

func (e *AssetWarehouse) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetWarehouse) GetId() interface{} {
	return e.Id
}
