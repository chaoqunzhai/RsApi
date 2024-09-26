package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetOutboundOrderGetPageReq struct {
	dto.Pagination `search:"-"`
	CreateBy       string `form:"createBy"  search:"type:exact;column:create_by;table:asset_outbound_order" comment:"申请人"`
	Status         string `form:"status"  search:"type:exact;column:status;table:asset_outbound_order" comment:"状态"`
	Code           string `form:"code"  search:"type:contains;column:code;table:asset_outbound_order" comment:"出库编码"`
	CustomId       string `form:"customId"  search:"type:exact;column:custom_id;table:asset_outbound_order" comment:"所属客户ID"`
	Region         string `form:"region"  search:"type:contains;column:region;table:asset_outbound_order" comment:"省份城市多ID"`
	Ems            string `form:"ems"  search:"type:exact;column:ems;table:asset_outbound_order" comment:"物流公司"`
	TrackingNumber string `form:"trackingNumber"  search:"type:exact;column:tracking_number;table:asset_outbound_order" comment:"物流单号"`
	Address        string `form:"address"  search:"type:exact;column:address;table:asset_outbound_order" comment:"联系地址"`
	User           string `form:"user"  search:"type:exact;column:user;table:asset_outbound_order" comment:"联系人"`
	IdcId          string `form:"idcId"  search:"type:exact;column:idc_id;table:asset_outbound_order" comment:"idcId"`
	StartTime      string `form:"startTime"  search:"type:gte;column:created_at;table:asset_outbound_order" comment:"出库开始时间"`
	EndTime        string `form:"endTime"  search:"type:lte;column:created_at;table:asset_outbound_order" comment:"出库结束时间"`
	AssetOutboundOrderOrder
}

type AssetOutboundOrderOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:asset_outbound_order"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_outbound_order"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_outbound_order"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_outbound_order"`
	UpdatedAt      string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_outbound_order"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_outbound_order"`
	Desc           string `form:"descOrder"  search:"type:order;column:desc;table:asset_outbound_order"`
	Code           string `form:"codeOrder"  search:"type:order;column:code;table:asset_outbound_order"`
	CustomId       string `form:"customIdOrder"  search:"type:order;column:custom_id;table:asset_outbound_order"`
	PhoneNumber    string `form:"phoneNumberOrder"  search:"type:order;column:phone_number;table:asset_outbound_order"`
	Region         string `form:"regionOrder"  search:"type:order;column:region;table:asset_outbound_order"`
	Ems            string `form:"emsOrder"  search:"type:order;column:ems;table:asset_outbound_order"`
	TrackingNumber string `form:"trackingNumberOrder"  search:"type:order;column:tracking_number;table:asset_outbound_order"`
	Address        string `form:"addressOrder"  search:"type:order;column:address;table:asset_outbound_order"`
	User           string `form:"userOrder"  search:"type:order;column:user;table:asset_outbound_order"`
	IdcId          string `form:"idcIdOrder"  search:"type:order;column:idc_id;table:asset_outbound_order"`
	Count          string `form:"countOrder"  search:"type:order;column:count;table:asset_outbound_order"`
}

func (m *AssetOutboundOrderGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetOutboundOrderInsertReq struct {
	Id             int    `json:"-" comment:"主键编码"` // 主键编码
	Desc           string `json:"desc" comment:"描述信息"`
	Code           string `json:"code" comment:"出库编码"`
	CustomId       int    `json:"customId" comment:"所属客户ID"`
	PhoneNumber    string `json:"phoneNumber" comment:""`
	Region         string `json:"region" comment:"省份城市多ID"`
	Ems            string `json:"ems" comment:"物流公司"`
	TrackingNumber string `json:"trackingNumber" comment:"物流单号"`
	Address        string `json:"address" comment:"联系地址"`
	UserId         int    `json:"userId" comment:"联系人"`
	IdcId          int    `json:"idcId" comment:"idcId"`
	Asset          []int  `json:"asset"`
	Combination    []int  `json:"combination"`
	common.ControlBy
}

func (s *AssetOutboundOrderInsertReq) Generate(model *models.AssetOutboundOrder) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc
	model.Code = s.Code
	model.CustomId = s.CustomId
	model.PhoneNumber = s.PhoneNumber
	model.Region = s.Region
	model.Ems = s.Ems
	model.TrackingNumber = s.TrackingNumber
	model.Address = s.Address
	model.UserId = s.UserId
	model.IdcId = s.IdcId
}

func (s *AssetOutboundOrderInsertReq) GetId() interface{} {
	return s.Id
}

type AssetOutboundOrderUpdateReq struct {
	Id             int    `uri:"id" comment:"主键编码"` // 主键编码
	Desc           string `json:"desc" comment:"描述信息"`
	Code           string `json:"code" comment:"出库编码"`
	CustomId       int    `json:"customId" comment:"所属客户ID"`
	PhoneNumber    string `json:"phoneNumber" comment:""`
	Region         string `json:"region" comment:"省份城市多ID"`
	Ems            string `json:"ems" comment:"物流公司"`
	TrackingNumber string `json:"trackingNumber" comment:"物流单号"`
	Address        string `json:"address" comment:"联系地址"`
	UserId         int    `json:"userId" gorm:"comment:联系人"`
	IdcId          int    `json:"idcId" comment:"idcId"`
	common.ControlBy
}

func (s *AssetOutboundOrderUpdateReq) Generate(model *models.AssetOutboundOrder) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.Code = s.Code
	model.CustomId = s.CustomId
	model.PhoneNumber = s.PhoneNumber
	model.Region = s.Region
	model.Ems = s.Ems
	model.TrackingNumber = s.TrackingNumber
	model.Address = s.Address
	model.UserId = s.UserId
	model.IdcId = s.IdcId
}

func (s *AssetOutboundOrderUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetOutboundOrderGetReq 功能获取请求参数
type AssetOutboundOrderGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetOutboundOrderGetReq) GetId() interface{} {
	return s.Id
}

// AssetOutboundOrderDeleteReq 功能删除请求参数
type AssetOutboundOrderDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetOutboundOrderDeleteReq) GetId() interface{} {
	return s.Ids
}
