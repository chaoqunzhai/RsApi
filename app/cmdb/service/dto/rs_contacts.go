package dto

import (

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsContactsGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserName string `form:"userName"  search:"type:contains;column:user_name;table:rs_contacts" comment:"用户名"`
    CustomerId int64 `form:"customerId"  search:"type:exact;column:customer_id;table:rs_contacts" comment:"客户ID"`
    BuUser int64 `form:"buUser"  search:"type:exact;column:bu_user;table:rs_contacts" comment:"商务人员"`
    RsContactsOrder
}

type RsContactsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:rs_contacts"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_contacts"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_contacts"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_contacts"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_contacts"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_contacts"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:rs_contacts"`
    UserName string `form:"userNameOrder"  search:"type:order;column:user_name;table:rs_contacts"`
    CustomerId string `form:"customerIdOrder"  search:"type:order;column:customer_id;table:rs_contacts"`
    BuUser string `form:"buUserOrder"  search:"type:order;column:bu_user;table:rs_contacts"`
    Phone string `form:"phoneOrder"  search:"type:order;column:phone;table:rs_contacts"`
    Landline string `form:"landlineOrder"  search:"type:order;column:landline;table:rs_contacts"`
    Region string `form:"regionOrder"  search:"type:order;column:region;table:rs_contacts"`
    Email string `form:"emailOrder"  search:"type:order;column:email;table:rs_contacts"`
    Address string `form:"addressOrder"  search:"type:order;column:address;table:rs_contacts"`
    Department string `form:"departmentOrder"  search:"type:order;column:department;table:rs_contacts"`
    Duties string `form:"dutiesOrder"  search:"type:order;column:duties;table:rs_contacts"`
    
}

func (m *RsContactsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsContactsInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Desc string `json:"desc" comment:"描述信息"`
    UserName string `json:"userName" comment:"用户名"`
    CustomerId int64 `json:"customerId" comment:"客户ID"`
    BuUser int64 `json:"buUser" comment:"商务人员"`
    Phone string `json:"phone" comment:"电话号码"`
    Landline string `json:"landline" comment:"座机号"`
    Region string `json:"region" comment:"管理区域,也是城市ID"`
    Email string `json:"email" comment:"电话号码"`
    Address string `json:"address" comment:"地址"`
    Department string `json:"department" comment:"部门"`
    Duties string `json:"duties" comment:"职务"`
    common.ControlBy
}

func (s *RsContactsInsertReq) Generate(model *models.RsContacts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.Desc = s.Desc
    model.UserName = s.UserName
    model.CustomerId = s.CustomerId
    model.BuUser = s.BuUser
    model.Phone = s.Phone
    model.Landline = s.Landline
    model.Region = s.Region
    model.Email = s.Email
    model.Address = s.Address
    model.Department = s.Department
    model.Duties = s.Duties
}

func (s *RsContactsInsertReq) GetId() interface{} {
	return s.Id
}

type RsContactsUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Desc string `json:"desc" comment:"描述信息"`
    UserName string `json:"userName" comment:"用户名"`
    CustomerId int64 `json:"customerId" comment:"客户ID"`
    BuUser int64 `json:"buUser" comment:"商务人员"`
    Phone string `json:"phone" comment:"电话号码"`
    Landline string `json:"landline" comment:"座机号"`
    Region string `json:"region" comment:"管理区域,也是城市ID"`
    Email string `json:"email" comment:"电话号码"`
    Address string `json:"address" comment:"地址"`
    Department string `json:"department" comment:"部门"`
    Duties string `json:"duties" comment:"职务"`
    common.ControlBy
}

func (s *RsContactsUpdateReq) Generate(model *models.RsContacts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.Desc = s.Desc
    model.UserName = s.UserName
    model.CustomerId = s.CustomerId
    model.BuUser = s.BuUser
    model.Phone = s.Phone
    model.Landline = s.Landline
    model.Region = s.Region
    model.Email = s.Email
    model.Address = s.Address
    model.Department = s.Department
    model.Duties = s.Duties
}

func (s *RsContactsUpdateReq) GetId() interface{} {
	return s.Id
}

// RsContactsGetReq 功能获取请求参数
type RsContactsGetReq struct {
     Id int `uri:"id"`
}
func (s *RsContactsGetReq) GetId() interface{} {
	return s.Id
}

// RsContactsDeleteReq 功能删除请求参数
type RsContactsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsContactsDeleteReq) GetId() interface{} {
	return s.Ids
}
