package models

import (
	"go-admin/common/models"
)

type Asset struct {
	models.Model

	AssetCode     string `json:"assetCode" gorm:"type:varchar(128);comment:资产编号"`
	SnCode        string `json:"snCode" gorm:"type:varchar(255);comment:SN编码"`
	CategoryId    string `json:"categoryId" gorm:"type:int;comment:资产类别"`
	Specification string `json:"specification" gorm:"type:varchar(100);comment:规格型号"`
	Brand         string `json:"brand" gorm:"type:varchar(100);comment:品牌"`
	Unit          string `json:"unit" gorm:"type:varchar(50);comment:计量单位"`
	UnitPrice     string `json:"unitPrice" gorm:"type:decimal(10,2);comment:单价"`
	Attachment    string `json:"attachment" gorm:"type:varchar(255);comment:附件"`
	Status        string `json:"status" gorm:"type:enum('InStock','OutStock');comment:状态(在库、出库)"`
	Remark        string `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (Asset) TableName() string {
	return "asset"
}

func (e *Asset) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Asset) GetId() interface{} {
	return e.Id
}
