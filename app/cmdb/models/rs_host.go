package models

import (

	"go-admin/common/models"

)

type RsHost struct {
    models.Model
    
    Layer string `json:"layer" gorm:"type:tinyint;comment:排序"` 
    Enable string `json:"enable" gorm:"type:tinyint(1);comment:开关"` 
    Desc string `json:"desc" gorm:"type:varchar(35);comment:描述信息"` 
    HostName string `json:"hostName" gorm:"type:varchar(100);comment:主机名"` 
    Sn string `json:"sn" gorm:"type:varchar(191);comment:sn"` 
    Cpu string `json:"cpu" gorm:"type:bigint;comment:总核数"` 
    Ip string `json:"ip" gorm:"type:varchar(20);comment:ip"` 
    Memory string `json:"memory" gorm:"type:longtext;comment:总内存"` 
    Disk string `json:"disk" gorm:"type:longtext;comment:总磁盘"` 
    Kernel string `json:"kernel" gorm:"type:varchar(100);comment:内核版本"` 
    Belong string `json:"belong" gorm:"type:int;comment:机器归属"` 
    Remark string `json:"remark" gorm:"type:varchar(60);comment:备注"` 
    Operator string `json:"operator" gorm:"type:int;comment:运营商"` 
    Status string `json:"status" gorm:"type:int;comment:主机状态"` 
    NetDevice string `json:"netDevice" gorm:"type:varchar(120);comment:网卡信息"` 
    Balance string `json:"balance" gorm:"type:varchar(50);comment:总带宽"` 
    BusinessSn string `json:"businessSn" gorm:"type:varchar(120);comment:业务SN"` 
    Province string `json:"province" gorm:"type:varchar(20);comment:省份"` 
    City string `json:"city" gorm:"type:varchar(30);comment:城市"` 
    Isp string `json:"isp" gorm:"type:varchar(16);comment:运营商"` 
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

func (e *RsHost) GetId() interface{} {
	return e.Id
}