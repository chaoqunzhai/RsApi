package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetGetPageReq struct {
	dto.Pagination `search:"-"`
	AssetCode      string `form:"assetCode"  search:"type:exact;column:asset_code;table:asset" comment:"资产编号"`
	SnCode         string `form:"snCode"  search:"type:exact;column:sn_code;table:asset" comment:"SN编码"`
	CategoryId     string `form:"categoryId"  search:"type:exact;column:category_id;table:asset" comment:"资产类别"`
	Specification  string `form:"specification"  search:"type:exact;column:specification;table:asset" comment:"规格型号"`
	Brand          string `form:"brand"  search:"type:exact;column:brand;table:asset" comment:"品牌"`
	Unit           string `form:"unit"  search:"type:exact;column:unit;table:asset" comment:"计量单位"`
	UnitPrice      string `form:"unitPrice"  search:"type:exact;column:unit_price;table:asset" comment:"单价"`
	Status         string `form:"status"  search:"type:exact;column:status;table:asset" comment:"状态(1=在库, 2=出库, 3=在用, 4=处置)"`
	AssetOrder
}

type AssetOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:asset"`
	AssetCode     string `form:"assetCodeOrder"  search:"type:order;column:asset_code;table:asset"`
	SnCode        string `form:"snCodeOrder"  search:"type:order;column:sn_code;table:asset"`
	CategoryId    string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:asset"`
	Specification string `form:"specificationOrder"  search:"type:order;column:specification;table:asset"`
	Brand         string `form:"brandOrder"  search:"type:order;column:brand;table:asset"`
	Unit          string `form:"unitOrder"  search:"type:order;column:unit;table:asset"`
	UnitPrice     string `form:"unitPriceOrder"  search:"type:order;column:unit_price;table:asset"`
	Status        string `form:"statusOrder"  search:"type:order;column:status;table:asset"`
	Remark        string `form:"remarkOrder"  search:"type:order;column:remark;table:asset"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:asset"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset"`
}

func (m *AssetGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetInsertReq struct {
	Id            int     `json:"-" comment:"主键"` // 主键
	AssetCode     string  `json:"assetCode" comment:"资产编号"`
	SnCode        string  `json:"snCode" comment:"SN编码"`
	CategoryId    int     `json:"categoryId" comment:"资产类别"`
	Specification string  `json:"specification" comment:"规格型号"`
	Brand         string  `json:"brand" comment:"品牌"`
	Unit          string  `json:"unit" comment:"计量单位"`
	UnitPrice     float64 `json:"unitPrice" comment:"单价"`
	Status        int8    `json:"status" comment:"状态(1=在库, 2=出库, 3=在用, 4=处置)"`
	Remark        string  `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetInsertReq) Generate(model *models.Asset) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetCode = s.AssetCode
	model.SnCode = s.SnCode
	model.CategoryId = s.CategoryId
	model.Specification = s.Specification
	model.Brand = s.Brand
	model.Unit = s.Unit
	model.UnitPrice = s.UnitPrice
	model.Status = s.Status
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetInsertReq) GetId() interface{} {
	return s.Id
}

type AssetUpdateReq struct {
	Id            int     `uri:"id" comment:"主键"` // 主键
	AssetCode     string  `json:"assetCode" comment:"资产编号"`
	SnCode        string  `json:"snCode" comment:"SN编码"`
	CategoryId    int     `json:"categoryId" comment:"资产类别"`
	Specification string  `json:"specification" comment:"规格型号"`
	Brand         string  `json:"brand" comment:"品牌"`
	Unit          string  `json:"unit" comment:"计量单位"`
	UnitPrice     float64 `json:"unitPrice" comment:"单价"`
	Status        int8    `json:"status" comment:"状态(1=在库, 2=出库, 3=在用, 4=处置)"`
	Remark        string  `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetUpdateReq) Generate(model *models.Asset) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetCode = s.AssetCode
	model.SnCode = s.SnCode
	model.CategoryId = s.CategoryId
	model.Specification = s.Specification
	model.Brand = s.Brand
	model.Unit = s.Unit
	model.UnitPrice = s.UnitPrice
	model.Status = s.Status
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetGetReq 功能获取请求参数
type AssetGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetGetReq) GetId() interface{} {
	return s.Id
}

// AssetDeleteReq 功能删除请求参数
type AssetDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetDeleteReq) GetId() interface{} {
	return s.Ids
}
