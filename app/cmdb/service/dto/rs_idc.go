package dto

import (
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsIdcGetPageReq struct {
	dto.Pagination `search:"-"`
	Number         int    `form:"number"  search:"type:exact;column:number;table:rs_idc" comment:"机房编号"`
	Name           string `form:"name"  search:"type:contains;column:name;table:rs_idc" comment:"机房名称"`
	Search         string `form:"search" search:"-"`
	CustomUser     int    `form:"customUser"  search:"type:exact;column:custom_user;table:rs_idc" comment:"所属客户"`
	Region         string `form:"region"  search:"type:contains;column:region;table:rs_idc" comment:"所在区域"`
	TypeId         int    `form:"typeId"  search:"type:exact;column:type_id;table:rs_idc" comment:"机房类型"`
	BuId           int    `form:"buId"  search:"type:exact;column:bu_id;table:rs_idc" comment:"商务人员"`
	Status         int    `form:"status"  search:"type:exact;column:status;table:rs_idc" comment:"机房状态"`
	Belong         int    `form:"belong"  search:"type:exact;column:belong;table:rs_idc" comment:"机房归属"`
	RsIdcOrder
}

type RsIdcOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:rs_idc"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_idc"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_idc"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_idc"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_idc"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_idc"`
	Desc         string `form:"descOrder"  search:"type:order;column:desc;table:rs_idc"`
	Number       string `form:"numberOrder"  search:"type:order;column:number;table:rs_idc"`
	Name         string `form:"nameOrder"  search:"type:order;column:name;table:rs_idc"`
	CustomUser   string `form:"customUserOrder"  search:"type:order;column:custom_user;table:rs_idc"`
	Region       string `form:"regionOrder"  search:"type:order;column:region;table:rs_idc"`
	Address      string `form:"addressOrder"  search:"type:order;column:address;table:rs_idc"`
	IpV6         string `form:"ipV6Order"  search:"type:order;column:ip_v6;table:rs_idc"`
	TypeId       string `form:"typeIdOrder"  search:"type:order;column:type_id;table:rs_idc"`
	BusinessUser string `form:"businessUserOrder"  search:"type:order;column:business_user;table:rs_idc"`
	WechatName   string `form:"wechatNameOrder"  search:"type:order;column:wechat_name;table:rs_idc"`
	WebHookUrl   string `form:"webHookUrlOrder"  search:"type:order;column:web_hook_url;table:rs_idc"`
	Status       string `form:"statusOrder"  search:"type:order;column:status;table:rs_idc"`
	Belong       string `form:"belongOrder"  search:"type:order;column:belong;table:rs_idc"`
}

func (m *RsIdcGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsIdcInsertReq struct {
	Id     int    `json:"-" comment:"主键编码"` // 主键编码
	Desc   string `json:"desc" comment:"描述信息"`
	Number int    `json:"number" comment:"机房编号"`
	Name   string `json:"name" comment:"机房名称"`

	Region        string `json:"region" comment:"所在地区"`
	Address       string `json:"address" comment:"详细地址"`
	IpV6          int    `json:"ipV6" comment:"是否IPV6"`
	TypeId        int    `json:"typeId" comment:"机房类型"`
	BuId          int    `json:"buId" comment:"商务人员"`
	CustomId      int    `json:"customId" comment:"所属客户ID"`
	WechatName    string `json:"wechatName" comment:"企业微信群名称"`
	WebHookUrl    string `json:"webHookUrl" comment:"企业微信webhookUrl"`
	Status        int    `json:"status" comment:"机房状态"`
	Belong        int    `json:"belong" comment:"机房归属"`
	TransProvince int    `json:"transProvince" comment:"跨省"`
	MoreDialing   int    `json:"moreDialing" comment:"多拨"`
	common.ControlBy
}

func (s *RsIdcInsertReq) Generate(model *models.RsIdc) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc
	model.Number = s.Number
	model.Name = s.Name
	model.CustomId = s.CustomId
	model.Region = s.Region
	model.Address = s.Address
	model.IpV6 = s.IpV6
	model.TypeId = s.TypeId
	model.BuId = s.BuId
	model.WechatName = s.WechatName
	model.WebHookUrl = s.WebHookUrl
	model.Status = s.Status
	model.Belong = s.Belong
	model.TransProvince = s.TransProvince
}

func (s *RsIdcInsertReq) GetId() interface{} {
	return s.Id
}

type RsIdcUpdateReq struct {
	Id            int    `uri:"id" comment:"主键编码"` // 主键编码
	Desc          string `json:"desc" comment:"描述信息"`
	Number        int    `json:"number" comment:"机房编号"`
	Name          string `json:"name" comment:"机房名称"`
	Region        string `json:"region" comment:"所在地区"`
	Address       string `json:"address" comment:"详细地址"`
	IpV6          int    `json:"ipV6" comment:"是否IPV6"`
	TypeId        int    `json:"typeId" comment:"机房类型"`
	BuId          int    `json:"buId" comment:"商务人员"`
	CustomId      int    `json:"customId" comment:"所属客户ID"`
	WechatName    string `json:"wechatName" comment:"企业微信群名称"`
	WebHookUrl    string `json:"webHookUrl" comment:"企业微信webhookUrl"`
	Status        int    `json:"status" comment:"机房状态"`
	Belong        int    `json:"belong" comment:"机房归属"`
	TransProvince int    `json:"transProvince" comment:"跨省"`
	MoreDialing   int    `json:"moreDialing" comment:"多拨"`
	common.ControlBy
}

func (s *RsIdcUpdateReq) Generate(model *models.RsIdc) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.Number = s.Number
	model.Name = s.Name
	model.BuId = s.BuId
	model.Region = s.Region
	model.Address = s.Address
	model.IpV6 = s.IpV6
	model.TypeId = s.TypeId
	model.CustomId = s.CustomId
	model.WechatName = s.WechatName
	model.WebHookUrl = s.WebHookUrl
	model.Status = s.Status
	model.Belong = s.Belong
}

func (s *RsIdcUpdateReq) GetId() interface{} {
	return s.Id
}

// RsIdcGetReq 功能获取请求参数
type RsIdcGetReq struct {
	Id int `uri:"id"`
}

func (s *RsIdcGetReq) GetId() interface{} {
	return s.Id
}

// RsIdcDeleteReq 功能删除请求参数
type RsIdcDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsIdcDeleteReq) GetId() interface{} {
	return s.Ids
}
