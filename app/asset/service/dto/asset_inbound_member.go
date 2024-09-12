package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetInboundMemberGetPageReq struct {
	dto.Pagination   `search:"-"`
	AssetInboundId   int    `form:"assetInboundId"  search:"type:exact;column:asset_inbound_id;table:asset_inbound_member" comment:"资产入库编码"`
	AssetInboundCode string `form:"assetInboundCode"  search:"type:exact;column:asset_inbound_code;table:asset_inbound_member" comment:"资产入库单号"`
	AssetId          int    `form:"assetId"  search:"type:exact;column:asset_id;table:asset_inbound_member" comment:"资产编码"`
	AssetInboundMemberOrder
}

type AssetInboundMemberOrder struct {
	Id               string `form:"idOrder"  search:"type:order;column:id;table:asset_inbound_member"`
	AssetInboundId   string `form:"assetInboundIdOrder"  search:"type:order;column:asset_inbound_id;table:asset_inbound_member"`
	AssetInboundCode string `form:"assetInboundCodeOrder"  search:"type:order;column:asset_inbound_code;table:asset_inbound_member"`
	AssetId          string `form:"assetIdOrder"  search:"type:order;column:asset_id;table:asset_inbound_member"`
	CreatedAt        string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_inbound_member"`
	UpdatedAt        string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_inbound_member"`
	DeletedAt        string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_inbound_member"`
	CreateBy         string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_inbound_member"`
	UpdateBy         string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_inbound_member"`
}

func (m *AssetInboundMemberGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetInboundMemberInsertReq struct {
	Id               int    `json:"-" comment:"主键"` // 主键
	AssetInboundId   int    `json:"assetInboundId" comment:"资产入库编码"`
	AssetInboundCode string `json:"assetInboundCode" comment:"资产入库单号"`
	AssetId          int    `json:"assetId" comment:"资产编码"`
	common.ControlBy
}

func (s *AssetInboundMemberInsertReq) Generate(model *models.AssetInboundMember) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetInboundId = s.AssetInboundId
	model.AssetInboundCode = s.AssetInboundCode
	model.AssetId = s.AssetId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetInboundMemberInsertReq) GetId() interface{} {
	return s.Id
}

type AssetInboundMemberUpdateReq struct {
	Id               int    `uri:"id" comment:"主键"` // 主键
	AssetInboundId   int    `json:"assetInboundId" comment:"资产入库编码"`
	AssetInboundCode string `json:"assetInboundCode" comment:"资产入库单号"`
	AssetId          int    `json:"assetId" comment:"资产编码"`
	common.ControlBy
}

func (s *AssetInboundMemberUpdateReq) Generate(model *models.AssetInboundMember) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetInboundId = s.AssetInboundId
	model.AssetInboundCode = s.AssetInboundCode
	model.AssetId = s.AssetId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetInboundMemberUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetInboundMemberGetReq 功能获取请求参数
type AssetInboundMemberGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetInboundMemberGetReq) GetId() interface{} {
	return s.Id
}

// AssetInboundMemberDeleteReq 功能删除请求参数
type AssetInboundMemberDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetInboundMemberDeleteReq) GetId() interface{} {
	return s.Ids
}
