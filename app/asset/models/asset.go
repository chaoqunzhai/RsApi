package models

import (
	"database/sql"
	"go-admin/common/models"
	"go-admin/global"
	"gorm.io/gorm"
)

type AdditionsOrder struct {
	models.Model
	models.ModelTime
	CreateBy    int         `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy    int         `json:"updateBy" gorm:"index;comment:更新者"`
	OrderId     string      `json:"orderId" gorm:"type:varchar(50);index;comment:关联的入库单号"`
	StoreRoomId int         `json:"storeRoomId" gorm:"index;comment:关联库房"`
	Asset       interface{} `json:"asset" gorm:"-"`
	CreateUser  string      `json:"createUser" gorm:"-"`
	Desc        string      `json:"desc" gorm:"size:35;comment:描述信息"` //描述
}

func (e *AdditionsOrder) AfterFind(tx *gorm.DB) (err error) {

	if row, _ := global.UserDatMap.Get(e.CreateBy); row != nil {
		e.CreateUser = row.Username
	}
	return nil
}

func (AdditionsOrder) TableName() string {
	return "asset_additions_order"
}

type AdditionsWarehousing struct {
	models.Model
	CreateUser       string       `json:"createUser" gorm:"-"`
	Code             string       `json:"code"  gorm:"type:varchar(50);comment:资产编码" `
	StoreRoomId      int          `json:"storeRoomId" gorm:"index;comment:关联库房"`
	PurchaseAt       sql.NullTime `json:"-" gorm:"type:datetime(3);comment:采购日期"`
	ExpireAt         sql.NullTime `json:"-" gorm:"type:datetime(3);comment:维保到期日"`
	CategoryId       int64        `json:"categoryId" gorm:"type:bigint;comment:关联的资产分类ID"`
	SupplierId       int64        `json:"supplierId" gorm:"type:bigint;comment:供应商ID"`
	WId              int64        `json:"wId" gorm:"type:bigint;comment:关联的入库单号"`
	OutId            int64        `json:"outId" gorm:"index;comment:关联的出库ID"`
	CombinationId    int          `json:"combinationId" gorm:"index;comment:组合ID"`
	CombinationSN    string       `json:"combinationSn" gorm:"-"`
	Name             string       `json:"name" gorm:"type:varchar(50);comment:资产名称"`
	Spec             string       `json:"spec" gorm:"type:varchar(50);comment:规格型号"`
	Brand            string       `json:"brand" gorm:"type:varchar(50);comment:品牌名称"`
	Sn               string       `json:"sn" gorm:"type:varchar(100);comment:资产SN"`
	Status           int          `json:"status" gorm:"type:int(1);default:1;comment:资产状态"`
	UnitId           int64        `json:"unitId" gorm:"type:bigint;comment:单位"`
	Price            float64      `json:"price" gorm:"comment:价格"`
	UserId           int64        `json:"userId" gorm:"type:bigint;comment:采购人员ID"`
	Desc             string       `json:"desc" gorm:"type:varchar(30);comment:备注"`
	PurchaseAtFormat string       `json:"purchaseAtFormat" gorm:"-"`
	ExpireAtFormat   string       `json:"ExpireAtFormat" gorm:"-"`
	CreatedAt        models.XTime `json:"-" gorm:"comment:创建时间"`
	CreateBy         int          `json:"createBy" gorm:"index;comment:创建者"`
}

func (AdditionsWarehousing) TableName() string {
	return "asset_additions_warehousing"
}

func (e *AdditionsWarehousing) GetId() interface{} {
	return e.Id
}
func (e *AdditionsWarehousing) AfterFind(tx *gorm.DB) (err error) {

	if e.PurchaseAt.Valid {
		e.PurchaseAtFormat = e.PurchaseAt.Time.Format("2006-01-02 15:04:05")

	}
	if e.ExpireAt.Valid {
		e.ExpireAtFormat = e.ExpireAt.Time.Format("2006-01-02 15:04:05")

	}
	if row, _ := global.UserDatMap.Get(e.CreateBy); row != nil {
		e.CreateUser = row.Username
	}
	return nil
}
