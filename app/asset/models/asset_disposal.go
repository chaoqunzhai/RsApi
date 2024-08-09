package models

import (
	"time"

	"go-admin/common/models"
)

type AssetDisposal struct {
	models.Model

	AssetId        string    `json:"assetId" gorm:"type:int;comment:资产编码"`
	DisposalPerson string    `json:"disposalPerson" gorm:"type:int;comment:处置人编码"`
	Reason         string    `json:"reason" gorm:"type:varchar(50);comment:处置原因"`
	DisposalType   string    `json:"disposalType" gorm:"type:enum('Scrap','Sell','Rent','ReturnRent','Donate','Other');comment:处置方式(报废、出售、出租、退租、捐赠、其它)"`
	Amount         string    `json:"amount" gorm:"type:decimal(10,2);comment:处置金额"`
	DisposalAt     time.Time `json:"disposalAt" gorm:"type:timestamp;comment:处置时间"`
	Attachment     string    `json:"attachment" gorm:"type:varchar(255);comment:附件"`
	Remark         string    `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetDisposal) TableName() string {
	return "asset_disposal"
}

func (e *AssetDisposal) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetDisposal) GetId() interface{} {
	return e.Id
}
