package models

import (
	"database/sql"
	"time"
)

type AssetWarehouse struct {
	RichGlobal
	WarehouseName   string `json:"warehouseName" gorm:"type:varchar(50);index;comment:名字"`
	AdministratorId int    `json:"administratorId"`
	Remark          string `json:"remark" gorm:"type:text;comment:备注"`
}

func (AssetWarehouse) TableName() string {
	return "asset_warehouse"
}

type AssetSupplier struct {
	RichGlobal
	SupplierName  string `json:"supplierName" gorm:"type:varchar(50);comment:名字"`
	ContactPerson string `json:"contactPerson" gorm:"type:varchar(50)"`
	PhoneNumber   string `json:"phoneNumber" gorm:"type:varchar(20)"`
	Email         string `json:"email" gorm:"type:varchar(50)"`
	Address       string `json:"address" gorm:"type:varchar(255)"`
	Remark        string `json:"remark" gorm:"type:varchar(255)"`
}

func (AssetSupplier) TableName() string {
	return "asset_supplier"
}

type AdditionsOrder struct {
	RichGlobal
	OrderId     string `json:"orderId" gorm:"type:varchar(50);index;comment:关联的入库单号"`
	StoreRoomId int    `json:"storeRoomId" gorm:"index;comment:关联库房"`
}

func (AdditionsOrder) TableName() string {
	return "asset_additions_order"
}

type AdditionsWarehousing struct {
	Model
	CreateBy      int          `json:"createBy" gorm:"index;comment:创建者"`
	Code          string       `json:"code"  gorm:"type:varchar(50);comment:资产编码" `
	CreatedAt     time.Time    `json:"createdAt" gorm:"comment:创建时间"`
	PurchaseAt    sql.NullTime `json:"purchaseAt" gorm:"comment:采购日期"`
	StoreRoomId   int          `json:"storeRoomId" gorm:"index;comment:关联库房"`
	ExpireAt      sql.NullTime `json:"expireAt" gorm:"comment:维保到期日"`
	HostId        int          `json:"hostId"  gorm:"comment:关联的CMDB上线的主机ID"`
	CategoryId    int          `json:"categoryId" gorm:"index;comment:关联的资产分类ID"`
	SupplierId    int          `json:"supplierId"  gorm:"index;comment:供应商ID"`
	WId           int          `json:"WId" gorm:"index;comment:关联的入库ID"`
	OutId         int          `json:"outId" gorm:"index;comment:关联的出库ID"`
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
	return "asset_additions_warehousing"
}

type Combination struct {
	RichGlobal
	Code     string `json:"code"  gorm:"type:varchar(50);comment:组合编号" `
	CustomId int    `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	Status   int    `json:"status" gorm:"index;type:int(1);default:1;comment:资产状态"`
	IdcId    int    `json:"idcId" gorm:"type:bigint;comment:关联IDC,定时更新"`
	HostId   int    `json:"hostId" gorm:"type:bigint;comment:关联的上线CMDB ID"`
}

func (Combination) TableName() string {
	return "asset_combination"
}

type AssetRecording struct {
	Model
	AssetType int       `json:"assetType" gorm:"default:1;comment:资产类型 1:资产 2:组合"`
	CreateBy  int       `json:"createBy" gorm:"index;comment:创建者"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	AssetId   int       `json:"assetId"  gorm:"comment:关联资产ID"`
	User      string    `json:"user" gorm:"index;comment:操作人"`
	Type      int       `json:"type" gorm:"comment:操作类型"`
	Info      string    `json:"info" gorm:"comment:处理内容"`
	BindOrder string    `json:"bindOrder" gorm:"type:varchar(50);comment:关联单据"`
}

func (AssetRecording) TableName() string {
	return "asset_recording"
}
