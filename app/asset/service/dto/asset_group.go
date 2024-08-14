package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetGroupGetPageReq struct {
	dto.Pagination `search:"-"`
	GroupName      string `form:"groupName"  search:"type:exact;column:group_name;table:asset_group" comment:"资产组合名称"`
	MainAssetId    string `form:"mainAssetId"  search:"type:exact;column:main_asset_id;table:asset_group" comment:"主资产编码"`
	AssetGroupOrder
}

type AssetGroupOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:asset_group"`
	GroupName   string `form:"groupNameOrder"  search:"type:order;column:group_name;table:asset_group"`
	MainAssetId string `form:"mainAssetIdOrder"  search:"type:order;column:main_asset_id;table:asset_group"`
	Remark      string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_group"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_group"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_group"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_group"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_group"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_group"`
}

func (m *AssetGroupGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetGroupInsertReq struct {
	Id          int    `json:"-" comment:"主键"` // 主键
	GroupName   string `json:"groupName" comment:"资产组合名称"`
	MainAssetId int    `json:"mainAssetId" comment:"主资产编码"`
	Remark      string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetGroupInsertReq) Generate(model *models.AssetGroup) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GroupName = s.GroupName
	model.MainAssetId = s.MainAssetId
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetGroupInsertReq) GetId() interface{} {
	return s.Id
}

type AssetGroupUpdateReq struct {
	Id          int    `uri:"id" comment:"主键"` // 主键
	GroupName   string `json:"groupName" comment:"资产组合名称"`
	MainAssetId int    `json:"mainAssetId" comment:"主资产编码"`
	Remark      string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetGroupUpdateReq) Generate(model *models.AssetGroup) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GroupName = s.GroupName
	model.MainAssetId = s.MainAssetId
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetGroupUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetGroupGetReq 功能获取请求参数
type AssetGroupGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetGroupGetReq) GetId() interface{} {
	return s.Id
}

// AssetGroupDeleteReq 功能删除请求参数
type AssetGroupDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetGroupDeleteReq) GetId() interface{} {
	return s.Ids
}
