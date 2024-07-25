package models

import "go-admin/common/models"

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

	HostName string `json:"hostName" gorm:"type:varchar(100);comment:主机名;not null"`
	Sn       string `json:"sn" gorm:"index;comment:sn"`
	CPU      string `json:"cpu" gorm:"type:varchar(20);comment:cpu"`
	Ip string `json:"ip" gorm:"type:varchar(20);comment:ip"`
	Memory   string `json:"memory" gorm:"type:varchar(30);comment:memory"`
	Disk     string `json:"disk" gorm:"type:varchar(30);comment:disk"`
	Kernel   string `json:"kernel" gorm:"type:varchar(50);comment:内核"`
	Belong   int    `json:"belong" gorm:"type:int(1);default:0;comment:机器归属"`
	Remark   string `json:"remark" gorm:"type:varchar(100);comment:备注;default:'';"`
	Operator int    `json:"operator" gorm:"type:int(1);default:1;comment:运营商"`
	Status   int    `json:"status" gorm:"type:int(1);default:0;comment:主机状态"`

	Idc []Idc `gorm:"many2many:host_bind_idc;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:idc_id;"`
}

func (Host) TableName() string {
	return "rs_host"
}

type HostSoftware struct {
	models.Model
	HostId int `json:"host_id" gorm:"index;comment:关联主机ID"`
	Key string `json:"key" gorm:"type:varchar(30);comment:服务名称"`
	Value string `json:"value" gorm:"type:varchar(100);comment:服务内容"`
	Desc string `json:"desc" gorm:"type:varchar(30);comment:备注"`
	models.ModelTime
}
func (HostSoftware) TableName() string {
	return "rs_host_software"
}
