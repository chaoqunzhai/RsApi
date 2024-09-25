package models

import (
	"go-admin/common/models"
	"go-admin/global"
	"gorm.io/gorm"
)

type Combination struct {
	models.Model
	CreateUser string                 `json:"createUser" gorm:"-"`
	CustomId   int                    `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	HostId     int                    `json:"hostId" gorm:"type:bigint;comment:关联的上线CMDB ID"`
	Desc       string                 `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Code       string                 `json:"code" gorm:"type:varchar(50);comment:组合编号"`
	Status     string                 `json:"status" gorm:"type:int;comment:资产状态"`
	AssetCount int                    `json:"assetCount" gorm:"-"`
	Asset      []AdditionsWarehousing `json:"asset" gorm:"-"`
	Price      float64                `json:"price" gorm:"-"`
	RegionInfo interface{}            `json:"regionInfo" gorm:"-"`
	CustomInfo interface{}            `json:"customInfo" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (Combination) TableName() string {
	return "asset_combination"
}

func (e *Combination) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Combination) GetId() interface{} {
	return e.Id
}

func (e *Combination) AfterFind(tx *gorm.DB) (err error) {
	if row, _ := global.UserDatMap.Get(e.CreateBy); row != nil {

		e.CreateUser = row.Username
	}
	return nil
}
