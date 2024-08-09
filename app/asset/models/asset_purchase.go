package models

import (
	"go-admin/common/models"
)

type AssetPurchase struct {
	models.Model

	PurchaseCode  string `json:"purchaseCode" gorm:"type:varchar(100);comment:采购单编号"`
	CategoryId    string `json:"categoryId" gorm:"type:int;comment:资产类型编码"`
	SupplierId    string `json:"supplierId" gorm:"type:int;comment:供应商编码"`
	PurchaseUser  string `json:"purchaseUser" gorm:"type:int;comment:采购人编码"`
	Specification string `json:"specification" gorm:"type:varchar(100);comment:规格型号"`
	Brand         string `json:"brand" gorm:"type:varchar(100);comment:品牌"`
	Quantity      string `json:"quantity" gorm:"type:int;comment:采购数量"`
	Unit          string `json:"unit" gorm:"type:varchar(50);comment:计量单位"`
	UnitPrice     string `json:"unitPrice" gorm:"type:decimal(10,2);comment:采购单价"`
	TotalAmount   string `json:"totalAmount" gorm:"type:decimal(10,2);comment:采购金额"`
	PurchaseAt    string `json:"purchaseAt" gorm:"type:date;comment:采购日期"`
	Remark        string `json:"remark" gorm:"type:text;comment:备注"`
	Attachment    string `json:"attachment" gorm:"type:varchar(255);comment:附件"`
	models.ModelTime
	models.ControlBy
}

func (AssetPurchase) TableName() string {
	return "asset_purchase"
}

func (e *AssetPurchase) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetPurchase) GetId() interface{} {
	return e.Id
}
