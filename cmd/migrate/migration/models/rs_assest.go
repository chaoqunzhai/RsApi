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

type Additions struct {
	Model
	CreatedAt   time.Time    `json:"createdAt" gorm:"comment:创建时间"`
	PurchaseAt  sql.NullTime `json:"purchaseAt" gorm:"comment:采购日期"`
	StoreRoomId int          `json:"storeRoomId" gorm:"index;comment:关联库房"`
	ExpireAt    sql.NullTime `json:"expireAt" gorm:"comment:维保到期日"`
	CategoryId  int          `json:"categoryId" gorm:"index;comment:关联的资产分类ID"`
	SupplierId  int          `json:"supplierId"  gorm:"index;comment:供应商ID"`
	WId         int          `json:"WId" gorm:"index;comment:关联的入库ID"`
	Name        string       `json:"name"  gorm:"type:varchar(50);comment:资产名称" `
	Spec        string       `json:"spec" gorm:"type:varchar(50);comment:规格型号" `
	Brand       string       `json:"brand" gorm:"type:varchar(50);comment:品牌名称" `
	Sn          string       `json:"sn" gorm:"type:varchar(100);comment:资产SN" `
	UnitId      int          `json:"unitId" gorm:"comment:单位"`
	Price       string       `json:"price" gorm:"type:varchar(50);comment:价格"`
	UserId      int          `json:"userId" gorm:"index;comment:采购人员ID"`
	Desc        string       `json:"desc" gorm:"type:varchar(30);comment:备注"`
}

func (Additions) TableName() string {
	return "additions_warehousing"
}
