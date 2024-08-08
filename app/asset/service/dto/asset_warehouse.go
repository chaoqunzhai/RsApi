package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetWarehouseGetPageReq struct {
	dto.Pagination  `search:"-"`
	WarehouseName   string `form:"warehouseName"  search:"type:exact;column:warehouse_name;table:asset_warehouse" comment:"库房名称"`
	AdministratorId string `form:"administratorId"  search:"type:exact;column:administrator_id;table:asset_warehouse" comment:"管理员编码"`
	AssetWarehouseOrder
}

type AssetWarehouseOrder struct {
	Id              string `form:"idOrder"  search:"type:order;column:id;table:asset_warehouse"`
	WarehouseName   string `form:"warehouseNameOrder"  search:"type:order;column:warehouse_name;table:asset_warehouse"`
	AdministratorId string `form:"administratorIdOrder"  search:"type:order;column:administrator_id;table:asset_warehouse"`
	Remark          string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_warehouse"`
	CreatedAt       string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_warehouse"`
	UpdatedAt       string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_warehouse"`
	DeletedAt       string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_warehouse"`
	CreateBy        string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_warehouse"`
	UpdateBy        string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_warehouse"`
}

func (m *AssetWarehouseGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetWarehouseInsertReq struct {
	Id              int    `json:"-" comment:"主键"` // 主键
	WarehouseName   string `json:"warehouseName" comment:"库房名称"`
	AdministratorId string `json:"administratorId" comment:"管理员编码"`
	Remark          string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetWarehouseInsertReq) Generate(model *models.AssetWarehouse) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.WarehouseName = s.WarehouseName
	model.AdministratorId = s.AdministratorId
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetWarehouseInsertReq) GetId() interface{} {
	return s.Id
}

type AssetWarehouseUpdateReq struct {
	Id              int    `uri:"id" comment:"主键"` // 主键
	WarehouseName   string `json:"warehouseName" comment:"库房名称"`
	AdministratorId string `json:"administratorId" comment:"管理员编码"`
	Remark          string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetWarehouseUpdateReq) Generate(model *models.AssetWarehouse) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.WarehouseName = s.WarehouseName
	model.AdministratorId = s.AdministratorId
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetWarehouseUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetWarehouseGetReq 功能获取请求参数
type AssetWarehouseGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetWarehouseGetReq) GetId() interface{} {
	return s.Id
}

// AssetWarehouseDeleteReq 功能删除请求参数
type AssetWarehouseDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetWarehouseDeleteReq) GetId() interface{} {
	return s.Ids
}
