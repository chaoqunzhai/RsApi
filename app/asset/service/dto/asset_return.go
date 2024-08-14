package dto

import (
	"time"

	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetReturnGetPageReq struct {
	dto.Pagination `search:"-"`
	AssetId        int64     `form:"assetId"  search:"type:exact;column:asset_id;table:asset_return" comment:"资产编码"`
	ReturnPerson   int64     `form:"returnPerson"  search:"type:exact;column:return_person;table:asset_return" comment:"退还人编码"`
	Reason         string    `form:"reason"  search:"type:exact;column:reason;table:asset_return" comment:"退还原因"`
	ReturnAt       time.Time `form:"returnAt"  search:"type:exact;column:return_at;table:asset_return" comment:"退还时间"`
	AssetReturnOrder
}

type AssetReturnOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:asset_return"`
	AssetId      string `form:"assetIdOrder"  search:"type:order;column:asset_id;table:asset_return"`
	ReturnPerson string `form:"returnPersonOrder"  search:"type:order;column:return_person;table:asset_return"`
	Reason       string `form:"reasonOrder"  search:"type:order;column:reason;table:asset_return"`
	ReturnAt     string `form:"returnAtOrder"  search:"type:order;column:return_at;table:asset_return"`
	Remark       string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_return"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_return"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_return"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_return"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_return"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_return"`
}

func (m *AssetReturnGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetReturnInsertReq struct {
	Id           int       `json:"-" comment:"主键"` // 主键
	AssetId      int       `json:"assetId" comment:"资产编码"`
	ReturnPerson int       `json:"returnPerson" comment:"退还人编码"`
	Reason       string    `json:"reason" comment:"退还原因"`
	ReturnAt     time.Time `json:"returnAt" comment:"退还时间"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetReturnInsertReq) Generate(model *models.AssetReturn) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.ReturnPerson = s.ReturnPerson
	model.Reason = s.Reason
	model.ReturnAt = s.ReturnAt
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetReturnInsertReq) GetId() interface{} {
	return s.Id
}

type AssetReturnUpdateReq struct {
	Id           int       `uri:"id" comment:"主键"` // 主键
	AssetId      int       `json:"assetId" comment:"资产编码"`
	ReturnPerson int       `json:"returnPerson" comment:"退还人编码"`
	Reason       string    `json:"reason" comment:"退还原因"`
	ReturnAt     time.Time `json:"returnAt" comment:"退还时间"`
	Remark       string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetReturnUpdateReq) Generate(model *models.AssetReturn) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.ReturnPerson = s.ReturnPerson
	model.Reason = s.Reason
	model.ReturnAt = s.ReturnAt
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetReturnUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetReturnGetReq 功能获取请求参数
type AssetReturnGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetReturnGetReq) GetId() interface{} {
	return s.Id
}

// AssetReturnDeleteReq 功能删除请求参数
type AssetReturnDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetReturnDeleteReq) GetId() interface{} {
	return s.Ids
}
