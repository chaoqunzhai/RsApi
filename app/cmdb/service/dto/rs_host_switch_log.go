package dto

import (
	"go-admin/common/dto"
)

type RsHostSwitchLogGetPageReq struct {
	dto.Pagination `search:"-"`
	HostId         string `form:"hostId"  search:"type:exact;column:host_id;table:rs_host_switch_log" comment:"切换的主机ID"`
	BuTargetId     string `form:"buTargetId"  search:"type:exact;column:bu_target_id;table:rs_host_switch_log" comment:"切换的新业务ID"`
	BuSource       string `form:"buSource"  search:"type:contains;column:bu_source;table:rs_host_switch_log" comment:"原来的业务SN"`
	RsHostSwitchLogOrder
}

type RsHostSwitchLogOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:rs_host_switch_log"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_host_switch_log"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_host_switch_log"`
	JobId     string `form:"jobIdOrder"  search:"type:order;column:job_id;table:rs_host_switch_log"`
	HostId    string `form:"hostIdOrder"  search:"type:order;column:host_id;table:rs_host_switch_log"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:rs_host_switch_log"`
}

func (m *RsHostSwitchLogGetPageReq) GetNeedSearch() interface{} {
	return *m
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
