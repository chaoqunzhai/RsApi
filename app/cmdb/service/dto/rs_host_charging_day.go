package dto

import (

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsHostChargingDayGetPageReq struct {
	dto.Pagination     `search:"-"`
    BusinessId string `form:"businessId"  search:"type:exact;column:business_id;table:rs_host_charging_day" comment:"切换的业务ID"`
    HostId string `form:"hostId"  search:"type:exact;column:host_id;table:rs_host_charging_day" comment:"关联的主机ID"`
    RsHostChargingDayOrder
}

type RsHostChargingDayOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:rs_host_charging_day"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_host_charging_day"`
    BusinessId string `form:"businessIdOrder"  search:"type:order;column:business_id;table:rs_host_charging_day"`
    HostId string `form:"hostIdOrder"  search:"type:order;column:host_id;table:rs_host_charging_day"`
    Cost string `form:"costOrder"  search:"type:order;column:cost;table:rs_host_charging_day"`
    Banlance95 string `form:"banlance95Order"  search:"type:order;column:banlance95;table:rs_host_charging_day"`
    Sla string `form:"slaOrder"  search:"type:order;column:sla;table:rs_host_charging_day"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:rs_host_charging_day"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_host_charging_day"`
    
}

func (m *RsHostChargingDayGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsHostChargingDayInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    BusinessId string `json:"businessId" comment:"切换的业务ID"`
    HostId string `json:"hostId" comment:"关联的主机ID"`
    Cost string `json:"cost" comment:"计算的费用"`
    Banlance95 string `json:"banlance95" comment:"95带宽值"`
    Sla string `json:"sla" comment:"触发SLA原因"`
    Desc string `json:"desc" comment:"计费备注"`
    common.ControlBy
}

func (s *RsHostChargingDayInsertReq) Generate(model *models.RsHostChargingDay)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.BusinessId = s.BusinessId
    model.HostId = s.HostId
    model.Cost = s.Cost
    model.Banlance95 = s.Banlance95
    model.Sla = s.Sla
    model.Desc = s.Desc
}

func (s *RsHostChargingDayInsertReq) GetId() interface{} {
	return s.Id
}

type RsHostChargingDayUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    BusinessId string `json:"businessId" comment:"切换的业务ID"`
    HostId string `json:"hostId" comment:"关联的主机ID"`
    Cost string `json:"cost" comment:"计算的费用"`
    Banlance95 string `json:"banlance95" comment:"95带宽值"`
    Sla string `json:"sla" comment:"触发SLA原因"`
    Desc string `json:"desc" comment:"计费备注"`
    common.ControlBy
}

func (s *RsHostChargingDayUpdateReq) Generate(model *models.RsHostChargingDay)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BusinessId = s.BusinessId
    model.HostId = s.HostId
    model.Cost = s.Cost
    model.Banlance95 = s.Banlance95
    model.Sla = s.Sla
    model.Desc = s.Desc
}

func (s *RsHostChargingDayUpdateReq) GetId() interface{} {
	return s.Id
}

// RsHostChargingDayGetReq 功能获取请求参数
type RsHostChargingDayGetReq struct {
     Id int `uri:"id"`
}
func (s *RsHostChargingDayGetReq) GetId() interface{} {
	return s.Id
}

// RsHostChargingDayDeleteReq 功能删除请求参数
type RsHostChargingDayDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsHostChargingDayDeleteReq) GetId() interface{} {
	return s.Ids
}
