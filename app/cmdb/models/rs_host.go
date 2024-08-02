package models

import (
	"database/sql"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsHost struct {
	models.Model
	HealthyAt     sql.NullTime `json:"healthy" gorm:"comment:存活上报时间"`
	Layer         int          `json:"layer" gorm:"default:1;type:tinyint;comment:排序"`
	Enable        int          `json:"enable" gorm:"default:1;type:tinyint(1);comment:开关"`
	Desc          string       `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	NetworkType   int          `json:"networkType" gorm:"default:2;comment:网络类型"`
	HostName      string       `json:"hostname" gorm:"type:varchar(100);comment:主机名"`
	Sn            string       `json:"sn" gorm:"type:varchar(191);comment:sn"`
	Cpu           string       `json:"cpu" gorm:"type:bigint;comment:总核数"`
	Ip            string       `json:"ip" gorm:"type:varchar(20);comment:ip"`
	Memory        uint64       `json:"memory" gorm:"comment:总内存"`
	Disk          string       `json:"disk" gorm:"type:longtext;comment:总磁盘"`
	Kernel        string       `json:"kernel" gorm:"type:varchar(100);comment:内核版本"`
	Belong        int          `json:"belong" gorm:"type:int;comment:机器归属"`
	Remark        string       `json:"remark" gorm:"type:varchar(60);comment:备注"` //166陕西延安宜川集义郭东机房电信1-2-11(30*100M) 拆分解析到线路和带宽
	Status        int          `json:"status" gorm:"type:int;comment:主机状态"`
	NetDevice     string       `json:"netDevice" gorm:"type:varchar(120);comment:网卡信息"`
	Balance       float64      `json:"balance" gorm:"type:varchar(50);comment:总带宽"`
	Region        string       `json:"region" gorm:"type:varchar(80)comment:省份城市多ID"`
	Isp           int          `json:"isp" gorm:"type:varchar(16);comment:运营商"`
	TransProvince int          `json:"transProvince" gorm:"default:0;comment:是否跨省"`
	LineType      int          `json:"lineType" gorm:"type:int(1);default:0;comment:线路类型"`
	Idc           int          `json:"idc" gorm:"type:int(11);comment:idc"`
	Business      []RsBusiness `json:"business" gorm:"many2many:host_bind_business;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:business_id;"`
	Tag           []RsTag      `json:"tag" gorm:"many2many:host_bind_tag;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:tag_id;"`
	models.ExtendUserBy
	models.ModelTime
	models.ControlBy
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
	tx.Model(&user).Select("user_id,username").Where("user_id = ?", userId).Limit(1).Find(&user)

	if user.UserId > 0 {
		e.UpdatedUser = user.Username
	}
	return
}
func (e *RsHost) GetId() interface{} {
	return e.Id
}
