package models

import (
	"go-admin/common/models"
)

type Combination struct {
	models.Model
	HostId     int                    `json:"hostId" gorm:"type:bigint;comment:关联的上线CMDB ID"`
	Desc       string                 `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Code       string                 `json:"code" gorm:"type:varchar(50);comment:组合编号"`
	Status     string                 `json:"status" gorm:"type:int;comment:资产状态"`
	AssetCount int                    `json:"assetCount" gorm:"-"`
	Asset      []AdditionsWarehousing `json:"asset" gorm:"-"`
	Price      float64                `json:"price" gorm:"-"`
	RegionInfo interface{}            `json:"regionInfo" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (Combination) TableName() string {
	return "combination"
}

func (e *Combination) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Combination) GetId() interface{} {
	return e.Id
}
