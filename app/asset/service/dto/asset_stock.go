package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetStockGetPageReq struct {
	dto.Pagination `search:"-"`
	WarehouseId    string `form:"warehouseId"  search:"type:exact;column:warehouse_id;table:asset_stock" comment:"库房编码"`
	CategoryId     string `form:"categoryId"  search:"type:exact;column:category_id;table:asset_stock" comment:"资产类别编码"`
	AssetStockOrder
}

type AssetStockOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:asset_stock"`
	WarehouseId string `form:"warehouseIdOrder"  search:"type:order;column:warehouse_id;table:asset_stock"`
	CategoryId  string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:asset_stock"`
	Quantity    string `form:"quantityOrder"  search:"type:order;column:quantity;table:asset_stock"`
	Remark      string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_stock"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_stock"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_stock"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_stock"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_stock"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_stock"`
}

func (m *AssetStockGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetStockInsertReq struct {
	Id          int    `json:"-" comment:"主键"` // 主键
	WarehouseId string `json:"warehouseId" comment:"库房编码"`
	CategoryId  string `json:"categoryId" comment:"资产类别编码"`
	Quantity    string `json:"quantity" comment:"资产库存数量"`
	Remark      string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetStockInsertReq) Generate(model *models.AssetStock) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.WarehouseId = s.WarehouseId
	model.CategoryId = s.CategoryId
	model.Quantity = s.Quantity
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetStockInsertReq) GetId() interface{} {
	return s.Id
}

type AssetStockUpdateReq struct {
	Id          int    `uri:"id" comment:"主键"` // 主键
	WarehouseId string `json:"warehouseId" comment:"库房编码"`
	CategoryId  string `json:"categoryId" comment:"资产类别编码"`
	Quantity    string `json:"quantity" comment:"资产库存数量"`
	Remark      string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetStockUpdateReq) Generate(model *models.AssetStock) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.WarehouseId = s.WarehouseId
	model.CategoryId = s.CategoryId
	model.Quantity = s.Quantity
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetStockUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetStockGetReq 功能获取请求参数
type AssetStockGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetStockGetReq) GetId() interface{} {
	return s.Id
}

// AssetStockDeleteReq 功能删除请求参数
type AssetStockDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetStockDeleteReq) GetId() interface{} {
	return s.Ids
}
