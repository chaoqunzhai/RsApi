package models

import (
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
	HostName   string `json:"hostName" gorm:"type:varchar(100);comment:主机名;not null"`
	Sn         string `json:"sn" gorm:"index;comment:sn"`
	CPU        int    `json:"cpu" gorm:"comment:总核数"`
	Ip         string `json:"ip" gorm:"type:varchar(20);comment:ip"`
	Memory     string `json:"memory" gorm:"comment:总内存"`
	Disk       string `json:"disk" gorm:"comment:总磁盘"`
	NetDevice  string `json:"netDevice" gorm:"type:varchar(120);comment:网卡信息"`
	Kernel     string `json:"kernel" gorm:"type:varchar(100);comment:内核版本"`
	Balance    string `json:"balance" gorm:"type:varchar(50);comment:总带宽"`
	Belong     int    `json:"belong" gorm:"type:int(1);default:0;comment:机器归属"`
	Remark     string `json:"remark" gorm:"type:varchar(60);comment:备注;default:'';"`
	Operator   int    `json:"operator" gorm:"type:int(1);default:1;comment:运营商"`
	Status     int    `json:"status" gorm:"type:int(1);default:0;comment:主机状态"`
	BusinessSn string `json:"businessSn" gorm:"type:varchar(120);comment:业务SN"`
	Province   string `json:"province" gorm:"type:varchar(20);comment:省份"`
	City       string `json:"city" gorm:"type:varchar(30);comment:城市"`
	Isp        string `json:"isp" gorm:"type:varchar(16);comment:运营商"`
	Idc        []Idc  `gorm:"many2many:host_bind_idc;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:idc_id;"`
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
	Balance        float64   `json:"balance" gorm:"type:varchar(30);comment:总带宽"`
	TransmitNumber float64   `json:"transmit_number" gorm:"type:varchar(30);comment:TransmitNumber"`
	ReceiveNumber  float64   `json:"receive_number" gorm:"type:varchar(30);comment:ReceiveNumber"`
	MemoryData     string    `json:"memory" gorm:"type:varchar(255);comment:当前内容使用率"`
}

func (HostSystem) TableName() string {
	return "rs_host_system"
}
