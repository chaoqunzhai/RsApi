package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetOutboundMemberGetPageReq struct {
	dto.Pagination    `search:"-"`
	AssetOutboundId   int    `form:"assetOutboundId"  search:"type:exact;column:asset_outbound_id;table:asset_outbound_member" comment:"资产出库编码"`
	AssetOutboundCode string `form:"assetOutboundCode"  search:"type:exact;column:asset_outbound_code;table:asset_outbound_member" comment:"资产出库单号"`
	AssetId           int    `form:"assetId"  search:"type:exact;column:asset_id;table:asset_outbound_member" comment:"资产编码"`
	AssetOutboundMemberOrder
}

type AssetOutboundMemberOrder struct {
	Id                string `form:"idOrder"  search:"type:order;column:id;table:asset_outbound_member"`
	AssetOutboundId   string `form:"assetOutboundIdOrder"  search:"type:order;column:asset_outbound_id;table:asset_outbound_member"`
	AssetOutboundCode string `form:"assetOutboundCodeOrder"  search:"type:order;column:asset_outbound_code;table:asset_outbound_member"`
	AssetId           string `form:"assetIdOrder"  search:"type:order;column:asset_id;table:asset_outbound_member"`
	CreatedAt         string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_outbound_member"`
	UpdatedAt         string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_outbound_member"`
	DeletedAt         string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_outbound_member"`
	CreateBy          string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_outbound_member"`
	UpdateBy          string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_outbound_member"`
}

func (m *AssetOutboundMemberGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetOutboundMemberInsertReq struct {
	Id                int    `json:"-" comment:"主键"` // 主键
	AssetOutboundId   int    `json:"assetOutboundId" comment:"资产出库编码"`
	AssetOutboundCode string `json:"assetOutboundCode" comment:"资产出库单号"`
	AssetId           int    `json:"assetId" comment:"资产编码"`
	common.ControlBy
}

func (s *AssetOutboundMemberInsertReq) Generate(model *models.AssetOutboundMember) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetOutboundId = s.AssetOutboundId
	model.AssetOutboundCode = s.AssetOutboundCode
	model.AssetId = s.AssetId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetOutboundMemberInsertReq) GetId() interface{} {
	return s.Id
}

type AssetOutboundMemberUpdateReq struct {
	Id                int    `uri:"id" comment:"主键"` // 主键
	AssetOutboundId   int    `json:"assetOutboundId" comment:"资产出库编码"`
	AssetOutboundCode string `json:"assetOutboundCode" comment:"资产出库单号"`
	AssetId           int    `json:"assetId" comment:"资产编码"`
	common.ControlBy
}

func (s *AssetOutboundMemberUpdateReq) Generate(model *models.AssetOutboundMember) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetOutboundId = s.AssetOutboundId
	model.AssetOutboundCode = s.AssetOutboundCode
	model.AssetId = s.AssetId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetOutboundMemberUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetOutboundMemberGetReq 功能获取请求参数
type AssetOutboundMemberGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetOutboundMemberGetReq) GetId() interface{} {
	return s.Id
}

// AssetOutboundMemberDeleteReq 功能删除请求参数
type AssetOutboundMemberDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetOutboundMemberDeleteReq) GetId() interface{} {
	return s.Ids
}
