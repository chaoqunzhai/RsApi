package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetGroupMemberGetPageReq struct {
	dto.Pagination `search:"-"`
	AssetGroupId   string `form:"assetGroupId"  search:"type:exact;column:asset_group_id;table:asset_group_member" comment:"资产组合编码"`
	AssetId        string `form:"assetId"  search:"type:exact;column:asset_id;table:asset_group_member" comment:"资产编码"`
	IsMain         string `form:"isMain"  search:"type:exact;column:is_main;table:asset_group_member" comment:"是否为主资产(1=是,2=否)"`
	AssetGroupMemberOrder
}

type AssetGroupMemberOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:asset_group_member"`
	AssetGroupId string `form:"assetGroupIdOrder"  search:"type:order;column:asset_group_id;table:asset_group_member"`
	AssetId      string `form:"assetIdOrder"  search:"type:order;column:asset_id;table:asset_group_member"`
	IsMain       string `form:"isMainOrder"  search:"type:order;column:is_main;table:asset_group_member"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_group_member"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_group_member"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_group_member"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_group_member"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_group_member"`
}

func (m *AssetGroupMemberGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetGroupMemberInsertReq struct {
	Id           int  `json:"-" comment:"主键"` // 主键
	AssetGroupId int  `json:"assetGroupId" comment:"资产组合编码"`
	AssetId      int  `json:"assetId" comment:"资产编码"`
	IsMain       int8 `json:"isMain" comment:"是否为主资产(1=是,2=否)"`
	common.ControlBy
}

func (s *AssetGroupMemberInsertReq) Generate(model *models.AssetGroupMember) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetGroupId = s.AssetGroupId
	model.AssetId = s.AssetId
	model.IsMain = s.IsMain
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetGroupMemberInsertReq) GetId() interface{} {
	return s.Id
}

type AssetGroupMemberUpdateReq struct {
	Id           int  `uri:"id" comment:"主键"` // 主键
	AssetGroupId int  `json:"assetGroupId" comment:"资产组合编码"`
	AssetId      int  `json:"assetId" comment:"资产编码"`
	IsMain       int8 `json:"isMain" comment:"是否为主资产(1=是,2=否)"`
	common.ControlBy
}

func (s *AssetGroupMemberUpdateReq) Generate(model *models.AssetGroupMember) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetGroupId = s.AssetGroupId
	model.AssetId = s.AssetId
	model.IsMain = s.IsMain
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetGroupMemberUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetGroupMemberGetReq 功能获取请求参数
type AssetGroupMemberGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetGroupMemberGetReq) GetId() interface{} {
	return s.Id
}

// AssetGroupMemberDeleteReq 功能删除请求参数
type AssetGroupMemberDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetGroupMemberDeleteReq) GetId() interface{} {
	return s.Ids
}
