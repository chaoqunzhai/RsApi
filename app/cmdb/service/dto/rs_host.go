package dto

import (
	"fmt"
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strings"
)

// 主机拨号的信息

type RegisterDial struct {
	BU     string `json:"bu"`
	VlanId string `json:"vlan_id"`
	A      string `json:"a"` //账号
	P      string `json:"p"` //密码
	I      string `json:"i"` //绑定物理网卡
	D      string `json:"d"` //ppo线路名称 虚拟网卡名称
	S      int    `json:"s"` //状态
	Ip     string `json:"ip"`
	IpV6   string `json:"ipv_6"`
	Mac    string `json:"mac"`
	NT     string `json:"nt"` //网络类型
	NS     int    `json:"ns"` //联网状态
}

// 内存使用率的格式化

type HostMemory struct {
	T           uint64  `json:"t"`   //total
	A           uint64  `json:"a"`   //available
	U           uint64  `json:"u"`   //used
	SC          uint64  `json:"s_c"` //swap_cached
	SF          uint64  `json:"s_f"` //swap_free
	ST          uint64  `json:"s_t"` //swap_total
	CpuUsedRate float64 `json:"cpu_used_rate"`
	MemUsedRate float64 `json:"mem_used_rate"`
}

//磁盘的JSON格式化

type HDDevUsage struct {
	Dev string `json:"dev"`
	FS  string `json:"fs"`
	U   string `json:"u"`  //用户使用
	T   string `json:"t"`  //总量
	M   string `json:"m"`  //挂载点
	UP  string `json:"up"` //用户使用占比
}
type HostBusiness struct {
	Id interface{} `json:"id" form:"id"`
}
type BusinessSwitch struct {
	HostIds  []int          `json:"hostIds" form:"hostIds" comment:"需要切换的主机"`
	Business []HostBusiness `json:"business" form:"business" comment:"切换的新业务ID"`
	Desc     string         `json:"desc"  form:"desc"`
}

type UpdateStatus struct {
	Ids    []int  `json:"ids"`
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}
type HostBindDial struct {
	HostId   int `json:"hostId" form:"hostId" `
	DriverId int `json:"driverId" form:"driverId" `
	DialId   int `json:"dialId" form:"dialId" `
}

type HostBindIdc struct {
	IdcId   int   `json:"idcId" form:"idcId" `
	HostIds []int `json:"hostIds" form:"hostIds" `
}

type LabelRow struct {
	Id    int    `json:"id"`
	Label string `json:"label" form:"label" `
	Value string `json:"value" form:"value" `
}

type RsHostMonitorFlow struct {
	Id    int    `uri:"id"`
	Start string `form:"start"`
	End   string `form:"end"`
	Setup int    `form:"setup"`
}
type RsHostGetPageReq struct {
	dto.Pagination `search:"-"`
	Enable         string `form:"enable"  search:"type:exact;column:enable;table:rs_host" comment:"开关"`
	Idc            string `form:"idc"  search:"type:exact;column:idc;table:rs_host" comment:"关联机房"`
	IdcName        string `form:"idcName" search:"-"`
	IdcId          string `form:"idcId"  search:"type:exact;column:idc;table:rs_host" comment:"关联IDC"`
	IdcNumber      string `form:"idcNumber" search:"-"`
	BusinessId     string `form:"businessId" search:"-"`
	Region         string `form:"region"  search:"-" comment:"所在区域"`
	LineType       int    `form:"lineType" search:"type:exact;column:line_type;table:rs_host" comment:"线路类型"`
	HostName       string `form:"hostname"  search:"-"`
	HostId         string `form:"hostId"  search:"-" comment:"主机名"`
	Sn             string `form:"sn"  search:"type:contains;column:sn;table:rs_host" comment:"sn"`
	Ip             string `form:"ip"  search:"type:contains;column:ip;table:rs_host" comment:"ip"`
	NetworkType    int    `form:"networkType" search:"type:exact;column:network_type;table:rs_host" comment:"网络类型"`
	Kernel         string `form:"kernel"  search:"type:exact;column:kernel;table:rs_host" comment:"内核版本"`
	Belong         string `form:"belong"  search:"type:exact;column:belong;table:rs_host" comment:"机器归属"`
	TransProd      string `form:"transProd"  search:"type:exact;column:trans_province;table:rs_host" comment:"是否跨省"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:rs_host" comment:"备注"`
	Isp            string `form:"isp"  search:"type:exact;column:isp;table:rs_host" comment:"运营商"`
	Status         string `form:"status"  search:"type:exact;column:status;table:rs_host" comment:"主机状态"`
	BusinessSn     string `form:"businessSn"  search:"-" comment:"业务SN"`
	TransProvince  int    `form:"transProvince" search:"type:exact;column:trans_province;table:rs_host" comment:"是否跨省"`
	CustomId       int    `form:"customId"  search:"-" comment:"客户ID"`

	//给计费的请求设置的
	IncomeMonth string `form:"incomeMonth" search:"-" `

	RsHostOrder
}

type RsHostOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:rs_host"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_host"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_host"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_host"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_host"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_host"`
	Layer         string `form:"layerOrder"  search:"type:order;column:layer;table:rs_host"`
	Enable        string `form:"enableOrder"  search:"type:order;column:enable;table:rs_host"`
	Desc          string `form:"descOrder"  search:"type:order;column:desc;table:rs_host"`
	HostNameOrder string `form:"hostNameOrder"  search:"type:order;column:host_name;table:rs_host"`
	HealthyAtOrder string `form:"healthyAtOrder"  search:"type:order;column:healthy_at;table:rs_host"`
	UsageOrder    string `form:"usageOrder"  search:"type:order;column:usage;table:rs_host"`
	Sn            string `form:"snOrder"  search:"type:order;column:sn;table:rs_host"`
	Cpu           string `form:"cpuOrder"  search:"type:order;column:cpu;table:rs_host"`
	Ip            string `form:"ipOrder"  search:"type:order;column:ip;table:rs_host"`
	Memory        string `form:"memoryOrder"  search:"type:order;column:memory;table:rs_host"`
	Disk          string `form:"diskOrder"  search:"type:order;column:disk;table:rs_host"`
	Kernel        string `form:"kernelOrder"  search:"type:order;column:kernel;table:rs_host"`
	Belong        string `form:"belongOrder"  search:"type:order;column:belong;table:rs_host"`
	TransProvince string `form:"transProvinceOrder"  search:"type:order;column:trans_province;table:rs_host"`
	Remark        string `form:"remarkOrder"  search:"type:order;column:remark;table:rs_host"`
	Operator      string `form:"operatorOrder"  search:"type:order;column:operator;table:rs_host"`
	Status        string `form:"statusOrder"  search:"-"`
	NetDevice     string `form:"netDeviceOrder"  search:"type:order;column:net_device;table:rs_host"`
	Balance       string `form:"balanceOrder"  search:"type:order;column:balance;table:rs_host"`
	Isp           string `form:"ispOrder"  search:"type:order;column:isp;table:rs_host"`
}

func (m *RsHostGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsHostInsertReq struct {
	Id        int     `json:"-" comment:"主键编码"` // 主键编码
	Layer     int     `json:"layer" comment:"排序"`
	Desc      string  `json:"desc" comment:"描述信息"`
	HostName  string  `json:"hostName" comment:"主机名"`
	Sn        string  `json:"sn" comment:"sn"`
	Cpu       int     `json:"cpu" comment:"总核数"`
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
	model.Desc = s.Desc
	model.HostName = s.HostName
	model.Sn = s.Sn
	model.Cpu = s.Cpu
	model.Ip = s.Ip
	model.Memory = s.Memory

	model.Kernel = s.Kernel
	model.Belong = s.Belong
	model.Remark = s.Remark
	model.Status = s.Status

	model.Balance = s.Balance

	model.Region = s.Region
	model.Isp = s.Isp
}

func (s *RsHostInsertReq) GetId() interface{} {
	return s.Id
}

type RsHostUpdateReq struct {
	Id        int     `uri:"id" comment:"主键编码"` // 主键编码
	Layer     int     `json:"layer" comment:"排序"`
	Desc      string  `json:"desc" comment:"描述信息"`
	HostName  string  `json:"hostName" comment:"主机名"`
	Sn        string  `json:"sn" comment:"sn"`
	Cpu       int     `json:"cpu" comment:"总核数"`
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
	model.Desc = s.Desc
	model.HostName = s.HostName
	model.Sn = s.Sn
	model.Cpu = s.Cpu
	model.Ip = s.Ip
	model.Memory = s.Memory

	model.Kernel = s.Kernel
	model.Belong = s.Belong
	model.Remark = s.Remark

	model.Status = s.Status

	model.Balance = s.Balance

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

func (s *RsHostDeleteReq) GetIdStr() string {
	ca := make([]string, 0)
	for _, row := range s.Ids {
		ca = append(ca, fmt.Sprintf("%v", row))
	}
	return strings.Join(ca, ",")
}
