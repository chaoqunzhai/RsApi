package dto

import (
	"time"

	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetOutboundGetPageReq struct {
	dto.Pagination `search:"-"`
	OutboundCode   string `form:"outboundCode"  search:"type:exact;column:outbound_code;table:asset_outbound" comment:"出库单号"`
	WarehouseId    int    `form:"warehouseId"  search:"type:exact;column:warehouse_id;table:asset_outbound" comment:"库房编码"`
	OutboundTo     int    `form:"outboundTo"  search:"type:exact;column:outbound_to;table:asset_outbound" comment:"出库去向(客户编码)"`
	OutboundBy     int    `form:"outboundBy"  search:"type:exact;column:outbound_by;table:asset_outbound" comment:"出库人编码"`
	AssetOutboundOrder
}

type AssetOutboundOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:asset_outbound"`
	OutboundCode string `form:"outboundCodeOrder"  search:"type:order;column:outbound_code;table:asset_outbound"`
	WarehouseId  string `form:"warehouseIdOrder"  search:"type:order;column:warehouse_id;table:asset_outbound"`
	OutboundTo   string `form:"outboundToOrder"  search:"type:order;column:outbound_to;table:asset_outbound"`
	OutboundBy   string `form:"outboundByOrder"  search:"type:order;column:outbound_by;table:asset_outbound"`
	OutboundAt   string `form:"outboundAtOrder"  search:"type:order;column:outbound_at;table:asset_outbound"`
	Remark       string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_outbound"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_outbound"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_outbound"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_outbound"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_outbound"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_outbound"`
}

func (m *AssetOutboundGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetOutboundInsertReq struct {
	Id           int       `json:"-" comment:"主键"` // 主键
	OutboundCode string    `json:"outboundCode" comment:"出库单号"`
	WarehouseId  int       `json:"warehouseId" comment:"库房编码"`
	OutboundTo   int       `json:"outboundTo" comment:"出库去向(客户编码)"`
	OutboundBy   int       `json:"outboundBy" comment:"出库人编码"`
	OutboundAt   time.Time `json:"outboundAt" comment:"出库时间"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetOutboundInsertReq) Generate(model *models.AssetOutbound) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.OutboundCode = s.OutboundCode
	model.WarehouseId = s.WarehouseId
	model.OutboundTo = s.OutboundTo
	model.OutboundBy = s.OutboundBy
	model.OutboundAt = s.OutboundAt
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetOutboundInsertReq) GetId() interface{} {
	return s.Id
}

type AssetOutboundUpdateReq struct {
	Id           int       `uri:"id" comment:"主键"` // 主键
	OutboundCode string    `json:"outboundCode" comment:"出库单号"`
	WarehouseId  int       `json:"warehouseId" comment:"库房编码"`
	OutboundTo   int       `json:"outboundTo" comment:"出库去向(客户编码)"`
	OutboundBy   int       `json:"outboundBy" comment:"出库人编码"`
	OutboundAt   time.Time `json:"outboundAt" comment:"出库时间"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetOutboundUpdateReq) Generate(model *models.AssetOutbound) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.OutboundCode = s.OutboundCode
	model.WarehouseId = s.WarehouseId
	model.OutboundTo = s.OutboundTo
	model.OutboundBy = s.OutboundBy
	model.OutboundAt = s.OutboundAt
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetOutboundUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetOutboundGetReq 功能获取请求参数
type AssetOutboundGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetOutboundGetReq) GetId() interface{} {
	return s.Id
}

// AssetOutboundDeleteReq 功能删除请求参数
type AssetOutboundDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetOutboundDeleteReq) GetId() interface{} {
	return s.Ids
}
