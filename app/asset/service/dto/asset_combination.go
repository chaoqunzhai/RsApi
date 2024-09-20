package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CombinationGetPageReq struct {
	dto.Pagination `search:"-"`
	CustomId       int    `form:"customId" search:"type:exact;column:custom_id;table:combination"`
	HostId         int    `form:"hostId" search:"type:exact;column:host_id;table:combination"`
	IdcId          int    `form:"idcId" search:"type:exact;column:idc_id;table:combination"`
	Code           string `form:"code"  search:"type:contains;column:code;table:combination" comment:"组合编号"`
	Status         string `form:"status"  search:"type:exact;column:status;table:combination" comment:"资产状态"`
	CombinationOrder
}

type CombinationOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:combination"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:combination"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:combination"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:combination"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:combination"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:combination"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:combination"`
	JobId     string `form:"jobIdOrder"  search:"type:order;column:job_id;table:combination"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:combination"`
}

func (m *CombinationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CombinationAutoReq struct {
	Sn       string     `json:"sn" comment:"服务器SN"`
	DiskSn   []AutoDisk `json:"diskSn" comment:"磁盘SN"`
	MemorySn []string   `json:"memorySn" comment:"内存条SN"`
}

type AutoDisk struct {
	Name string `json:"name"`
	Sn   string `json:"sn"`
}
type CombinationInsertReq struct {
	Id   int    `json:"-" comment:"主键编码"` // 主键编码
	Desc string `json:"desc" comment:"描述信息"`

	Status string `json:"status" comment:"资产状态"`
	Asset  []int  `json:"asset"`
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

	Status string `json:"status" comment:"资产状态"`
	Asset  []int  `json:"asset"`
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
