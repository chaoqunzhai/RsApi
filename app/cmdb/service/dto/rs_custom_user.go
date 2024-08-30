package dto

import (
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsCustomUserGetPageReq struct {
	dto.Pagination `search:"-"`
	Search         string `form:"search" search:"-"`
	UserName       string `form:"userName"  search:"type:contains;column:user_name;table:rs_custom_user" comment:"姓名"`
	CustomId       string `form:"customId"  search:"type:exact;column:custom_id;table:rs_custom_user" comment:"所属客户"`
	BuId           string `form:"buId"  search:"type:exact;column:bu_id;table:rs_custom_user" comment:"所属商务人员"`
	Phone          string `form:"phone"  search:"type:contains;column:phone;table:rs_custom_user" comment:"联系号码"`
	RsCustomUserOrder
}

type RsCustomUserOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:rs_custom_user"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_custom_user"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_custom_user"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_custom_user"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_custom_user"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_custom_user"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:rs_custom_user"`
	UserName  string `form:"userNameOrder"  search:"type:order;column:user_name;table:rs_custom_user"`
	CustomId  string `form:"customIdOrder"  search:"type:order;column:custom_id;table:rs_custom_user"`
	BuId      string `form:"buIdOrder"  search:"type:order;column:bu_id;table:rs_custom_user"`
	Phone     string `form:"phoneOrder"  search:"type:order;column:phone;table:rs_custom_user"`
	Email     string `form:"emailOrder"  search:"type:order;column:email;table:rs_custom_user"`
	Region    string `form:"regionOrder"  search:"type:order;column:region;table:rs_custom_user"`
	Dept      string `form:"deptOrder"  search:"type:order;column:dept;table:rs_custom_user"`
	Duties    string `form:"dutiesOrder"  search:"type:order;column:duties;table:rs_custom_user"`
	Address   string `form:"addressOrder"  search:"type:order;column:address;table:rs_custom_user"`
}

func (m *RsCustomUserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsCustomUserInsertReq struct {
	Id       int    `json:"-" comment:"主键编码"` // 主键编码
	Desc     string `json:"desc" comment:"描述信息"`
	UserName string `json:"userName" comment:"姓名"`
	CustomId int    `json:"customId" comment:"所属客户"`
	BuId     int    `json:"buId" comment:"所属商务人员"`
	Phone    string `json:"phone" comment:"联系号码"`
	Email    string `json:"email" comment:"联系邮箱"`
	Region   string `json:"region" comment:"省份城市多ID"`
	Dept     string `json:"dept" comment:"部门"`
	Duties   string `json:"duties" comment:"职务"`
	Address  string `json:"address" comment:"详细地址"`
	common.ControlBy
}

func (s *RsCustomUserInsertReq) Generate(model *models.RsCustomUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc
	model.UserName = s.UserName
	model.CustomId = s.CustomId
	model.BuId = s.BuId
	model.Phone = s.Phone
	model.Email = s.Email
	model.Region = s.Region
	model.Dept = s.Dept
	model.Duties = s.Duties
	model.Address = s.Address
}

func (s *RsCustomUserInsertReq) GetId() interface{} {
	return s.Id
}

type RsCustomUserUpdateReq struct {
	Id       int    `uri:"id" comment:"主键编码"` // 主键编码
	Desc     string `json:"desc" comment:"描述信息"`
	UserName string `json:"userName" comment:"姓名"`
	CustomId int    `json:"customId" comment:"所属客户"`
	BuId     int    `json:"buId" comment:"所属商务人员"`
	Phone    string `json:"phone" comment:"联系号码"`
	Email    string `json:"email" comment:"联系邮箱"`
	Region   string `json:"region" comment:"省份城市多ID"`
	Dept     string `json:"dept" comment:"部门"`
	Duties   string `json:"duties" comment:"职务"`
	Address  string `json:"address" comment:"详细地址"`
	common.ControlBy
}

func (s *RsCustomUserUpdateReq) Generate(model *models.RsCustomUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.UserName = s.UserName
	model.CustomId = s.CustomId
	model.BuId = s.BuId
	model.Phone = s.Phone
	model.Email = s.Email
	model.Region = s.Region
	model.Dept = s.Dept
	model.Duties = s.Duties
	model.Address = s.Address
}

func (s *RsCustomUserUpdateReq) GetId() interface{} {
	return s.Id
}

// RsCustomUserGetReq 功能获取请求参数
type RsCustomUserGetReq struct {
	Id int `uri:"id"`
}

func (s *RsCustomUserGetReq) GetId() interface{} {
	return s.Id
}

// RsCustomUserDeleteReq 功能删除请求参数
type RsCustomUserDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsCustomUserDeleteReq) GetId() interface{} {
	return s.Ids
}
