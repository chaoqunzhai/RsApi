package models

import "database/sql"

type Contract struct {
	RichGlobal
	Name           string       `json:"name" gorm:"type:varchar(100);comment:合同名称"`
	Number         string       `json:"number" gorm:"type:varchar(100);comment:合同编号"`
	BuId           int          `json:"buId" gorm:"comment:商务人员"`
	CustomId       int          `json:"customId" gorm:"comment:所属客户ID"`
	SignatoryId    int          `json:"signatoryId" gorm:"comment:签订人ID"`
	UserId         int          `json:"userId" gorm:"comment:联系人ID"`
	Type           int          `json:"type" gorm:"comment:合同类型,contract_type"`
	SettlementType int          `json:"settlementType" gorm:"comment:结算方式,settlement_type"`
	StartTime      sql.NullTime `json:"startTime" gorm:"comment:合同开始时间"`
	EndTime        sql.NullTime `json:"endTime" gorm:"comment:合同结束时间"`
	AccountName    string       `json:"accountName"  gorm:"type:varchar(100);comment:开户名称"`
	BankAccount    string       `json:"bankAccount"  gorm:"type:varchar(100);comment:银行账号"`
	BankName       string       `json:"bankName"  gorm:"type:varchar(100);comment:开户银行"`
	IdentifyNumber string       `json:"identifyNumber" gorm:"type:varchar(100);comment:纳税人识别号"`
	Address        string       `json:"address"  gorm:"type:varchar(200);comment:地址"`
	Phone          string       `json:"phone"  gorm:"type:varchar(30);comment:电话"`
}

func (Contract) TableName() string {
	return "rs_contract"
}

type BandwidthFees struct {
	Model
	ModelTime
	Region          string  `json:"region" gorm:"type:varchar(120);comment:所在地区"`
	ContractId      int     `json:"contractId" gorm:"index;type:int(11);comment:关联的合同"`
	Isp             int     `json:"isp" gorm:"type:int(1);default:1;comment:运营商"`
	Up              float64 `json:"up" gorm:"default:0;comment:上行带宽"`
	Down            float64 `json:"down" gorm:"default:0;comment:下行带宽"`
	LinePrice       float64 `json:"LinePrice" gorm:"comment:业务线单价"`
	ManagerLineCost float64 `json:"managerLineCost" gorm:"comment:管理线价格"`
	Charging        int     `json:"charging" gorm:"type:int(1);default:0;comment:计费方式"`
	TransProvince   bool    `json:"transProd" gorm:"default:false;comment:是否跨省"`
	MoreDialing     bool    `json:"moreDialing" gorm:"default:false;comment:是否支持多拨"`
}

func (BandwidthFees) TableName() string {
	return "rs_bandwidth_fees"
}
