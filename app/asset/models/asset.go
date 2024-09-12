package models

import (
	"go-admin/common/models"
)

type Asset struct {
	models.Model

	AssetCode     string  `json:"assetCode" gorm:"type:varchar(128);comment:资产编号"`
	SnCode        string  `json:"snCode" gorm:"type:varchar(255);comment:SN编码"`
	CategoryId    int     `json:"categoryId" gorm:"type:int;comment:资产类别"`
	Specification string  `json:"specification" gorm:"type:varchar(100);comment:规格型号"`
	Brand         string  `json:"brand" gorm:"type:varchar(100);comment:品牌"`
	Unit          string  `json:"unit" gorm:"type:varchar(50);comment:计量单位"`
	UnitPrice     float64 `json:"unitPrice" gorm:"type:decimal(10,2);comment:单价"`
	Status        int8    `json:"status" gorm:"type:tinyint(1);comment:状态(1=在库, 2=出库, 3=在用, 4=处置)"`
	Remark        string  `json:"remark" gorm:"type:text;comment:备注"`
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
