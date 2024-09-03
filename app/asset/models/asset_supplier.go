package models

import (
	"go-admin/common/models"
)

type AssetSupplier struct {
	models.Model

	SupplierName  string `json:"supplierName" gorm:"type:varchar(100);comment:供应商名称"`
	ContactPerson string `json:"contactPerson" gorm:"type:varchar(100);comment:联系人"`
	PhoneNumber   string `json:"phoneNumber" gorm:"type:varchar(20);comment:联系电话"`
	Email         string `json:"email" gorm:"type:varchar(100);comment:邮箱"`
	Address       string `json:"address" gorm:"type:varchar(200);comment:地址"`
	Remark        string `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetSupplier) TableName() string {
	return "asset_supplier"
}

func (e *AssetSupplier) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetSupplier) GetId() interface{} {
	return e.Id
}
