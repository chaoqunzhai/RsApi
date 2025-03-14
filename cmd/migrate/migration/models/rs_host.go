package models

import (
	"database/sql"
	"go-admin/common/models"
	"gorm.io/gorm"
	"time"
)

type Host struct {
	RichGlobal

	Layer         int          `json:"layer" gorm:"comment:自定义排序"`
	HealthyAt     sql.NullTime `json:"healthy" gorm:"comment:存活上报时间"`
	HostName      string       `json:"hostname" gorm:"type:varchar(100);comment:主机名;not null"`
	Sn            string       `json:"sn" gorm:"type:varchar(100);index;comment:sn"`
	CPU           int          `json:"cpu" gorm:"comment:总核数"`
	PublicIp      string       `json:"publicIp" gorm:"type:varchar(20);comment:公网IP"`
	Ip            string       `json:"ip" gorm:"type:varchar(20);comment:ip"`
	Mask          string       `json:"mask" gorm:"type:varchar(30);comment:mask"`
	Mac           string       `json:"mac" gorm:"type:varchar(30);comment:mac"`
	Gateway       string       `json:"gateway" gorm:"type:varchar(30);comment:gateway"`
	Memory        uint64       `json:"memory" gorm:"comment:总内存"`
	NetworkType   int          `json:"networkType" gorm:"default:2;comment:网络类型"`
	Kernel        string       `json:"kernel" gorm:"type:varchar(100);comment:内核版本"`
	Version       string       `json:"version" gorm:"type:varchar(20);comment:客户端版本"`
	Balance       float64      `json:"balance" gorm:"type:varchar(50);comment:总带宽"`
	Belong        int          `json:"belong" gorm:"type:int(1);default:1;comment:机器归属"`
	RemotePort    string       `json:"remotePort" gorm:"type:varchar(12);comment映射端口号"`
	Remark        string       `json:"remark" gorm:"type:varchar(60);comment:备注;default:'';"`
	Isp           int          `json:"isp" gorm:"type:int(1);default:1;comment:运营商"`
	Status        int          `json:"status" gorm:"type:int(1);default:0;comment:主机状态"`
	Region        string       `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	LineType      int          `json:"lineType" gorm:"type:int(1);default:0;comment:线路类型"`
	AllLine       int          `json:"allLine" gorm:"type:int(1);default:0;comment:机器总线路"`
	LineBandwidth float64      `json:"lineBandwidth"  gorm:"default:0;comment:单条线路带宽"`
	Usage         float64      `json:"usage"  gorm:"comment:利用率"`
	PercentValue  float64      `json:"percentValue" gorm:"comment:计算昨天95带宽值"`
	Idc           int          `json:"idc" gorm:"index;type:int(11);comment:关联的IDC"`
	Auth          int          `json:"auth" gorm:"type:int(1);default:1;comment:是否有主机权限"`
	ProbeShell    string       `json:"probeShell" gorm:"type:varchar(200);comment:主动探测主机命令"`
	TransProvince int          `form:"transProvince" search:"type:exact;column:trans_province;table:rs_host" comment:"是否跨省"`
	Business      []Business   `gorm:"many2many:host_bind_business;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:business_id;"`
	Tag           []Tag        `gorm:"many2many:host_bind_tag;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:tag_id;"`
	SuspendBilling bool `json:"suspend_billing" gorm:"default:1;comment:是否暂停计费"`
	Desc string `json:"desc" gorm:"comment:描述信息"` //描述
}

func (Host) TableName() string {
	return "rs_host"
}
type RsHostSuspendLog struct {
	models.Model
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	CreatedAt models.XTime          `json:"createdAt" gorm:"comment:创建时间"`
	Desc string `json:"desc" gorm:"type:varchar(200);comment:内容"`
	HostId int64 `json:"host_id" gorm:"comment:主机ID"`
	Enable bool `json:"enable" gorm:"comment:开启还是关闭"`
}
func (RsHostSuspendLog) TableName() string {
	return "rs_host_suspend_log"
}
type HostNetDevice struct {
	models.Model
	HostId    int          `json:"host_id" gorm:"index;comment:关联主机ID"`
	UpdatedAt models.XTime `json:"updatedAt" gorm:"comment:更新时间"`
	Name      string       `json:"name" gorm:"type:varchar(20);comment:网卡名称"`
	Status    int          `json:"status" gorm:"type:int(1);default:0;comment:网卡状态,1:正常 非1:异常"`
	Ip        string       `json:"ip" gorm:"type:varchar(50);comment:ip"`
	Mac       string       `json:"mac" gorm:"type:varchar(50);comment:mac"`
}

func (HostNetDevice) TableName() string {
	return "rs_host_netdevice"
}

type HostSoftware struct {
	models.Model
	UpdatedAt models.XTime `json:"updatedAt" gorm:"comment:更新时间"`
	HostId    int          `json:"host_id" gorm:"index;comment:关联主机ID"`
	Key       string       `json:"key" gorm:"type:varchar(30);comment:服务名称"`
	Value     string       `json:"value" gorm:"type:varchar(100);comment:服务内容"`
	Desc      string       `json:"desc" gorm:"type:varchar(30);comment:备注"`
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
	Disk           string    `json:"disk" gorm:"type:text;comment:所有磁盘信息"`
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
	BuTargetId int       `json:"bu_target_id" gorm:"index;comment:切换的新业务ID"`
	BuSource   string    `json:"bu_source"  gorm:"type:varchar(50);index;comment:原来的业务中文名称"`
	BuEnSource string    `json:"bu_en_source"  gorm:"type:varchar(50);index;comment:原来的业务英文名称"`
	Desc       string    `json:"desc" gorm:"type:varchar(100);comment:切换业务备注" `
}

func (HostSwitchLog) TableName() string {
	return "rs_host_switch_log"
}

//主机命令操作记录

type HostExecLog struct {
	Model
	JobId       string       `json:"job_id" gorm:"type:varchar(50);comment:任务ID" `
	CreatedAt   models.XTime `json:"createdAt" gorm:"comment:执行时间"`
	CreateBy    int          `json:"createBy" gorm:"index;comment:创建者"`
	HostId      int          `json:"host_id" gorm:"index;comment:关联的主机ID"`
	Status      int          `json:"status" gorm:"index;comment:执行状态,0:执行中  1:执行成功 -1:执行失败"`
	Module      string       `json:"module"  gorm:"type:varchar(50);comment:执行的模块"`
	Exec        string       `json:"exec" gorm:"type:text;comment:执行的命令"`
	OutPut      string       `json:"outPut" gorm:"comment:返回的结果"`
	UpdatedUser string       `json:"updatedUser" gorm:"-"`
}

func (e *HostExecLog) AfterFind(tx *gorm.DB) (err error) {
	var user SysUser
	userId := e.CreateBy

	if userId == 0 {
		return
	}
	tx.Model(&user).Select("user_id,username").Where("user_id = ?", userId).Limit(1).Find(&user)

	if user.UserId > 0 {
		e.UpdatedUser = user.Username
	}
	return
}
func (HostExecLog) TableName() string {
	return "rs_host_exec_log"
}
