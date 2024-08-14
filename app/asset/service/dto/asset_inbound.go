package dto

import (
	"time"

	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetInboundGetPageReq struct {
	dto.Pagination `search:"-"`
	AssetId        int    `form:"assetId"  search:"type:exact;column:asset_id;table:asset_inbound" comment:"资产编码"`
	WarehouseId    int    `form:"warehouseId"  search:"type:exact;column:warehouse_id;table:asset_inbound" comment:"库房编码"`
	InboundFrom    int    `form:"inboundFrom"  search:"type:exact;column:inbound_from;table:asset_inbound" comment:"来源(1=采购、0=直接入库)"`
	FromCode       string `form:"fromCode"  search:"type:exact;column:from_code;table:asset_inbound" comment:"来源凭证编码(采购编码)"`
	InboundBy      int    `form:"inboundBy"  search:"type:exact;column:inbound_by;table:asset_inbound" comment:"入库人编码"`
	AssetInboundOrder
}

type AssetInboundOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:asset_inbound"`
	AssetId     string `form:"assetIdOrder"  search:"type:order;column:asset_id;table:asset_inbound"`
	WarehouseId string `form:"warehouseIdOrder"  search:"type:order;column:warehouse_id;table:asset_inbound"`
	InboundFrom string `form:"inboundFromOrder"  search:"type:order;column:inbound_from;table:asset_inbound"`
	FromCode    string `form:"fromCodeOrder"  search:"type:order;column:from_code;table:asset_inbound"`
	InboundBy   string `form:"inboundByOrder"  search:"type:order;column:inbound_by;table:asset_inbound"`
	InboundAt   string `form:"inboundAtOrder"  search:"type:order;column:inbound_at;table:asset_inbound"`
	Remark      string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_inbound"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_inbound"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_inbound"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_inbound"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_inbound"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_inbound"`
}

func (m *AssetInboundGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetInboundInsertReq struct {
	Id          int       `json:"-" comment:"主键"` // 主键
	AssetId     int       `json:"assetId" comment:"资产编码"`
	WarehouseId int       `json:"warehouseId" comment:"库房编码"`
	InboundFrom int8      `json:"inboundFrom" comment:"来源(1=采购、0=直接入库)"`
	FromCode    string    `json:"fromCode" comment:"来源凭证编号(采购单编号)"`
	InboundBy   int       `json:"inboundBy" comment:"入库人编码"`
	InboundAt   time.Time `json:"inboundAt" comment:"入库时间"`
	Remark      string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetInboundInsertReq) Generate(model *models.AssetInbound) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.WarehouseId = s.WarehouseId
	model.InboundFrom = s.InboundFrom
	model.FromCode = s.FromCode
	model.InboundBy = s.InboundBy
	model.InboundAt = s.InboundAt
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetInboundInsertReq) GetId() interface{} {
	return s.Id
}

type AssetInboundUpdateReq struct {
	Id          int       `uri:"id" comment:"主键"` // 主键
	AssetId     int       `json:"assetId" comment:"资产编码"`
	WarehouseId int       `json:"warehouseId" comment:"库房编码"`
	InboundFrom int8      `json:"inboundFrom" comment:"来源(1=采购、0=直接入库)"`
	FromCode    string    `json:"fromCode" comment:"来源凭证编码(采购编码)"`
	InboundBy   int       `json:"inboundBy" comment:"入库人编码"`
	InboundAt   time.Time `json:"inboundAt" comment:"入库时间"`
	Remark      string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetInboundUpdateReq) Generate(model *models.AssetInbound) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.WarehouseId = s.WarehouseId
	model.InboundFrom = s.InboundFrom
	model.FromCode = s.FromCode
	model.InboundBy = s.InboundBy
	model.InboundAt = s.InboundAt
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetInboundUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetInboundGetReq 功能获取请求参数
type AssetInboundGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetInboundGetReq) GetId() interface{} {
	return s.Id
}

// AssetInboundDeleteReq 功能删除请求参数
type AssetInboundDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetInboundDeleteReq) GetId() interface{} {
	return s.Ids
}
