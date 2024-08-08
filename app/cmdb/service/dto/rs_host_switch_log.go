package dto

import (
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsHostSwitchLogGetPageReq struct {
	dto.Pagination `search:"-"`
	HostId         string `form:"hostId"  search:"type:exact;column:host_id;table:rs_host_switch_log" comment:"切换的主机ID"`
	BusinessId     string `form:"businessId"  search:"type:exact;column:business_id;table:rs_host_switch_log" comment:"切换的新业务ID"`
	BusinessSn     string `form:"businessSn"  search:"type:contains;column:business_sn;table:rs_host_switch_log" comment:"原来的业务SN"`
	RsHostSwitchLogOrder
}

type RsHostSwitchLogOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:rs_host_switch_log"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_host_switch_log"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_host_switch_log"`
	JobId      string `form:"jobIdOrder"  search:"type:order;column:job_id;table:rs_host_switch_log"`
	HostId     string `form:"hostIdOrder"  search:"type:order;column:host_id;table:rs_host_switch_log"`
	BusinessId string `form:"businessIdOrder"  search:"type:order;column:business_id;table:rs_host_switch_log"`
	BusinessSn string `form:"businessSnOrder"  search:"type:order;column:business_sn;table:rs_host_switch_log"`
	Desc       string `form:"descOrder"  search:"type:order;column:desc;table:rs_host_switch_log"`
}

func (m *RsHostSwitchLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsHostSwitchLogInsertReq struct {
	Id         int    `json:"-" comment:"主键编码"` // 主键编码
	JobId      string `json:"jobId" comment:"任务ID"`
	HostId     string `json:"hostId" comment:"切换的主机ID"`
	BusinessId int    `json:"businessId" comment:"切换的新业务ID"`
	BusinessSn string `json:"businessSn" comment:"原来的业务SN"`
	Desc       string `json:"desc" comment:"切换业务备注"`
	common.ControlBy
}

func (s *RsHostSwitchLogInsertReq) Generate(model *models.RsHostSwitchLog) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.JobId = s.JobId
	model.HostId = s.HostId
	model.BusinessId = s.BusinessId
	model.BusinessSn = s.BusinessSn
	model.Desc = s.Desc
}

func (s *RsHostSwitchLogInsertReq) GetId() interface{} {
	return s.Id
}

type RsHostSwitchLogUpdateReq struct {
	Id         int    `uri:"id" comment:"主键编码"` // 主键编码
	JobId      string `json:"jobId" comment:"任务ID"`
	HostId     string `json:"hostId" comment:"切换的主机ID"`
	BusinessId int    `json:"businessId" comment:"切换的新业务ID"`
	BusinessSn string `json:"businessSn" comment:"原来的业务SN"`
	Desc       string `json:"desc" comment:"切换业务备注"`
	common.ControlBy
}

func (s *RsHostSwitchLogUpdateReq) Generate(model *models.RsHostSwitchLog) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.JobId = s.JobId
	model.HostId = s.HostId
	model.BusinessId = s.BusinessId
	model.BusinessSn = s.BusinessSn
	model.Desc = s.Desc
}

func (s *RsHostSwitchLogUpdateReq) GetId() interface{} {
	return s.Id
}

// RsHostSwitchLogGetReq 功能获取请求参数
type RsHostSwitchLogGetReq struct {
	Id int `uri:"id"`
}

func (s *RsHostSwitchLogGetReq) GetId() interface{} {
	return s.Id
}

// RsHostSwitchLogDeleteReq 功能删除请求参数
type RsHostSwitchLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsHostSwitchLogDeleteReq) GetId() interface{} {
	return s.Ids
}
