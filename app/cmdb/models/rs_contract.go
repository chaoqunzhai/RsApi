package models

import (
	"database/sql"
	"go-admin/common/models"
)

type RsContract struct {
	models.Model

	Desc           string       `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Name           string       `json:"name" gorm:"type:varchar(20);comment:合同名称"`
	Number         string       `json:"number" gorm:"type:varchar(20);comment:合同编号"`
	BuId           int          `json:"buId" gorm:"type:bigint;comment:商务人员"`
	CustomId       int          `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	SignatoryId    int          `json:"signatoryId" gorm:"type:bigint;comment:签订人"`
	User           string       `json:"user" gorm:"type:varchar(20);comment:联系人名称"`
	Type           int          `json:"type" gorm:"type:bigint;comment:合同类型,contract_type"`
	SettlementType int          `json:"settlementType" gorm:"type:bigint;comment:结算方式,settlement_type"`
	StartTime      sql.NullTime `json:"startTime" gorm:"type:datetime(3);comment:合同开始时间"`
	EndTime        sql.NullTime `json:"endTime" gorm:"type:datetime(3);comment:合同结束时间"`
	AccountName    string       `json:"accountName" gorm:"type:varchar(30);comment:开户名称"`
	BankAccount    string       `json:"bankAccount" gorm:"type:varchar(50);comment:银行账号"`
	BankName       string       `json:"bankName" gorm:"type:varchar(30);comment:开户银行"`
	IdentifyNumber string       `json:"identifyNumber" gorm:"type:varchar(30);comment:纳税人识别号"`
	Address        string       `json:"address" gorm:"type:varchar(120);comment:地址"`
	Phone          string       `json:"phone" gorm:"type:varchar(30);comment:电话"`
	models.ModelTime
	models.ControlBy
}

func (RsContract) TableName() string {
	return "rs_contract"
}

func (e *RsContract) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsContract) GetId() interface{} {
	return e.Id
}
