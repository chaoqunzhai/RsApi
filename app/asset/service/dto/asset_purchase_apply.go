package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type AssetPurchaseApplyGetPageReq struct {
	dto.Pagination `search:"-"`
	ApplyCode      string `form:"applyCode"  search:"type:exact;column:apply_code;table:asset_purchase_apply" comment:"申请单编号"`
	CategoryId     string `form:"categoryId"  search:"type:exact;column:category_id;table:asset_purchase_apply" comment:"资产类型编码"`
	SupplierId     string `form:"supplierId"  search:"type:exact;column:supplier_id;table:asset_purchase_apply" comment:"供应商编码"`
	ApplyUser      string `form:"applyUser"  search:"type:exact;column:apply_user;table:asset_purchase_apply" comment:"申购人编码"`
	Specification  string `form:"specification"  search:"type:exact;column:specification;table:asset_purchase_apply" comment:"规格型号"`
	Brand          string `form:"brand"  search:"type:exact;column:brand;table:asset_purchase_apply" comment:"品牌"`
	ApplyAt        string `form:"applyAt"  search:"type:exact;column:apply_at;table:asset_purchase_apply" comment:"申购日期"`
	Approver       string `form:"approver"  search:"type:exact;column:approver;table:asset_purchase_apply" comment:"审批人编码"`
	ApproveAt      string `form:"approveAt"  search:"type:exact;column:approve_at;table:asset_purchase_apply" comment:"审批时间"`
	AssetPurchaseApplyOrder
}

type AssetPurchaseApplyOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:asset_purchase_apply"`
	ApplyCode     string `form:"applyCodeOrder"  search:"type:order;column:apply_code;table:asset_purchase_apply"`
	CategoryId    string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:asset_purchase_apply"`
	SupplierId    string `form:"supplierIdOrder"  search:"type:order;column:supplier_id;table:asset_purchase_apply"`
	ApplyUser     string `form:"applyUserOrder"  search:"type:order;column:apply_user;table:asset_purchase_apply"`
	Quantity      string `form:"quantityOrder"  search:"type:order;column:quantity;table:asset_purchase_apply"`
	Specification string `form:"specificationOrder"  search:"type:order;column:specification;table:asset_purchase_apply"`
	Brand         string `form:"brandOrder"  search:"type:order;column:brand;table:asset_purchase_apply"`
	Unit          string `form:"unitOrder"  search:"type:order;column:unit;table:asset_purchase_apply"`
	UnitPrice     string `form:"unitPriceOrder"  search:"type:order;column:unit_price;table:asset_purchase_apply"`
	TotalAmount   string `form:"totalAmountOrder"  search:"type:order;column:total_amount;table:asset_purchase_apply"`
	ApplyReason   string `form:"applyReasonOrder"  search:"type:order;column:apply_reason;table:asset_purchase_apply"`
	ApplyAt       string `form:"applyAtOrder"  search:"type:order;column:apply_at;table:asset_purchase_apply"`
	Status        string `form:"statusOrder"  search:"type:order;column:status;table:asset_purchase_apply"`
	Approver      string `form:"approverOrder"  search:"type:order;column:approver;table:asset_purchase_apply"`
	ApproveAt     string `form:"approveAtOrder"  search:"type:order;column:approve_at;table:asset_purchase_apply"`
	Remark        string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_purchase_apply"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_purchase_apply"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_purchase_apply"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_purchase_apply"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_purchase_apply"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_purchase_apply"`
}

func (m *AssetPurchaseApplyGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetPurchaseApplyInsertReq struct {
	Id            int       `json:"-" comment:"主键"` // 主键
	ApplyCode     string    `json:"applyCode" comment:"申请单编号"`
	CategoryId    int       `json:"categoryId" comment:"资产类型编码"`
	SupplierId    int       `json:"supplierId" comment:"供应商编码"`
	ApplyUser     int       `json:"applyUser" comment:"申购人编码"`
	Quantity      int64     `json:"quantity" comment:"申购数量"`
	Specification string    `json:"specification" comment:"规格型号"`
	Brand         string    `json:"brand" comment:"品牌"`
	Unit          string    `json:"unit" comment:"计量单位"`
	UnitPrice     float64   `json:"unitPrice" comment:"预估单价"`
	TotalAmount   float64   `json:"totalAmount" comment:"预估金额"`
	ApplyReason   string    `json:"applyReason" comment:"申购理由"`
	ApplyAt       time.Time `json:"applyAt" comment:"申购日期"`
	Status        int8      `json:"status" comment:"申购状态(1=待审批, 2=已审批, 3=已驳回, 4=已取消)"`
	Approver      int       `json:"approver" comment:"审批人编码"`
	ApproveAt     time.Time `json:"approveAt" comment:"审批时间"`
	Remark        string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetPurchaseApplyInsertReq) Generate(model *models.AssetPurchaseApply) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ApplyCode = s.ApplyCode
	model.CategoryId = s.CategoryId
	model.SupplierId = s.SupplierId
	model.ApplyUser = s.ApplyUser
	model.Quantity = s.Quantity
	model.Specification = s.Specification
	model.Brand = s.Brand
	model.Unit = s.Unit
	model.UnitPrice = s.UnitPrice
	model.TotalAmount = s.TotalAmount
	model.ApplyReason = s.ApplyReason
	model.ApplyAt = s.ApplyAt
	model.Status = s.Status
	model.Approver = s.Approver
	model.ApproveAt = s.ApproveAt
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetPurchaseApplyInsertReq) GetId() interface{} {
	return s.Id
}

type AssetPurchaseApplyUpdateReq struct {
	Id            int       `uri:"id" comment:"主键"` // 主键
	ApplyCode     string    `json:"applyCode" comment:"申请单编号"`
	CategoryId    int       `json:"categoryId" comment:"资产类型编码"`
	SupplierId    int       `json:"supplierId" comment:"供应商编码"`
	ApplyUser     int       `json:"applyUser" comment:"申购人编码"`
	Quantity      int64     `json:"quantity" comment:"申购数量"`
	Specification string    `json:"specification" comment:"规格型号"`
	Brand         string    `json:"brand" comment:"品牌"`
	Unit          string    `json:"unit" comment:"计量单位"`
	UnitPrice     float64   `json:"unitPrice" comment:"预估单价"`
	TotalAmount   float64   `json:"totalAmount" comment:"预估金额"`
	ApplyReason   string    `json:"applyReason" comment:"申购理由"`
	ApplyAt       time.Time `json:"applyAt" comment:"申购日期"`
	Status        int8      `json:"status" comment:"申购状态(1=待审批, 2=已审批, 3=已驳回, 4=已取消)"`
	Approver      int       `json:"approver" comment:"审批人编码"`
	ApproveAt     time.Time `json:"approveAt" comment:"审批时间"`
	Remark        string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetPurchaseApplyUpdateReq) Generate(model *models.AssetPurchaseApply) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ApplyCode = s.ApplyCode
	model.CategoryId = s.CategoryId
	model.SupplierId = s.SupplierId
	model.ApplyUser = s.ApplyUser
	model.Quantity = s.Quantity
	model.Specification = s.Specification
	model.Brand = s.Brand
	model.Unit = s.Unit
	model.UnitPrice = s.UnitPrice
	model.TotalAmount = s.TotalAmount
	model.ApplyReason = s.ApplyReason
	model.ApplyAt = s.ApplyAt
	model.Status = s.Status
	model.Approver = s.Approver
	model.ApproveAt = s.ApproveAt
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetPurchaseApplyUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetPurchaseApplyGetReq 功能获取请求参数
type AssetPurchaseApplyGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetPurchaseApplyGetReq) GetId() interface{} {
	return s.Id
}

// AssetPurchaseApplyDeleteReq 功能删除请求参数
type AssetPurchaseApplyDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetPurchaseApplyDeleteReq) GetId() interface{} {
	return s.Ids
}
