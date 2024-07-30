package dto

import (
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsIdcGetPageReq struct {
	dto.Pagination `search:"-"`
	Enable         string `form:"enable"  search:"type:exact;column:enable;table:rs_idc" comment:"开关"`
	Name           string `form:"name"  search:"type:contains;column:name;table:rs_idc" comment:"机房名称"`
	Status         string `form:"status"  search:"type:exact;column:status;table:rs_idc" comment:"机房状态"`
	Belong         string `form:"belong"  search:"type:exact;column:belong;table:rs_idc" comment:"机房归属"`
	TypeId         string `form:"typeId"  search:"type:exact;column:type_id;table:rs_idc" comment:"机房类型"`
	BusinessUser   string `form:"businessUser"  search:"type:exact;column:business_user;table:rs_idc" comment:"商务人员"`
	CustomUser     string `form:"customUser"  search:"type:exact;column:custom_user;table:rs_idc" comment:"所属客户"`
	Region         string `form:"region"  search:"type:exact;column:region;table:rs_idc" comment:"所在区域"`
	Charging       string `form:"charging"  search:"type:exact;column:charging;table:rs_idc" comment:"计费方式"`
	TransProvince  string `form:"transProvince"  search:"type:exact;column:trans_province;table:rs_idc" comment:"是否跨省"`
	Address        string `form:"address"  search:"type:contains;column:address;table:rs_idc" comment:"详细地址"`
	RsIdcOrder
}

type RsIdcOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:rs_idc"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_idc"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_idc"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_idc"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_idc"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_idc"`
	Layer         string `form:"layerOrder"  search:"type:order;column:layer;table:rs_idc"`
	Enable        string `form:"enableOrder"  search:"type:order;column:enable;table:rs_idc"`
	Desc          string `form:"descOrder"  search:"type:order;column:desc;table:rs_idc"`
	Name          string `form:"nameOrder"  search:"type:order;column:name;table:rs_idc"`
	Status        string `form:"statusOrder"  search:"type:order;column:status;table:rs_idc"`
	Belong        string `form:"belongOrder"  search:"type:order;column:belong;table:rs_idc"`
	TypeId        string `form:"typeIdOrder"  search:"type:order;column:type_id;table:rs_idc"`
	BusinessUser  string `form:"businessUserOrder"  search:"type:order;column:business_user;table:rs_idc"`
	CustomUser    string `form:"customUserOrder"  search:"type:order;column:custom_user;table:rs_idc"`
	Region        string `form:"regionOrder"  search:"type:order;column:region;table:rs_idc"`
	Charging      string `form:"chargingOrder"  search:"type:order;column:charging;table:rs_idc"`
	Price         string `form:"priceOrder"  search:"type:order;column:price;table:rs_idc"`
	WeChatName    string `form:"weChatNameOrder"  search:"type:order;column:we_chat_name;table:rs_idc"`
	IpV6          string `form:"ipV6Order"  search:"type:order;column:ip_v6;table:rs_idc"`
	TransProvince string `form:"transProvinceOrder"  search:"type:order;column:trans_province;table:rs_idc"`
	Address       string `form:"addressOrder"  search:"type:order;column:address;table:rs_idc"`
}

func (m *RsIdcGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsIdcInsertReq struct {
	Id            int    `json:"-" comment:"主键编码"` // 主键编码
	Layer         string `json:"layer" comment:"排序"`
	Enable        string `json:"enable" comment:"开关"`
	Desc          string `json:"desc" comment:"备注"`
	Name          string `json:"name" comment:"机房名称"`
	Status        string `json:"status" comment:"机房状态"`
	Belong        string `json:"belong" comment:"机房归属"`
	TypeId        string `json:"typeId" comment:"机房类型"`
	BusinessUser  string `json:"businessUser" comment:"商务人员"`
	CustomUser    string `json:"customUser" comment:"所属客户"`
	Region        string `json:"region" comment:"所在区域"`
	Charging      string `json:"charging" comment:"计费方式"`
	Price         string `json:"price" comment:"单价"`
	WeChatNumber  string `json:"weChatNumber" comment:"企业微信"`
	IpV6          string `json:"ipV6" comment:"是否IPV6"`
	TransProvince string `json:"transProvince" comment:"是否跨省"`
	Address       string `json:"address" comment:"详细地址"`
	common.ControlBy
}

func (s *RsIdcInsertReq) Generate(model *models.RsIdc) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Layer = s.Layer
	model.Enable = s.Enable
	model.Desc = s.Desc
	model.Name = s.Name
	model.Status = s.Status
	model.Belong = s.Belong
	model.TypeId = s.TypeId
	model.BusinessUser = s.BusinessUser
	model.CustomUser = s.CustomUser
	model.Region = s.Region
	model.Charging = s.Charging
	model.Price = s.Price
	model.WeChatNumber = s.WeChatNumber
	model.IpV6 = s.IpV6
	model.TransProvince = s.TransProvince
	model.Address = s.Address
}

func (s *RsIdcInsertReq) GetId() interface{} {
	return s.Id
}

type RsIdcUpdateReq struct {
	Id            int    `uri:"id" comment:"主键编码"` // 主键编码
	Layer         string `json:"layer" comment:"排序"`
	Enable        string `json:"enable" comment:"开关"`
	Desc          string `json:"desc" comment:"备注"`
	Name          string `json:"name" comment:"机房名称"`
	Status        string `json:"status" comment:"机房状态"`
	Belong        string `json:"belong" comment:"机房归属"`
	TypeId        string `json:"typeId" comment:"机房类型"`
	BusinessUser  string `json:"businessUser" comment:"商务人员"`
	CustomUser    string `json:"customUser" comment:"所属客户"`
	Region        string `json:"region" comment:"所在区域"`
	Charging      string `json:"charging" comment:"计费方式"`
	Price         string `json:"price" comment:"单价"`
	WeChatNumber  string `json:"weChatNumber" comment:"企业微信"`
	IpV6          string `json:"ipV6" comment:"是否IPV6"`
	TransProvince string `json:"transProvince" comment:"是否跨省"`
	Address       string `json:"address" comment:"详细地址"`
	common.ControlBy
}

func (s *RsIdcUpdateReq) Generate(model *models.RsIdc) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Layer = s.Layer
	model.Enable = s.Enable
	model.Desc = s.Desc
	model.Name = s.Name
	model.Status = s.Status
	model.Belong = s.Belong
	model.TypeId = s.TypeId
	model.BusinessUser = s.BusinessUser
	model.CustomUser = s.CustomUser
	model.Region = s.Region
	model.Charging = s.Charging
	model.Price = s.Price
	model.WeChatNumber = s.WeChatNumber
	model.IpV6 = s.IpV6
	model.TransProvince = s.TransProvince
	model.Address = s.Address
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
