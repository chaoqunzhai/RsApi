package dto

import (
	"database/sql"
	"go-admin/global"
	"time"

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsDialGetPageReq struct {
	dto.Pagination   `search:"-"`
	Search           string `form:"search" search:"-"`
	CustomId         int    `form:"customId"  search:"type:exact;column:custom_id;table:rs_dial" comment:"所属客户"`
	ContractId       int    `form:"contractId"  search:"type:exact;column:contract_id;table:rs_dial" comment:"关联合同"`
	BroadbandType    int    `form:"broadbandType"  search:"type:exact;column:broadband_type;table:rs_dial" comment:"带宽类型,broadband_type"`
	IsManager        int    `form:"isManager"  search:"type:exact;column:is_manager;table:rs_dial" comment:"是否管理线"`
	Ip               string `form:"ip"  search:"type:contains;column:ip;table:rs_dial" comment:"IP地址"`
	DialName         string `form:"dialName"  search:"type:contains;column:dial_name;table:rs_dial" comment:"线路名称"`
	NetworkingStatus int    `form:"networkingStatus"  search:"type:exact;column:networking_status;table:rs_dial" comment:"拨号状态,1:已联网 0:未联网 -1:联网异常"`
	Status           int    `form:"status"  search:"type:exact;column:status;table:rs_dial" comment:"拨号状态,1:已拨通 0:待使用 -1:拨号异常"`
	Source           int    `form:"source"  search:"type:exact;column:source;table:rs_dial" comment:"拨号状态,0:录入 1:自动上报"`
	IdcId            int    `form:"idcId"  search:"type:exact;column:idc_id;table:rs_dial" comment:"关联的IDC"`
	HostId           int    `form:"hostId"  search:"type:exact;column:host_id;table:rs_dial" comment:"关联主机ID"`
	DeviceId         int    `form:"deviceId"  search:"type:exact;column:device_id;table:rs_dial" comment:"关联网卡ID"`
	Account          string `form:"account"  search:"type:contains;column:account;table:rs_dial" comment:"账号"`
	RsDialOrder
}

type RsDialOrder struct {
	Id               string `form:"idOrder"  search:"type:order;column:id;table:rs_dial"`
	CreateBy         string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_dial"`
	UpdateBy         string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_dial"`
	CreatedAt        string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_dial"`
	UpdatedAt        string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_dial"`
	DeletedAt        string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_dial"`
	Desc             string `form:"descOrder"  search:"type:order;column:desc;table:rs_dial"`
	CustomId         string `form:"customIdOrder"  search:"type:order;column:custom_id;table:rs_dial"`
	ContractId       string `form:"contractIdOrder"  search:"type:order;column:contract_id;table:rs_dial"`
	BroadbandType    string `form:"broadbandTypeOrder"  search:"type:order;column:broadband_type;table:rs_dial"`
	IsManager        string `form:"isManagerOrder"  search:"type:order;column:is_manager;table:rs_dial"`
	Account          string `form:"accountOrder"  search:"type:order;column:account;table:rs_dial"`
	Ip               string `form:"ipOrder"  search:"type:order;column:ip;table:rs_dial"`
	Pass             string `form:"passOrder"  search:"type:order;column:pass;table:rs_dial"`
	Mac              string `form:"macOrder"  search:"type:order;column:mac;table:rs_dial"`
	DialName         string `form:"dialNameOrder"  search:"type:order;column:dial_name;table:rs_dial"`
	NetworkingStatus string `form:"networkingStatusOrder"  search:"type:order;column:networking_status;table:rs_dial"`
	Status           string `form:"statusOrder"  search:"type:order;column:status;table:rs_dial"`
	Source           string `form:"sourceOrder"  search:"type:order;column:source;table:rs_dial"`
	IdcId            string `form:"idcIdOrder"  search:"type:order;column:idc_id;table:rs_dial"`
	HostId           string `form:"hostIdOrder"  search:"type:order;column:host_id;table:rs_dial"`
	DeviceId         string `form:"deviceIdOrder"  search:"type:order;column:device_id;table:rs_dial"`
	RunTime          string `form:"runTimeOrder"  search:"type:order;column:run_time;table:rs_dial"`
}

func (m *RsDialGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsDialInsertReq struct {
	Id               int    `json:"-" comment:"主键编码"` // 主键编码
	Desc             string `json:"desc" comment:"描述信息"`
	CustomId         int    `json:"customId" comment:"所属客户"`
	ContractId       int    `json:"contractId" comment:"关联合同"`
	BroadbandType    int    `json:"broadbandType" comment:"带宽类型,broadband_type"`
	IsManager        int    `json:"isManager" comment:"是否管理线"`
	Account          string `json:"account" comment:"账号"`
	Ip               string `json:"ip" comment:"IP地址"`
	Pass             string `json:"pass" comment:"密码"`
	Mac              string `json:"mac" comment:"MAC地址"`
	DialName         string `json:"dialName" comment:"线路名称"`
	NetworkingStatus int    `json:"networkingStatus" comment:"拨号状态,1:已联网 0:未联网 -1:联网异常"`
	Status           int    `json:"status" comment:"拨号状态,1:已拨通 0:待使用 -1:拨号异常"`
	Source           int    `json:"source" comment:"拨号状态,0:录入 1:自动上报"`
	IdcId            int    `json:"idcId" comment:"关联的IDC"`
	HostId           int    `json:"hostId" comment:"关联主机ID"`
	DeviceId         int    `json:"deviceId" comment:"关联网卡ID"`
	RunTimeAt        string `json:"runTimeAt" comment:""`
	IspId            int    `json:"ispId"`
	common.ControlBy
}

func (s *RsDialInsertReq) Generate(model *models.RsDial) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc
	model.CustomId = s.CustomId
	model.ContractId = s.ContractId
	model.BroadbandType = s.BroadbandType
	model.IsManager = s.IsManager
	model.Account = s.Account
	model.Ip = s.Ip
	model.Pass = s.Pass
	model.Mac = s.Mac
	model.DialName = s.DialName
	model.NetworkingStatus = s.NetworkingStatus
	model.Status = s.Status
	model.Source = s.Source
	model.IdcId = s.IdcId
	model.HostId = s.HostId
	model.DeviceId = s.DeviceId
	if s.RunTimeAt != "" {
		if star, err := time.ParseInLocation(time.DateTime, s.RunTimeAt, global.LOC); err == nil {
			model.RunTime = sql.NullTime{
				Time: star,
			}
		}

	}
}

func (s *RsDialInsertReq) GetId() interface{} {
	return s.Id
}

type RsDialUpdateReq struct {
	Id               int    `uri:"id" comment:"主键编码"` // 主键编码
	Desc             string `json:"desc" comment:"描述信息"`
	CustomId         int    `json:"customId" comment:"所属客户"`
	ContractId       int    `json:"contractId" comment:"关联合同"`
	BroadbandType    int    `json:"broadbandType" comment:"带宽类型,broadband_type"`
	IsManager        int    `json:"isManager" comment:"是否管理线"`
	Account          string `json:"account" comment:"账号"`
	Ip               string `json:"ip" comment:"IP地址"`
	Pass             string `json:"pass" comment:"密码"`
	Mac              string `json:"mac" comment:"MAC地址"`
	DialName         string `json:"dialName" comment:"线路名称"`
	NetworkingStatus int    `json:"networkingStatus" comment:"拨号状态,1:已联网 0:未联网 -1:联网异常"`
	Status           int    `json:"status" comment:"拨号状态,1:已拨通 0:待使用 -1:拨号异常"`
	Source           int    `json:"source" comment:"拨号状态,0:录入 1:自动上报"`
	IdcId            int    `json:"idcId" comment:"关联的IDC"`
	HostId           int    `json:"hostId" comment:"关联主机ID"`
	DeviceId         int    `json:"deviceId" comment:"关联网卡ID"`
	RunTimeAt        string `json:"RunTimeAt" comment:""`
	IspId            int    `json:"ispId"`
	common.ControlBy
}

func (s *RsDialUpdateReq) Generate(model *models.RsDial) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.CustomId = s.CustomId
	model.ContractId = s.ContractId
	model.BroadbandType = s.BroadbandType
	model.IsManager = s.IsManager
	model.Account = s.Account
	model.Ip = s.Ip
	model.Pass = s.Pass
	model.Mac = s.Mac
	model.DialName = s.DialName
	model.NetworkingStatus = s.NetworkingStatus
	model.Status = s.Status
	model.Source = s.Source
	model.IdcId = s.IdcId
	model.HostId = s.HostId
	model.DeviceId = s.DeviceId

	if s.RunTimeAt != "" {
		if star, err := time.ParseInLocation(time.DateTime, s.RunTimeAt, global.LOC); err == nil {
			model.RunTime = sql.NullTime{
				Time:  star,
				Valid: true,
			}
		}

	} else {
		model.RunTime = sql.NullTime{}
	}
	model.IspId = s.IspId
}

func (s *RsDialUpdateReq) GetId() interface{} {
	return s.Id
}

// RsDialGetReq 功能获取请求参数
type RsDialGetReq struct {
	Id int `uri:"id"`
}

func (s *RsDialGetReq) GetId() interface{} {
	return s.Id
}

// RsDialDeleteReq 功能删除请求参数
type RsDialDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsDialDeleteReq) GetId() interface{} {
	return s.Ids
}
