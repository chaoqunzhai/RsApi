package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetRecordingGetPageReq struct {
	dto.Pagination `search:"-"`
	AssetId        string `form:"assetId"  search:"type:exact;column:asset_id;table:asset_recording" comment:"关联资产ID"`
	User           string `form:"user"  search:"type:exact;column:user;table:asset_recording" comment:"操作人"`
	Type           string `form:"type"  search:"type:exact;column:type;table:asset_recording" comment:"操作类型"`
	Info           string `form:"info"  search:"type:contains;column:info;table:asset_recording" comment:"处理内容"`
	BindOrder      string `form:"bindOrder"  search:"type:exact;column:bind_order;table:asset_recording" comment:"关联单据"`
	AssetRecordingOrder
}

type AssetRecordingOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:asset_recording"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_recording"`
	AssetId   string `form:"assetIdOrder"  search:"type:order;column:asset_id;table:asset_recording"`
	User      string `form:"userOrder"  search:"type:order;column:user;table:asset_recording"`
	Type      string `form:"typeOrder"  search:"type:order;column:type;table:asset_recording"`
	Info      string `form:"infoOrder"  search:"type:order;column:info;table:asset_recording"`
	BindOrder string `form:"bindOrderOrder"  search:"type:order;column:bind_order;table:asset_recording"`
}

func (m *AssetRecordingGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetRecordingInsertReq struct {
	Id        int    `json:"-" comment:"主键编码"` // 主键编码
	AssetId   int    `json:"assetId" comment:"关联资产ID"`
	User      string `json:"user" comment:"操作人"`
	Type      int    `json:"type" comment:"操作类型"`
	Info      string `json:"info" comment:"处理内容"`
	BindOrder string `json:"bindOrder" comment:"关联单据"`
	common.ControlBy
}

func (s *AssetRecordingInsertReq) Generate(model *models.AssetRecording) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.User = s.User
	model.Type = s.Type
	model.Info = s.Info
	model.BindOrder = s.BindOrder
}

func (s *AssetRecordingInsertReq) GetId() interface{} {
	return s.Id
}

type AssetRecordingUpdateReq struct {
	Id        int    `uri:"id" comment:"主键编码"` // 主键编码
	AssetId   int    `json:"assetId" comment:"关联资产ID"`
	User      string `json:"user" comment:"操作人"`
	Type      int    `json:"type" comment:"操作类型"`
	Info      string `json:"info" comment:"处理内容"`
	BindOrder string `json:"bindOrder" comment:"关联单据"`
	common.ControlBy
}

func (s *AssetRecordingUpdateReq) Generate(model *models.AssetRecording) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.User = s.User
	model.Type = s.Type
	model.Info = s.Info
	model.BindOrder = s.BindOrder
}

func (s *AssetRecordingUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetRecordingGetReq 功能获取请求参数
type AssetRecordingGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetRecordingGetReq) GetId() interface{} {
	return s.Id
}

// AssetRecordingDeleteReq 功能删除请求参数
type AssetRecordingDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetRecordingDeleteReq) GetId() interface{} {
	return s.Ids
}
