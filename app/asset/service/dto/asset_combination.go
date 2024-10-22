package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CombinationGetPageReq struct {
	dto.Pagination `search:"-"`
	Id             string `form:"id" search:"type:exact;column:id;table:asset_combination"`
	CustomId       int    `form:"customId" search:"type:exact;column:custom_id;table:asset_combination"`
	HostId         int    `form:"hostId" search:"type:exact;column:host_id;table:asset_combination"`
	IdcId          int    `form:"idcId" search:"type:exact;column:idc_id;table:asset_combination"`
	Code           string `form:"code"  search:"type:contains;column:code;table:asset_combination" comment:"组合编号"`
	Status         string `form:"status"  search:"type:exact;column:status;table:asset_combination" comment:"资产状态"`
	Extend         int    `form:"extend" search:"-"`
	CombinationOrder
}

type CombinationOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:asset_combination"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_combination"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_combination"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_combination"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_combination"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_combination"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:asset_combination"`
	JobId     string `form:"jobIdOrder"  search:"type:order;column:job_id;table:asset_combination"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:asset_combination"`
}

func (m *CombinationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CombinationAutoReq struct {
	Remark   string            `json:"remark"`
	Sn       string            `json:"sn"`
	Hostname string            `json:"hostname"`
	Brand    string            `json:"brand"`
	Spec     string            `json:"spec"`
	DiskSn   []AutoDisk        `json:"diskSn" comment:"磁盘SN"`
	MemorySn map[string]string `json:"memorySn" comment:"内存条SN"`
}
type AutoDisk struct {
	Name   string `json:"name"`
	Size   string `json:"size"`
	Sn     string `json:"sn"`
	Status int    `json:"status"`
}
type CombinationInsertReq struct {
	Id   int    `json:"-" comment:"主键编码"` // 主键编码
	Desc string `json:"desc" comment:"描述信息"`

	Status int   `json:"status" comment:"资产状态"`
	Asset  []int `json:"asset"`
	common.ControlBy
}

func (s *CombinationInsertReq) Generate(model *models.Combination) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc

	model.Status = s.Status
}

func (s *CombinationInsertReq) GetId() interface{} {
	return s.Id
}

type CombinationUpdateReq struct {
	Id   int    `uri:"id" comment:"主键编码"` // 主键编码
	Desc string `json:"desc" comment:"描述信息"`

	Status int   `json:"status" comment:"资产状态"`
	Asset  []int `json:"asset"`
	common.ControlBy
}

func (s *CombinationUpdateReq) Generate(model *models.Combination) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc

	model.Status = s.Status
}

func (s *CombinationUpdateReq) GetId() interface{} {
	return s.Id
}

// CombinationGetReq 功能获取请求参数
type CombinationGetReq struct {
	Id int `uri:"id"`
}

func (s *CombinationGetReq) GetId() interface{} {
	return s.Id
}

// CombinationDeleteReq 功能删除请求参数
type CombinationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CombinationDeleteReq) GetId() interface{} {
	return s.Ids
}
