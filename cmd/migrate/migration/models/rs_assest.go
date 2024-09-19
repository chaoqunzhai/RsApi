package models

import (
	"database/sql"
	"time"
)

type AdditionsOrder struct {
	RichGlobal
	OrderId     string `json:"orderId" gorm:"type:varchar(50);index;comment:关联的入库单号"`
	StoreRoomId int    `json:"storeRoomId" gorm:"index;comment:关联库房"`
}

func (AdditionsOrder) TableName() string {
	return "additions_order"
}

type AdditionsWarehousing struct {
	Model
	Code          string       `json:"code"  gorm:"type:varchar(50);comment:资产编码" `
	CreatedAt     time.Time    `json:"createdAt" gorm:"comment:创建时间"`
	PurchaseAt    sql.NullTime `json:"purchaseAt" gorm:"comment:采购日期"`
	StoreRoomId   int          `json:"storeRoomId" gorm:"index;comment:关联库房"`
	ExpireAt      sql.NullTime `json:"expireAt" gorm:"comment:维保到期日"`
	CategoryId    int          `json:"categoryId" gorm:"index;comment:关联的资产分类ID"`
	SupplierId    int          `json:"supplierId"  gorm:"index;comment:供应商ID"`
	WId           int          `json:"WId" gorm:"index;comment:关联的入库ID"`
	CombinationId int          `json:"combinationId" gorm:"index;default:0;comment:组合ID"`
	Name          string       `json:"name"  gorm:"type:varchar(50);comment:资产名称" `
	Spec          string       `json:"spec" gorm:"type:varchar(50);comment:规格型号" `
	Brand         string       `json:"brand" gorm:"type:varchar(50);comment:品牌名称" `
	Sn            string       `json:"sn" gorm:"type:varchar(100);comment:资产SN" `
	Status        int          `json:"status" gorm:"index;type:int(1);default:1;comment:资产状态"`
	UnitId        int          `json:"unitId" gorm:"comment:单位"`
	Price         float64      `json:"price" gorm:"comment:价格"`
	UserId        int          `json:"userId" gorm:"index;comment:采购人员ID"`
	Desc          string       `json:"desc" gorm:"type:varchar(30);comment:备注"`
}

func (AdditionsWarehousing) TableName() string {
	return "additions_warehousing"
}

type Combination struct {
	RichGlobal
	Code   string `json:"code"  gorm:"type:varchar(50);comment:组合编号" `
	Status int    `json:"status" gorm:"index;type:int(1);default:1;comment:资产状态"`
}

func (Combination) TableName() string {
	return "combination"
}
