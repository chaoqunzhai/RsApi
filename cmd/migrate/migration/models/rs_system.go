package models

import (
	"go-admin/common/models"
)

type Business struct {
	RichGlobal
	Status        int    `json:"status" gorm:"type:int(1);default:0;comment:拨号状态,1:正常 非1:异常"`
	Name          string `json:"name" gorm:"index;type:varchar(50);comment:业务中文名称"`
	EnName        string `json:"enName" gorm:"index;type:varchar(30);comment:业务英文名字"`
	BillingMethod int    `json:"billingMethod" gorm:"type:int(1);default:6,comment:计费方式"`
	ParentId      int    `json:"parentId" gorm:"comment:父业务"`
	OpeMonitor    bool   `json:"ope_monitor" gorm:"default:true;comment:是否支持业务监控"`
}

func (Business) TableName() string {
	return "rs_business"
}

type Tag struct {
	RichGlobal
	Name string `json:"name" gorm:"index;type:varchar(50);comment:标签名称"`
}

func (Tag) TableName() string {
	return "rs_tag"
}

type BusinessCostCnf struct {
	Model
	BuId           int     `json:"buId"  gorm:"index;comment:业务ID"`
	Isp            int     `json:"isp" gorm:"type:int(1);default:1;comment:运营商"`
	Minimum        float64 `json:"minimum" gorm:"保底带宽(G)"`
	DialType       int     `json:"dialType" gorm:"type:int(1);default:0;comment:0:静态拨号 1:动态拨号"`
	IpType         int     `json:"ipType" gorm:"type:int(1);default:0;comment:0:ipv4 1:ipv6"`
	BandwidthLower float64 `json:"bandwidthLower" gorm:"comment:带宽下限(G)"`
	BandwidthLimit float64 `json:"bandwidthLimit" gorm:"comment:带宽上限(G)"`
	Price          float64 `json:"price" gorm:"comment: 价格(元/G/月)"`
}

func (BusinessCostCnf) TableName() string {
	return "rs_business_cost_cnf"
}

type OperationLog struct {
	Model

	CreatedAt  models.XTime `json:"createdAt" gorm:"comment:操作时间"`
	CreateUser string       `json:"createBy" gorm:"index;comment:操作人"`
	Module     string       `json:"module" gorm:"index;type:varchar(30);comment:模块信息"`
	ObjectId   int          `json:"objectId" gorm:"index;comment:操作的对象ID"`
	TargetId   int          `json:"targetId" gorm:"index;comment:操作的目标ID"`
	Action     string       `json:"action" gorm:"type:varchar(20);comment:操作名称"`
	Info       string       `json:"info" gorm:"comment:操作内容"`
}

func (OperationLog) TableName() string {
	return "rs_operation_log"
}
