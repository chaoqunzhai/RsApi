package dto

import (

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsCustomerGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:contains;column:name;table:rs_customer" comment:"客户名称"`
    Level int64 `form:"level"  search:"type:exact;column:level;table:rs_customer" comment:"客户等级"`
    TypeId int64 `form:"typeId"  search:"type:exact;column:type_id;table:rs_customer" comment:"客户类型"`
    WorkStatus string `form:"workStatus"  search:"type:exact;column:work_status;table:rs_customer" comment:"合作状态"`
    RsCustomerOrder
}

type RsCustomerOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:rs_customer"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_customer"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_customer"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_customer"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_customer"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_customer"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:rs_customer"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:rs_customer"`
    Region string `form:"regionOrder"  search:"type:order;column:region;table:rs_customer"`
    Address string `form:"addressOrder"  search:"type:order;column:address;table:rs_customer"`
    Level string `form:"levelOrder"  search:"type:order;column:level;table:rs_customer"`
    TypeId string `form:"typeIdOrder"  search:"type:order;column:type_id;table:rs_customer"`
    WorkStatus string `form:"workStatusOrder"  search:"type:order;column:work_status;table:rs_customer"`
    
}

func (m *RsCustomerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsCustomerInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Desc string `json:"desc" comment:"描述信息"`
    Name string `json:"name" comment:"客户名称"`
    Region string `json:"region" comment:"省份城市多ID"`
    Address string `json:"address" comment:"地址"`
    Level int64 `json:"level" comment:"客户等级"`
    TypeId int64 `json:"typeId" comment:"客户类型"`
    WorkStatus string `json:"workStatus" comment:"合作状态"`
    common.ControlBy
}

func (s *RsCustomerInsertReq) Generate(model *models.RsCustomer)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.Desc = s.Desc
    model.Name = s.Name
    model.Region = s.Region
    model.Address = s.Address
    model.Level = s.Level
    model.TypeId = s.TypeId
    model.WorkStatus = s.WorkStatus
}

func (s *RsCustomerInsertReq) GetId() interface{} {
	return s.Id
}

type RsCustomerUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Desc string `json:"desc" comment:"描述信息"`
    Name string `json:"name" comment:"客户名称"`
    Region string `json:"region" comment:"省份城市多ID"`
    Address string `json:"address" comment:"地址"`
    Level int64 `json:"level" comment:"客户等级"`
    TypeId int64 `json:"typeId" comment:"客户类型"`
    WorkStatus string `json:"workStatus" comment:"合作状态"`
    common.ControlBy
}

func (s *RsCustomerUpdateReq) Generate(model *models.RsCustomer)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.Desc = s.Desc
    model.Name = s.Name
    model.Region = s.Region
    model.Address = s.Address
    model.Level = s.Level
    model.TypeId = s.TypeId
    model.WorkStatus = s.WorkStatus
}

func (s *RsCustomerUpdateReq) GetId() interface{} {
	return s.Id
}

// RsCustomerGetReq 功能获取请求参数
type RsCustomerGetReq struct {
     Id int `uri:"id"`
}
func (s *RsCustomerGetReq) GetId() interface{} {
	return s.Id
}

// RsCustomerDeleteReq 功能删除请求参数
type RsCustomerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsCustomerDeleteReq) GetId() interface{} {
	return s.Ids
}
