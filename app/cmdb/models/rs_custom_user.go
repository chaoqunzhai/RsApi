package models

import (
	"database/sql"
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsCustomUser struct {
	models.Model

	Desc     string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	UserName string `json:"userName" gorm:"type:varchar(50);comment:姓名"`
	CustomId int    `json:"customId" gorm:"type:bigint;comment:所属客户"`
	BuId     int    `json:"buId" gorm:"type:bigint;comment:所属商务人员"`
	Phone    string `json:"phone" gorm:"type:varchar(20);comment:联系号码"`
	Email    string `json:"email" gorm:"type:varchar(50);comment:联系邮箱"`
	Region   string `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	Dept     string `json:"dept" gorm:"type:varchar(30);comment:部门"`
	Duties   string `json:"duties" gorm:"type:varchar(30);comment:职务"`
	Address  string `json:"address" gorm:"type:varchar(255);comment:详细地址"`
	models.ModelTime
	models.ControlBy
}


func (RsCustomUser) TableName() string {
	return "rs_custom_user"
}

func (e *RsCustomUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsCustomUser) GetId() interface{} {
	return e.Id
}
type RsCustomUser2 struct {
	models.Model
	Id int  `json:"custom_user_id"`
	Desc     string `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	UserName string `json:"userName" gorm:"type:varchar(50);comment:姓名"`
	CustomId int    `json:"customId" gorm:"type:bigint;comment:所属客户"`
	BuId     int    `json:"buId" gorm:"type:bigint;comment:所属商务人员"`
	Phone    string `json:"phone" gorm:"type:varchar(20);comment:联系号码"`
	Email    string `json:"email" gorm:"type:varchar(50);comment:联系邮箱"`
	Region   string `json:"user_region" gorm:"type:varchar(100);comment:省份城市多ID"`
	Dept     string `json:"dept" gorm:"type:varchar(30);comment:部门"`
	Duties   string `json:"duties" gorm:"type:varchar(30);comment:职务"`
	Address  string `json:"user_address" gorm:"type:varchar(255);comment:详细地址"`
	models.ModelTime
	models.ControlBy
}
func (RsCustomUser2) TableName() string {
	return "rs_custom_user"
}

type RsContract2 struct {
	models.Model
	Id int  `json:"contract_id"`
	Name           string            `json:"contract_name" gorm:"type:varchar(20);comment:合同名称"`
	Number         string            `json:"contract_number" gorm:"type:varchar(20);comment:合同编号"`
	SignatoryId    int               `json:"contract_signatoryId" gorm:"type:bigint;comment:签订人"`
	Type           int               `json:"contract_type" gorm:"type:bigint;comment:合同类型,contract_type"`
	SettlementType int               `json:"contract_settlementType" gorm:"type:bigint;comment:结算方式,settlement_type"`
	StartTime      sql.NullTime      `json:"-" gorm:"type:datetime(3);comment:合同开始时间"`
	EndTime        sql.NullTime      `json:"-" gorm:"type:datetime(3);comment:合同结束时间"`
	AccountName    string            `json:"contract_accountName" gorm:"type:varchar(30);comment:开户名称"`
	BankAccount    string            `json:"contract_bankAccount" gorm:"type:varchar(50);comment:银行账号"`
	BankName       string            `json:"contract_bankName" gorm:"type:varchar(30);comment:开户银行"`
	IdentifyNumber string            `json:"contract_identifyNumber" gorm:"type:varchar(30);comment:纳税人识别号"`
	StartTimeAt    string            `json:"contract_startTimeAt" gorm:"-"`
	EndTimeAt      string            `json:"contract_endTimeAt" gorm:"-"`
	BandwidthFees  []RsBandwidthFees `json:"bandwidthFees" gorm:"-"`
	models.ModelTime
	models.ControlBy
}
func (RsContract2) TableName() string {
	return "rs_contract"
}
func (e *RsContract2) AfterFind(tx *gorm.DB) (err error) {

	if e.StartTime.Valid {
		e.StartTimeAt = e.StartTime.Time.Format("2006-01-02")

	}
	if e.EndTime.Valid {
		e.EndTimeAt = e.EndTime.Time.Format("2006-01-02")

	}
	var BandwidthFees []RsBandwidthFees

	tx.Model(&RsBandwidthFees{}).Where("contract_id = ?", e.Id).Find(&BandwidthFees)
	e.BandwidthFees = BandwidthFees

	return nil
}