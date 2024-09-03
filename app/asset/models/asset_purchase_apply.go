package models

import (
	"go-admin/common/models"
	"time"
)

type AssetPurchaseApply struct {
	models.Model

	ApplyCode     string    `json:"applyCode" gorm:"type:varchar(100);comment:申请单编号"`
	CategoryId    int       `json:"categoryId" gorm:"type:int;comment:资产类型编码"`
	SupplierId    int       `json:"supplierId" gorm:"type:int;comment:供应商编码"`
	ApplyUser     int       `json:"applyUser" gorm:"type:int;comment:申购人编码"`
	Quantity      int64     `json:"quantity" gorm:"type:int;comment:申购数量"`
	Specification string    `json:"specification" gorm:"type:varchar(100);comment:规格型号"`
	Brand         string    `json:"brand" gorm:"type:varchar(100);comment:品牌"`
	Unit          string    `json:"unit" gorm:"type:varchar(50);comment:计量单位"`
	UnitPrice     float64   `json:"unitPrice" gorm:"type:decimal(10,2);comment:预估单价"`
	TotalAmount   float64   `json:"totalAmount" gorm:"type:decimal(10,2);comment:预估金额"`
	ApplyReason   string    `json:"applyReason" gorm:"type:text;comment:申购理由"`
	ApplyAt       time.Time `json:"applyAt" gorm:"type:date;comment:申购日期"`
	Status        int8      `json:"status" gorm:"type:tinyint(1);comment:申购状态(0=待审批, 1=已审批, 2=已驳回, 3=已取消)"`
	Approver      int       `json:"approver" gorm:"type:int;comment:审批人编码"`
	ApproveAt     time.Time `json:"approveAt" gorm:"type:date;comment:审批时间"`
	Remark        string    `json:"remark" gorm:"type:text;comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (AssetPurchaseApply) TableName() string {
	return "asset_purchase_apply"
}

func (e *AssetPurchaseApply) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AssetPurchaseApply) GetId() interface{} {
	return e.Id
}
