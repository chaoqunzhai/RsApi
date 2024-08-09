package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetPurchaseGetPageReq struct {
	dto.Pagination `search:"-"`
	PurchaseCode   string `form:"purchaseCode"  search:"type:exact;column:purchase_code;table:asset_purchase" comment:"采购单编号"`
	CategoryId     string `form:"categoryId"  search:"type:exact;column:category_id;table:asset_purchase" comment:"资产类型编码"`
	SupplierId     string `form:"supplierId"  search:"type:exact;column:supplier_id;table:asset_purchase" comment:"供应商编码"`
	PurchaseUser   string `form:"purchaseUser"  search:"type:exact;column:purchase_user;table:asset_purchase" comment:"采购人编码"`
	Specification  string `form:"specification"  search:"type:exact;column:specification;table:asset_purchase" comment:"规格型号"`
	Brand          string `form:"brand"  search:"type:exact;column:brand;table:asset_purchase" comment:"品牌"`
	PurchaseAt     string `form:"purchaseAt"  search:"type:exact;column:purchase_at;table:asset_purchase" comment:"采购日期"`
	Remark         string `form:"remark"  search:"type:exact;column:remark;table:asset_purchase" comment:"备注"`
	AssetPurchaseOrder
}

type AssetPurchaseOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:asset_purchase"`
	PurchaseCode  string `form:"purchaseCodeOrder"  search:"type:order;column:purchase_code;table:asset_purchase"`
	CategoryId    string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:asset_purchase"`
	SupplierId    string `form:"supplierIdOrder"  search:"type:order;column:supplier_id;table:asset_purchase"`
	PurchaseUser  string `form:"purchaseUserOrder"  search:"type:order;column:purchase_user;table:asset_purchase"`
	Specification string `form:"specificationOrder"  search:"type:order;column:specification;table:asset_purchase"`
	Brand         string `form:"brandOrder"  search:"type:order;column:brand;table:asset_purchase"`
	Quantity      string `form:"quantityOrder"  search:"type:order;column:quantity;table:asset_purchase"`
	Unit          string `form:"unitOrder"  search:"type:order;column:unit;table:asset_purchase"`
	UnitPrice     string `form:"unitPriceOrder"  search:"type:order;column:unit_price;table:asset_purchase"`
	TotalAmount   string `form:"totalAmountOrder"  search:"type:order;column:total_amount;table:asset_purchase"`
	PurchaseAt    string `form:"purchaseAtOrder"  search:"type:order;column:purchase_at;table:asset_purchase"`
	Remark        string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_purchase"`
	Attachment    string `form:"attachmentOrder"  search:"type:order;column:attachment;table:asset_purchase"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_purchase"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_purchase"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_purchase"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_purchase"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_purchase"`
}

func (m *AssetPurchaseGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetPurchaseInsertReq struct {
	Id            int    `json:"-" comment:"主键"` // 主键
	PurchaseCode  string `json:"purchaseCode" comment:"采购单编号"`
	CategoryId    string `json:"categoryId" comment:"资产类型编码"`
	SupplierId    string `json:"supplierId" comment:"供应商编码"`
	PurchaseUser  string `json:"purchaseUser" comment:"采购人编码"`
	Specification string `json:"specification" comment:"规格型号"`
	Brand         string `json:"brand" comment:"品牌"`
	Quantity      string `json:"quantity" comment:"采购数量"`
	Unit          string `json:"unit" comment:"计量单位"`
	UnitPrice     string `json:"unitPrice" comment:"采购单价"`
	TotalAmount   string `json:"totalAmount" comment:"采购金额"`
	PurchaseAt    string `json:"purchaseAt" comment:"采购日期"`
	Remark        string `json:"remark" comment:"备注"`
	Attachment    string `json:"attachment" comment:"附件"`
	common.ControlBy
}

func (s *AssetPurchaseInsertReq) Generate(model *models.AssetPurchase) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PurchaseCode = s.PurchaseCode
	model.CategoryId = s.CategoryId
	model.SupplierId = s.SupplierId
	model.PurchaseUser = s.PurchaseUser
	model.Specification = s.Specification
	model.Brand = s.Brand
	model.Quantity = s.Quantity
	model.Unit = s.Unit
	model.UnitPrice = s.UnitPrice
	model.TotalAmount = s.TotalAmount
	model.PurchaseAt = s.PurchaseAt
	model.Remark = s.Remark
	model.Attachment = s.Attachment
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetPurchaseInsertReq) GetId() interface{} {
	return s.Id
}

type AssetPurchaseUpdateReq struct {
	Id            int    `uri:"id" comment:"主键"` // 主键
	PurchaseCode  string `json:"purchaseCode" comment:"采购单编号"`
	CategoryId    string `json:"categoryId" comment:"资产类型编码"`
	SupplierId    string `json:"supplierId" comment:"供应商编码"`
	PurchaseUser  string `json:"purchaseUser" comment:"采购人编码"`
	Specification string `json:"specification" comment:"规格型号"`
	Brand         string `json:"brand" comment:"品牌"`
	Quantity      string `json:"quantity" comment:"采购数量"`
	Unit          string `json:"unit" comment:"计量单位"`
	UnitPrice     string `json:"unitPrice" comment:"采购单价"`
	TotalAmount   string `json:"totalAmount" comment:"采购金额"`
	PurchaseAt    string `json:"purchaseAt" comment:"采购日期"`
	Remark        string `json:"remark" comment:"备注"`
	Attachment    string `json:"attachment" comment:"附件"`
	common.ControlBy
}

func (s *AssetPurchaseUpdateReq) Generate(model *models.AssetPurchase) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PurchaseCode = s.PurchaseCode
	model.CategoryId = s.CategoryId
	model.SupplierId = s.SupplierId
	model.PurchaseUser = s.PurchaseUser
	model.Specification = s.Specification
	model.Brand = s.Brand
	model.Quantity = s.Quantity
	model.Unit = s.Unit
	model.UnitPrice = s.UnitPrice
	model.TotalAmount = s.TotalAmount
	model.PurchaseAt = s.PurchaseAt
	model.Remark = s.Remark
	model.Attachment = s.Attachment
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetPurchaseUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetPurchaseGetReq 功能获取请求参数
type AssetPurchaseGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetPurchaseGetReq) GetId() interface{} {
	return s.Id
}

// AssetPurchaseDeleteReq 功能删除请求参数
type AssetPurchaseDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetPurchaseDeleteReq) GetId() interface{} {
	return s.Ids
}
