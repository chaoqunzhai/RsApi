package dto

import (

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsDialGetPageReq struct {
	dto.Pagination     `search:"-"`
    Enable int64 `form:"enable"  search:"type:exact;column:enable;table:rs_dial" comment:"开关"`
    User string `form:"user"  search:"type:contains;column:user;table:rs_dial" comment:"用户名"`
    Status int64 `form:"status"  search:"type:exact;column:status;table:rs_dial" comment:"拨号状态,1:正常 非1:异常"`
    IdcId int64 `form:"idcId"  search:"type:exact;column:idc_id;table:rs_dial" comment:"关联的IDC"`
    HostId int64 `form:"hostId"  search:"type:exact;column:host_id;table:rs_dial" comment:"关联主机ID"`
    DeviceId int64 `form:"deviceId"  search:"type:exact;column:device_id;table:rs_dial" comment:"关联网卡ID"`
    RsDialOrder
}

type RsDialOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:rs_dial"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_dial"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_dial"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_dial"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_dial"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_dial"`
    Layer string `form:"layerOrder"  search:"type:order;column:layer;table:rs_dial"`
    Enable string `form:"enableOrder"  search:"type:order;column:enable;table:rs_dial"`
    Desc string `form:"descOrder"  search:"type:order;column:desc;table:rs_dial"`
    Number string `form:"numberOrder"  search:"type:order;column:number;table:rs_dial"`
    User string `form:"userOrder"  search:"type:order;column:user;table:rs_dial"`
    Pass string `form:"passOrder"  search:"type:order;column:pass;table:rs_dial"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:rs_dial"`
    IdcId string `form:"idcIdOrder"  search:"type:order;column:idc_id;table:rs_dial"`
    HostId string `form:"hostIdOrder"  search:"type:order;column:host_id;table:rs_dial"`
    DeviceId string `form:"deviceIdOrder"  search:"type:order;column:device_id;table:rs_dial"`
    
}

func (m *RsDialGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsDialInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Layer string `json:"layer" comment:"排序"`
    Enable int64 `json:"enable" comment:"开关"`
    Desc string `json:"desc" comment:"描述信息"`
    Number string `json:"number" comment:"账号"`
    User string `json:"user" comment:"用户名"`
    Pass string `json:"pass" comment:"密码"`
    Status int64 `json:"status" comment:"拨号状态,1:正常 非1:异常"`
    IdcId int64 `json:"idcId" comment:"关联的IDC"`
    HostId int64 `json:"hostId" comment:"关联主机ID"`
    DeviceId int64 `json:"deviceId" comment:"关联网卡ID"`
    common.ControlBy
}

func (s *RsDialInsertReq) Generate(model *models.RsDial)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.Layer = s.Layer
    model.Enable = s.Enable
    model.Desc = s.Desc
    model.Number = s.Number
    model.User = s.User
    model.Pass = s.Pass
    model.Status = s.Status
    model.IdcId = s.IdcId
    model.HostId = s.HostId
    model.DeviceId = s.DeviceId
}

func (s *RsDialInsertReq) GetId() interface{} {
	return s.Id
}

type RsDialUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Layer string `json:"layer" comment:"排序"`
    Enable int64 `json:"enable" comment:"开关"`
    Desc string `json:"desc" comment:"描述信息"`
    Number string `json:"number" comment:"账号"`
    User string `json:"user" comment:"用户名"`
    Pass string `json:"pass" comment:"密码"`
    Status int64 `json:"status" comment:"拨号状态,1:正常 非1:异常"`
    IdcId int64 `json:"idcId" comment:"关联的IDC"`
    HostId int64 `json:"hostId" comment:"关联主机ID"`
    DeviceId int64 `json:"deviceId" comment:"关联网卡ID"`
    common.ControlBy
}

func (s *RsDialUpdateReq) Generate(model *models.RsDial)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.Layer = s.Layer
    model.Enable = s.Enable
    model.Desc = s.Desc
    model.Number = s.Number
    model.User = s.User
    model.Pass = s.Pass
    model.Status = s.Status
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
