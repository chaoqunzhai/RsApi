package dto

import (

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsTagGetPageReq struct {
	dto.Pagination     `search:"-"`
    Enable string `form:"enable"  search:"type:exact;column:enable;table:rs_tag" comment:"开关"`
    Name string `form:"name"  search:"type:contains;column:name;table:rs_tag" comment:"业务云名称"`
    RsTagOrder
}

type RsTagOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:rs_tag"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_tag"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_tag"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_tag"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_tag"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_tag"`
    Layer string `form:"layerOrder"  search:"type:order;column:layer;table:rs_tag"`
    Enable string `form:"enableOrder"  search:"type:order;column:enable;table:rs_tag"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:rs_tag"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:rs_tag"`
    
}

func (m *RsTagGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsTagInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Layer string `json:"layer" comment:"排序"`
    Enable string `json:"enable" comment:"开关"`
    Desc string `json:"desc" comment:"描述信息"`
    Name string `json:"name" comment:"业务云名称"`
    common.ControlBy
}

func (s *RsTagInsertReq) Generate(model *models.RsTag)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.Layer = s.Layer
    model.Enable = s.Enable
    model.Desc = s.Desc
    model.Name = s.Name
}

func (s *RsTagInsertReq) GetId() interface{} {
	return s.Id
}

type RsTagUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Layer string `json:"layer" comment:"排序"`
    Enable string `json:"enable" comment:"开关"`
    Desc string `json:"desc" comment:"描述信息"`
    Name string `json:"name" comment:"业务云名称"`
    common.ControlBy
}

func (s *RsTagUpdateReq) Generate(model *models.RsTag)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.Layer = s.Layer
    model.Enable = s.Enable
    model.Desc = s.Desc
    model.Name = s.Name
}

func (s *RsTagUpdateReq) GetId() interface{} {
	return s.Id
}

// RsTagGetReq 功能获取请求参数
type RsTagGetReq struct {
     Id int `uri:"id"`
}
func (s *RsTagGetReq) GetId() interface{} {
	return s.Id
}

// RsTagDeleteReq 功能删除请求参数
type RsTagDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsTagDeleteReq) GetId() interface{} {
	return s.Ids
}
