package models

import (
	"time"

	"go-admin/common/models"
)

type AssetReturn struct {
	models.Model

	AssetId      string    `json:"assetId" gorm:"type:int;comment:资产编码"`
	ReturnPerson string    `json:"returnPerson" gorm:"type:int;comment:退还人编码"`
	Reason       string    `json:"reason" gorm:"type:varchar(50);comment:退还原因"`
	ReturnAt     time.Time `json:"returnAt" gorm:"type:timestamp;comment:退还时间"`
	Attachment   string    `json:"attachment" gorm:"type:varchar(255);comment:附件"`
	Remark       string    `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetReturn) TableName() string {
	return "asset_return"
}

func (e *AssetReturn) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetReturn) GetId() interface{} {
	return e.Id
}
