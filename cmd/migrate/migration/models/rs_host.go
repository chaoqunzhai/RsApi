package models

import (
	"database/sql"
	"go-admin/common/models"
	"time"
)

const (
	YIDONG      = 1
	DIANXIN     = 2
	LIANTONG    = 3
	Other       = 3
	HostLoading = 0 //链接中
	HostUp      = 1 //在线
	HostOffline = 1 //离线
)

type Host struct {
	RichGlobal
	HealthyAt     sql.NullTime `json:"healthy" gorm:"comment:存活上报时间"`
	HostName      string       `json:"hostname" gorm:"type:varchar(100);comment:主机名;not null"`
	Sn            string       `json:"sn" gorm:"index;comment:sn"`
	CPU           int          `json:"cpu" gorm:"comment:总核数"`
	Ip            string       `json:"ip" gorm:"type:varchar(20);comment:ip"`
	Memory        uint64       `json:"memory" gorm:"comment:总内存"`
	Disk          string       `json:"disk" gorm:"type:varchar(20);comment:总磁盘"`
	NetDevice     string       `json:"netDevice" gorm:"type:varchar(120);comment:网卡信息"`
	Kernel        string       `json:"kernel" gorm:"type:varchar(100);comment:内核版本"`
	Balance       float64      `json:"balance" gorm:"type:varchar(50);comment:总带宽"`
	Belong        int          `json:"belong" gorm:"type:int(1);default:0;comment:机器归属"`
	Remark        string       `json:"remark" gorm:"type:varchar(60);comment:备注;default:'';"`
	Isp           int          `json:"isp" gorm:"type:int(1);default:1;comment:运营商"`
	Status        int          `json:"status" gorm:"type:int(1);default:0;comment:主机状态"`
	Region        string       `json:"region" gorm:"type:varchar(50);comment:省份城市多ID"`
	Address       string       `json:"address" gorm:"type:varchar(100);comment:详细地址"`
	TransProvince bool         `json:"transProd" gorm:"default:false;comment:是否跨省"`
	LineType      int          `json:"lineType" gorm:"type:int(1);default:0;comment:线路类型"`
	Idc           int          `json:"idc" gorm:"index;type:int(11);comment:关联的IDC"`
	Business      []Business   `gorm:"many2many:host_bind_business;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:business_id;"`
	Tag           []Tag        `gorm:"many2many:host_bind_tag;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:tag_id;"`
}

func (Host) TableName() string {
	return "rs_host"
}

type HostSoftware struct {
	models.Model
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	HostId    int       `json:"host_id" gorm:"index;comment:关联主机ID"`
	Key       string    `json:"key" gorm:"type:varchar(30);comment:服务名称"`
	Value     string    `json:"value" gorm:"type:varchar(100);comment:服务内容"`
	Desc      string    `json:"desc" gorm:"type:varchar(30);comment:备注"`
	models.ModelTime
}

func (HostSoftware) TableName() string {
	return "rs_host_software"
}

type HostSystem struct {
	models.Model
	UpdatedAt      time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	HostId         int       `json:"host_id" gorm:"index;comment:关联主机ID"`
	TransmitNumber float64   `json:"transmit_number" gorm:"type:varchar(30);comment:TransmitNumber"`
	ReceiveNumber  float64   `json:"receive_number" gorm:"type:varchar(30);comment:ReceiveNumber"`
	MemoryData     string    `json:"memory" gorm:"type:varchar(255);comment:当前内容使用率"`
}

func (HostSystem) TableName() string {
	return "rs_host_system"
}

//主机业务切换记录表

type HostSwitchLog struct {
	Model
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	CreateBy   int       `json:"createBy" gorm:"index;comment:创建者"`
	JobId      string    `json:"job_id" gorm:"type:varchar(50);comment:任务ID" `
	HostId     int       `json:"host_id" gorm:"index;comment:切换的主机ID"`
	BusinessId int       `json:"business_id" gorm:"index;comment:切换的新业务ID"`
	BusinessSn string    `json:"business_sn" gorm:"type:varchar(30);index;comment:原来的业务SN"`
	Desc       string    `json:"desc" gorm:"type:varchar(120);comment:切换业务备注" `
}

func (HostSwitchLog) TableName() string {
	return "rs_host_switch_log"
}

//主机计费 计算表

type HostChargingDay struct {
	Model
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	CreateBy   int       `json:"createBy" gorm:"index;comment:创建者"`
	BusinessId int       `json:"business_id" gorm:"index;comment:切换的业务ID"`
	HostId     int       `json:"host_id" gorm:"index;comment:关联的主机ID"`
	Cost       float64   `json:"cost" gorm:"index;comment:计算的费用"`
	Banlance95 float64   `json:"banlance95" gorm:"index;comment:95带宽值"`
	Sla        string    `json:"sla" gorm:"type:varchar(120);comment:触发SLA原因"`
	Desc       string    `json:"desc" gorm:"type:varchar(120);comment:计费备注" `
}

func (HostChargingDay) TableName() string {
	return "rs_host_charging_day"
}
