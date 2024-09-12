package models

import (
	"time"

	"go-admin/common/models"
)

type AssetDisposal struct {
	models.Model

	AssetId        int       `json:"assetId" gorm:"type:int;comment:资产编码"`
	DisposalPerson int       `json:"disposalPerson" gorm:"type:int;comment:处置人编码"`
	Reason         string    `json:"reason" gorm:"type:varchar(50);comment:处置原因"`
	DisposalWay    int8      `json:"disposalWay" gorm:"type:tinyint(1);comment:处置方式(1=报废, 2=出售, 3=出租, 4=退租, 5=捐赠, 6=其它)"`
	DisposalType   int8      `json:"disposalType" gorm:"type:tinyint(1);comment:处置地点类型(1=机房, 2=库房)"`
	LocationId     int       `json:"locationId" gorm:"type:int;comment:处置地点编码(机房编码/库房编码)"`
	Amount         float64   `json:"amount" gorm:"type:decimal(10,2);comment:处置金额"`
	DisposalAt     time.Time `json:"disposalAt" gorm:"type:timestamp;comment:处置时间"`
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
