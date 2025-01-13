package models

import (
	"database/sql"
	"encoding/json"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/models"
	"gorm.io/gorm"
)


type HostDesc struct {
	Desc string `json:"desc"`
	SuspendBilling  string `json:"suspend_billing"`
}
type RsHost struct {
	models.Model

	HealthyAt     sql.NullTime `json:"healthy" gorm:"comment:存活上报时间"`
	Desc          string       `json:"-" gorm:"comment:描述信息"`
	DescJson          *HostDesc       `json:"desc" gorm:"-"`
	NetworkType   int          `json:"networkType" gorm:"default:2;comment:网络类型"`
	HostName      string       `json:"hostname" gorm:"type:varchar(100);comment:主机名"`
	Sn            string       `json:"sn" gorm:"type:varchar(100);index;comment:sn"`
	Cpu           int          `json:"cpu" gorm:"type:bigint;comment:总核数"`
	Ip            string       `json:"ip" gorm:"type:varchar(20);comment:ip"`
	Mac           string       `json:"mac" gorm:"type:varchar(30);comment:mac"`
	Mask          string       `json:"mask" gorm:"type:varchar(30);comment:mask"`
	Gateway       string       `json:"gateway" gorm:"type:varchar(30);comment:gateway"`
	PublicIp      string       `json:"publicIp" gorm:"type:varchar(20);comment:公网IP"`
	Memory        uint64       `json:"memory" gorm:"comment:总内存"`
	Kernel        string       `json:"kernel" gorm:"type:varchar(100);comment:内核版本"`
	Version       string       `json:"version" gorm:"type:varchar(20);comment:客户端版本"`
	Belong        int          `json:"belong" gorm:"type:int;default:1;comment:机器归属"`
	RemotePort    string       `json:"remotePort" gorm:"type:varchar(12);comment映射端口号"`
	Remark        string       `json:"remark" gorm:"type:varchar(60);comment:备注"` //166陕西延安宜川集义郭东机房电信1-2-11(30*100M) 拆分解析到线路和带宽
	Status        int          `json:"status" gorm:"type:int;comment:主机状态"`
	Balance       float64      `json:"balance" gorm:"type:varchar(50);comment:总带宽"`
	Region        string       `json:"region" gorm:"type:varchar(80)comment:省份城市多ID"`
	Isp           int          `json:"isp" gorm:"type:varchar(16);comment:运营商"`
	TransProvince int          `json:"transProvince" gorm:"default:0;comment:是否跨省"`
	LineType      int          `json:"lineType" gorm:"type:int(1);default:0;comment:线路类型"`
	AllLine       int          `json:"allLine" gorm:"type:int(1);default:0;comment:机器总线路"`
	LineBandwidth float64      `json:"lineBandwidth"  gorm:"default:0;comment:单条线路带宽"`
	Usage         float64      `json:"usage"  gorm:"comment:利用率"`
	PercentValue  float64      `json:"percentValue" gorm:"comment:计算昨天95带宽值"`
	Auth          int          `json:"auth" gorm:"type:int(1);default:1;comment:是否有主机权限"`
	ProbeShell    string       `json:"probeShell" gorm:"type:varchar(100);comment:主动探测主机命令"`
	Idc           int          `json:"idc" gorm:"type:int(11);comment:idc"`
	IdcInfo       interface{}  `json:"idcInfo" gorm:"-"`
	SuspendBilling bool `json:"suspend_billing" gorm:"default:1;comment:是否暂停计费"`
	Business      []RsBusiness `json:"business" gorm:"many2many:host_bind_business;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:business_id;"`
	Tag           []RsTag      `json:"tag" gorm:"many2many:host_bind_tag;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:tag_id;"`
	models.ExtendUserBy
	models.ModelTime
	models.ControlBy
	ExtendHostInfo `gorm:"-"`
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
type ExtendHostInfo struct {
	System        map[string]interface{} `json:"system" gorm:"-"`
	NetDevice     interface{}            `json:"netDevice" gorm:"-"`
	DialList      interface{}            `json:"dialList" gorm:"-"`
	MemoryMonitor interface{}            `json:"memoryMonitor"  gorm:"-"`
	Disk          interface{}            `json:"disk" gorm:"-"`
}

func (RsHost) TableName() string {
	return "rs_host"
}

func (e *RsHost) Generate() models.ActiveRecord {
	o := *e
	return &o
}
func (e *RsHost) AfterFind(tx *gorm.DB) (err error) {
	var user models2.SysUser
	userId := e.CreateBy
	if e.UpdateBy != 0 {
		userId = e.UpdateBy
	}
	if userId == 0 {
		return
	}
	tx.Model(&user).Select("user_id,username").Where("user_id = ?", userId).Limit(1).Find(&user)


	if user.UserId > 0 {
		e.UpdatedUser = user.Username
	}
	descModel :=&HostDesc{}
	if e.Desc != ""{
		marErr:=json.Unmarshal([]byte(e.Desc),&descModel)
		if marErr!=nil{
			descModel.Desc = e.Desc
		}
		e.DescJson = descModel
	}
	return
}
func (e *RsHost) GetId() interface{} {
	return e.Id
}
