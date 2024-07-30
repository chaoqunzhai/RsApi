package dto

import (
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HostMemory struct {
	T  uint64 `json:"t"`   //total
	A  uint64 `json:"a"`   //available
	U  uint64 `json:"u"`   //used
	SC uint64 `json:"s_c"` //swap_cached
	SF uint64 `json:"s_f"` //swap_free
	ST uint64 `json:"s_t"` //swap_total
}

type HostBusiness struct {
	Id int    `json:"id" form:"id"`
	Sn string `json:"sn" from:"sn" `
}
type BusinessSwitch struct {
	HostIds  []int          `json:"host_ids" form:"host_ids" `
	Business []HostBusiness `json:"business" form:"business" `
}

type RsHostGetPageReq struct {
	dto.Pagination `search:"-"`
	Enable         string `form:"enable"  search:"type:exact;column:enable;table:rs_host" comment:"开关"`
	HostName       string `form:"hostname"  search:"type:contains;column:host_name;table:rs_host" comment:"主机名"`
	Sn             string `form:"sn"  search:"type:contains;column:sn;table:rs_host" comment:"sn"`
	Ip             string `form:"ip"  search:"type:contains;column:ip;table:rs_host" comment:"ip"`
	Kernel         string `form:"kernel"  search:"type:exact;column:kernel;table:rs_host" comment:"内核版本"`
	Belong         string `form:"belong"  search:"type:exact;column:belong;table:rs_host" comment:"机器归属"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:rs_host" comment:"备注"`
	Isp            string `form:"isp"  search:"type:exact;column:isp;table:rs_host" comment:"运营商"`
	Status         string `form:"status"  search:"type:exact;column:status;table:rs_host" comment:"主机状态"`
	BusinessSn     string `form:"businessSn"  search:"type:contains;column:business_sn;table:rs_host" comment:"业务SN"`
	Province       string `form:"province"  search:"type:exact;column:province;table:rs_host" comment:"省份"`
	RsHostOrder
}

type RsHostOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:rs_host"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_host"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_host"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_host"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_host"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_host"`
	Layer     string `form:"layerOrder"  search:"type:order;column:layer;table:rs_host"`
	Enable    string `form:"enableOrder"  search:"type:order;column:enable;table:rs_host"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:rs_host"`
	HostName  string `form:"hostNameOrder"  search:"type:order;column:host_name;table:rs_host"`
	Sn        string `form:"snOrder"  search:"type:order;column:sn;table:rs_host"`
	Cpu       string `form:"cpuOrder"  search:"type:order;column:cpu;table:rs_host"`
	Ip        string `form:"ipOrder"  search:"type:order;column:ip;table:rs_host"`
	Memory    string `form:"memoryOrder"  search:"type:order;column:memory;table:rs_host"`
	Disk      string `form:"diskOrder"  search:"type:order;column:disk;table:rs_host"`
	Kernel    string `form:"kernelOrder"  search:"type:order;column:kernel;table:rs_host"`
	Belong    string `form:"belongOrder"  search:"type:order;column:belong;table:rs_host"`
	Remark    string `form:"remarkOrder"  search:"type:order;column:remark;table:rs_host"`
	Operator  string `form:"operatorOrder"  search:"type:order;column:operator;table:rs_host"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:rs_host"`
	NetDevice string `form:"netDeviceOrder"  search:"type:order;column:net_device;table:rs_host"`
	Balance   string `form:"balanceOrder"  search:"type:order;column:balance;table:rs_host"`
	Isp       string `form:"ispOrder"  search:"type:order;column:isp;table:rs_host"`
}

