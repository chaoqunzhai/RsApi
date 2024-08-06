package dto

import (
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsDialGetPageReq struct {
	dto.Pagination `search:"-"`
	CustomUser     int64  `form:"customUser"  search:"type:exact;column:custom_user;table:rs_dial" comment:"所属客户"`
	Isp            int64  `form:"isp"  search:"type:exact;column:isp;table:rs_dial" comment:"运营商"`
	Charging       int64  `form:"charging"  search:"type:exact;column:charging;table:rs_dial" comment:"计费方式"`
	BandwidthType  int64  `form:"bandwidthType"  search:"type:exact;column:bandwidth_type;table:rs_dial" comment:"宽带类型"`
	TransProvince  int64  `form:"transProvince"  search:"type:exact;column:trans_province;table:rs_dial" comment:"是否跨省"`
	MoreDialing    int64  `form:"moreDialing"  search:"type:exact;column:more_dialing;table:rs_dial" comment:"是否支持多拨"`
	Account        string `form:"account"  search:"type:contains;column:account;table:rs_dial" comment:"账号"`
	DialName       string `form:"dialName"  search:"type:contains;column:dial_name;table:rs_dial" comment:"线路名称"`
	Status         int64  `form:"status"  search:"type:exact;column:status;table:rs_dial" comment:"拨号状态,1:正常 非1:异常"`
	Source         int64  `form:"source"  search:"type:exact;column:source;table:rs_dial" comment:"拨号状态,0:录入 1:自动上报"`
	IdcId          int64  `form:"idcId"  search:"type:exact;column:idc_id;table:rs_dial" comment:"关联的IDC"`
	HostId         int64  `form:"hostId"  search:"type:exact;column:host_id;table:rs_dial" comment:"关联主机ID"`
	DeviceId       int64  `form:"deviceId"  search:"type:exact;column:device_id;table:rs_dial" comment:"关联网卡ID"`
	RsDialOrder
}

type RsDialOrder struct {
	Id              string `form:"idOrder"  search:"type:order;column:id;table:rs_dial"`
	CreateBy        string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_dial"`
	UpdateBy        string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_dial"`
	CreatedAt       string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_dial"`
	UpdatedAt       string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_dial"`
	DeletedAt       string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_dial"`
	Desc            string `form:"descOrder"  search:"type:order;column:desc;table:rs_dial"`
	CustomUser      string `form:"customUserOrder"  search:"type:order;column:custom_user;table:rs_dial"`
	Isp             string `form:"ispOrder"  search:"type:order;column:isp;table:rs_dial"`
	Up              string `form:"upOrder"  search:"type:order;column:up;table:rs_dial"`
	Down            string `form:"downOrder"  search:"type:order;column:down;table:rs_dial"`
	Charging        string `form:"chargingOrder"  search:"type:order;column:charging;table:rs_dial"`
	Price           string `form:"priceOrder"  search:"type:order;column:price;table:rs_dial"`
	ManagerLineCost string `form:"managerLineCostOrder"  search:"type:order;column:manager_line_cost;table:rs_dial"`
	BandwidthType   string `form:"bandwidthTypeOrder"  search:"type:order;column:bandwidth_type;table:rs_dial"`
	TransProvince   string `form:"transProvinceOrder"  search:"type:order;column:trans_province;table:rs_dial"`
	MoreDialing     string `form:"moreDialingOrder"  search:"type:order;column:more_dialing;table:rs_dial"`
	Account         string `form:"accountOrder"  search:"type:order;column:account;table:rs_dial"`
	Pass            string `form:"passOrder"  search:"type:order;column:pass;table:rs_dial"`
	DialName        string `form:"dialNameOrder"  search:"type:order;column:dial_name;table:rs_dial"`
	Status          string `form:"statusOrder"  search:"type:order;column:status;table:rs_dial"`
	Source          string `form:"sourceOrder"  search:"type:order;column:source;table:rs_dial"`
	IdcId           string `form:"idcIdOrder"  search:"type:order;column:idc_id;table:rs_dial"`
	HostId          string `form:"hostIdOrder"  search:"type:order;column:host_id;table:rs_dial"`
	DeviceId        string `form:"deviceIdOrder"  search:"type:order;column:device_id;table:rs_dial"`
}

func (m *RsDialGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsDialInsertReq struct {
	Id              int     `json:"-" comment:"主键编码"` // 主键编码
	Desc            string  `json:"desc" comment:"描述信息"`
	CustomUser      int64   `json:"customUser" comment:"所属客户"`
	Isp             int64   `json:"isp" comment:"运营商"`
	Up              string  `json:"up" comment:"上行带宽"`
	Down            string  `json:"down" comment:"下行带宽"`
	Charging        int64   `json:"charging" comment:"计费方式"`
	Price           float64 `json:"price" comment:"业务线单价"`
	ManagerLineCost float64 `json:"managerLineCost" comment:"管理线价格"`
	BandwidthType   int64   `json:"bandwidthType" comment:"宽带类型"`
	TransProvince   int64   `json:"transProvince" comment:"是否跨省"`
	MoreDialing     int64   `json:"moreDialing" comment:"是否支持多拨"`
	IsManager       int64   `json:"isManager" comment:"是否管理线"`
	Account         string  `json:"account" comment:"账号"`
	Pass            string  `json:"pass" comment:"密码"`
	DialName        string  `json:"dialName" comment:"线路名称"`
	Status          int64   `json:"status" comment:"拨号状态,1:正常 非1:异常"`
	Source          int64   `json:"source" comment:"拨号状态,0:录入 1:自动上报"`
	IdcId           int64   `json:"idcId" comment:"关联的IDC"`
	HostId          int64   `json:"hostId" comment:"关联主机ID"`
	DeviceId        int64   `json:"deviceId" comment:"关联网卡ID"`
	common.ControlBy
}

func (s *RsDialInsertReq) Generate(model *models.RsDial) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc
	model.CustomUser = s.CustomUser
	model.Isp = s.Isp
	model.Up = s.Up
	model.Down = s.Down
	model.Charging = s.Charging
	model.Price = s.Price
	model.ManagerLineCost = s.ManagerLineCost
	model.BandwidthType = s.BandwidthType
	model.TransProvince = s.TransProvince
	model.MoreDialing = s.MoreDialing
	model.IsManager = s.IsManager
	model.Account = s.Account
	model.Pass = s.Pass
	model.DialName = s.DialName
	model.Status = s.Status
	model.Source = s.Source
	model.IdcId = s.IdcId
	model.HostId = s.HostId
	model.DeviceId = s.DeviceId
}

func (s *RsDialInsertReq) GetId() interface{} {
	return s.Id
}

type RsDialUpdateReq struct {
	Id              int     `uri:"id" comment:"主键编码"` // 主键编码
	Desc            string  `json:"desc" comment:"描述信息"`
	CustomUser      int64   `json:"customUser" comment:"所属客户"`
	Isp             int64   `json:"isp" comment:"运营商"`
	Up              string  `json:"up" comment:"上行带宽"`
	Down            string  `json:"down" comment:"下行带宽"`
	Charging        int64   `json:"charging" comment:"计费方式"`
	Price           float64 `json:"price" comment:"业务线单价"`
	IsManager       int64   `json:"isManager" comment:"是否管理线"`
	ManagerLineCost float64 `json:"managerLineCost" comment:"管理线价格"`
	BandwidthType   int64   `json:"bandwidthType" comment:"宽带类型"`
	TransProvince   int64   `json:"transProvince" comment:"是否跨省"`
	MoreDialing     int64   `json:"moreDialing" comment:"是否支持多拨"`
	Account         string  `json:"account" comment:"账号"`
	Pass            string  `json:"pass" comment:"密码"`
	DialName        string  `json:"dialName" comment:"线路名称"`
	Status          int64   `json:"status" comment:"拨号状态,1:正常 非1:异常"`
	Source          int64   `json:"source" comment:"拨号状态,0:录入 1:自动上报"`
	IdcId           int64   `json:"idcId" comment:"关联的IDC"`
	HostId          int64   `json:"hostId" comment:"关联主机ID"`
	DeviceId        int64   `json:"deviceId" comment:"关联网卡ID"`
	common.ControlBy
}

func (s *RsDialUpdateReq) Generate(model *models.RsDial) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.CustomUser = s.CustomUser
	model.Isp = s.Isp
	model.Up = s.Up
	model.Down = s.Down
	model.Charging = s.Charging
	model.Price = s.Price
	model.IsManager = s.IsManager
	model.ManagerLineCost = s.ManagerLineCost
	model.BandwidthType = s.BandwidthType
	model.TransProvince = s.TransProvince
	model.MoreDialing = s.MoreDialing
	model.Account = s.Account
	model.Pass = s.Pass
	model.DialName = s.DialName
	model.Status = s.Status
	model.Source = s.Source
	model.IdcId = s.IdcId
	model.HostId = s.HostId
	model.DeviceId = s.DeviceId
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
