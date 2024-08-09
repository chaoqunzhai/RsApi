package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetCategoryGetPageReq struct {
	dto.Pagination `search:"-"`
	CategoryName   string `form:"categoryName"  search:"type:exact;column:category_name;table:asset_category" comment:"类别名称"`
	AssetCategoryOrder
}

type AssetCategoryOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:asset_category"`
	CategoryName string `form:"categoryNameOrder"  search:"type:order;column:category_name;table:asset_category"`
	Remark       string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_category"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_category"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_category"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_category"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_category"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_category"`
}

func (m *AssetCategoryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetCategoryInsertReq struct {
	Id           int    `json:"-" comment:"主键"` // 主键
	CategoryName string `json:"categoryName" comment:"类别名称"`
	Remark       string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetCategoryInsertReq) Generate(model *models.AssetCategory) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CategoryName = s.CategoryName
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetCategoryInsertReq) GetId() interface{} {
	return s.Id
}

type AssetCategoryUpdateReq struct {
	Id           int    `uri:"id" comment:"主键"` // 主键
	CategoryName string `json:"categoryName" comment:"类别名称"`
	Remark       string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetCategoryUpdateReq) Generate(model *models.AssetCategory) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CategoryName = s.CategoryName
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetCategoryUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetCategoryGetReq 功能获取请求参数
type AssetCategoryGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetCategoryGetReq) GetId() interface{} {
	return s.Id
}

// AssetCategoryDeleteReq 功能删除请求参数
type AssetCategoryDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetCategoryDeleteReq) GetId() interface{} {
	return s.Ids
}