func (m *RsHostGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsHostInsertReq struct {
	Id        int     `json:"-" comment:"主键编码"` // 主键编码
	Layer     int     `json:"layer" comment:"排序"`
	Enable    bool    `json:"enable" comment:"开关"`
	Desc      string  `json:"desc" comment:"描述信息"`
	HostName  string  `json:"hostName" comment:"主机名"`
	Sn        string  `json:"sn" comment:"sn"`
	Cpu       string  `json:"cpu" comment:"总核数"`
	Ip        string  `json:"ip" comment:"ip"`
	Memory    uint64  `json:"memory" comment:"总内存"`
	Disk      string  `json:"disk" comment:"总磁盘"`
	Kernel    string  `json:"kernel" comment:"内核版本"`
	Belong    int     `json:"belong" comment:"机器归属"`
	Remark    string  `json:"remark" comment:"备注"`
	Operator  string  `json:"operator" comment:"运营商"`
	Status    int     `json:"status" comment:"主机状态"`
	NetDevice string  `json:"netDevice" comment:"网卡信息"`
	Balance   float64 `json:"balance" comment:"总带宽"`
	Address   string  `json:"address" comment:"具体地址"`
	Region    string  `json:"region" comment:"城市"`
	Isp       int     `json:"isp" comment:"运营商"`
	common.ControlBy
}

func (s *RsHostInsertReq) Generate(model *models.RsHost) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Layer = s.Layer
	model.Enable = s.Enable
	model.Desc = s.Desc
	model.HostName = s.HostName
	model.Sn = s.Sn
	model.Cpu = s.Cpu
	model.Ip = s.Ip
	model.Memory = s.Memory
	model.Disk = s.Disk
	model.Kernel = s.Kernel
	model.Belong = s.Belong
	model.Remark = s.Remark
	model.Status = s.Status
	model.NetDevice = s.NetDevice
	model.Balance = s.Balance

	model.Address = s.Address
	model.Region = s.Region
	model.Isp = s.Isp
}

func (s *RsHostInsertReq) GetId() interface{} {
	return s.Id
}

type RsHostUpdateReq struct {
	Id        int     `uri:"id" comment:"主键编码"` // 主键编码
	Layer     int     `json:"layer" comment:"排序"`
	Enable    bool    `json:"enable" comment:"开关"`
	Desc      string  `json:"desc" comment:"描述信息"`
	HostName  string  `json:"hostName" comment:"主机名"`
	Sn        string  `json:"sn" comment:"sn"`
	Cpu       string  `json:"cpu" comment:"总核数"`
	Ip        string  `json:"ip" comment:"ip"`
	Memory    uint64  `json:"memory" comment:"总内存"`
	Disk      string  `json:"disk" comment:"总磁盘"`
	Kernel    string  `json:"kernel" comment:"内核版本"`
	Belong    int     `json:"belong" comment:"机器归属"`
	Remark    string  `json:"remark" comment:"备注"`
	Status    int     `json:"status" comment:"主机状态"`
	NetDevice string  `json:"netDevice" comment:"网卡信息"`
	Balance   float64 `json:"balance" comment:"总带宽"`
	Address   string  `json:"address" comment:"具体地址"`
	Region    string  `json:"region" comment:"城市"`
	Isp       int     `json:"isp" comment:"运营商"`
	common.ControlBy
}

func (s *RsHostUpdateReq) Generate(model *models.RsHost) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Layer = s.Layer
	model.Enable = s.Enable
	model.Desc = s.Desc
	model.HostName = s.HostName
	model.Sn = s.Sn
	model.Cpu = s.Cpu
	model.Ip = s.Ip
	model.Memory = s.Memory
	model.Disk = s.Disk
	model.Kernel = s.Kernel
	model.Belong = s.Belong
	model.Remark = s.Remark

	model.Status = s.Status
	model.NetDevice = s.NetDevice
	model.Balance = s.Balance

	model.Address = s.Address
	model.Region = s.Region
	model.Isp = s.Isp
}

func (s *RsHostUpdateReq) GetId() interface{} {
	return s.Id
}

// RsHostGetReq 功能获取请求参数
type RsHostGetReq struct {
	Id int `uri:"id"`
}

func (s *RsHostGetReq) GetId() interface{} {
	return s.Id
}

// RsHostDeleteReq 功能删除请求参数
type RsHostDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsHostDeleteReq) GetId() interface{} {
	return s.Ids
}
