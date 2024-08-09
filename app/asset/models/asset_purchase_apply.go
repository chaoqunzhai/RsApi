package models

import (
	"go-admin/common/models"
)

type AssetPurchaseApply struct {
	models.Model

	ApplyCode     string `json:"applyCode" gorm:"type:varchar(100);comment:申请单编号"`
	CategoryId    string `json:"categoryId" gorm:"type:int;comment:资产类型编码"`
	SupplierId    string `json:"supplierId" gorm:"type:int;comment:供应商编码"`
	ApplyUser     string `json:"applyUser" gorm:"type:int;comment:申购人编码"`
	Quantity      string `json:"quantity" gorm:"type:int;comment:申购数量"`
	Specification string `json:"specification" gorm:"type:varchar(100);comment:规格型号"`
	Brand         string `json:"brand" gorm:"type:varchar(100);comment:品牌"`
	Unit          string `json:"unit" gorm:"type:varchar(50);comment:计量单位"`
	UnitPrice     string `json:"unitPrice" gorm:"type:decimal(10,2);comment:预估单价"`
	TotalAmount   string `json:"totalAmount" gorm:"type:decimal(10,2);comment:预估金额"`
	ApplyReason   string `json:"applyReason" gorm:"type:text;comment:申购理由"`
	ApplyAt       string `json:"applyAt" gorm:"type:date;comment:申购日期"`
	Status        string `json:"status" gorm:"type:enum('Pending','Approved','Rejected');comment:申购状态(待审批、已审批、已拒绝)"`
	Approver      string `json:"approver" gorm:"type:int;comment:审批人编码"`
	ApproveAt     string `json:"approveAt" gorm:"type:date;comment:审批时间"`
	Attachment    string `json:"attachment" gorm:"type:varchar(255);comment:附件"`
	Remark        string `json:"remark" gorm:"type:text;comment:备注"`
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
