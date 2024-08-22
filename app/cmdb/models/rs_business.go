package models

import (
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsBusiness struct {
	models.Model
	Status        int         `json:"status" gorm:"type:tinyint;comment:状态" comment:"状态"`
	Layer         int         `json:"layer" gorm:"type:tinyint;comment:排序"`
	Desc          string      `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Name          string      `json:"name" gorm:"type:varchar(50);comment:业务云名称"`
	ParentId      int         `json:"parentId" gorm:"comment:父业务"`
	BillingMethod int         `json:"billingMethod" gorm:"type:int(1);comment:业务计费方式"`
	EnName        string      `json:"enName" gorm:"index;type:varchar(30);comment:业务英文名字"`
	Children      interface{} `json:"children" gorm:"-"`

	CostCnf interface{} `json:"costCnf" gorm:"-"`
	models.ExtendUserBy
	models.ModelTime
	models.ControlBy
}

func (RsBusiness) TableName() string {
	return "rs_business"
}
func (e *RsBusiness) AfterFind(tx *gorm.DB) (err error) {
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

func (e *RsBusiness) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsBusiness) GetId() interface{} {
	return e.Id
}

type RsBusinessCostCnf struct {
	models.Model
	BuId           int     `json:"buId" gorm:"type:bigint;comment:业务ID"`
	Isp            int     `json:"isp" gorm:"type:int;comment:运营商"`
	Minimum        float64 `json:"minimum" gorm:"type:double;comment:Minimum"`
	DialType       int     `json:"dialType" gorm:"type:int;comment:0:静态拨号 1:动态拨号"`
	IpType         int     `json:"ipType" gorm:"type:int;comment:0:ipv4 1:ipv6"`
	BandwidthLower float64 `json:"bandwidthLower" gorm:"type:double;comment:带宽下限(G)"`
	BandwidthLimit float64 `json:"bandwidthLimit" gorm:"type:double;comment:带宽上限(G)"`
	Price          float64 `json:"price" gorm:"type:double;comment: 价格(元/G/月)"`
}

func (RsBusinessCostCnf) TableName() string {
	return "rs_business_cost_cnf"
}

func (e *RsBusinessCostCnf) GetId() interface{} {
	return e.Id
}
