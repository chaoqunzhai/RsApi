package models

import (
	"database/sql"
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsContract struct {
	models.Model

	Desc           string            `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	Name           string            `json:"name" gorm:"type:varchar(20);comment:合同名称"`
	Number         string            `json:"number" gorm:"type:varchar(20);comment:合同编号"`
	BuId           int               `json:"buId" gorm:"type:bigint;comment:商务人员"`
	CustomId       int               `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	SignatoryId    int               `json:"signatoryId" gorm:"type:bigint;comment:签订人"`
	UserId         int               `json:"userId" gorm:"comment:联系人名称"`
	Type           int               `json:"type" gorm:"type:bigint;comment:合同类型,contract_type"`
	SettlementType int               `json:"settlementType" gorm:"type:bigint;comment:结算方式,settlement_type"`
	StartTime      sql.NullTime      `json:"-" gorm:"type:datetime(3);comment:合同开始时间"`
	EndTime        sql.NullTime      `json:"-" gorm:"type:datetime(3);comment:合同结束时间"`
	AccountName    string            `json:"accountName" gorm:"type:varchar(30);comment:开户名称"`
	BankAccount    string            `json:"bankAccount" gorm:"type:varchar(50);comment:银行账号"`
	BankName       string            `json:"bankName" gorm:"type:varchar(30);comment:开户银行"`
	IdentifyNumber string            `json:"identifyNumber" gorm:"type:varchar(30);comment:纳税人识别号"`
	Address        string            `json:"address" gorm:"type:varchar(120);comment:地址"`
	Phone          string            `json:"phone" gorm:"type:varchar(30);comment:电话"`
	StartTimeAt    string            `json:"startTimeAt" gorm:"-"`
	EndTimeAt      string            `json:"endTimeAt" gorm:"-"`
	BandwidthFees  []RsBandwidthFees `json:"bandwidthFees" gorm:"-"`
	//BuInfo         map[string]interface{} `json:"buInfo" gorm:"-"`
	//CustomInfo     map[string]interface{} `json:"customInfo" gorm:"-"`
	//SignatoryInfo  map[string]interface{} `json:"signatoryInfo" gorm:"-"`
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

func (e *RsContract) AfterFind(tx *gorm.DB) (err error) {

	if e.StartTime.Valid {
		e.StartTimeAt = e.StartTime.Time.Format("2006-01-02 15:04:05")

	}
	if e.EndTime.Valid {
		e.EndTimeAt = e.EndTime.Time.Format("2006-01-02 15:04:05")

	}
	var BandwidthFees []RsBandwidthFees

	tx.Model(&RsBandwidthFees{}).Where("contract_id = ?", e.Id).Find(&BandwidthFees)
	e.BandwidthFees = BandwidthFees

	return nil
}

type RsBandwidthFees struct {
	models.Model
	ContractId      int     `json:"-" gorm:"index;type:int(11);comment:关联的合同"`
	Isp             int     `json:"isp" gorm:"type:int(1);default:1;comment:运营商"`
	Up              float64 `json:"up" gorm:"default:0;comment:上行带宽"`
	Down            float64 `json:"down" gorm:"default:0;comment:下行带宽"`
	LinePrice       float64 `json:"linePrice" gorm:"comment:业务线单价"`
	ManagerLineCost float64 `json:"managerLineCost" gorm:"comment:管理线价格"`
	Charging        int     `json:"charging" gorm:"type:int(1);default:0;comment:计费方式"`
	TransProvince   int     `json:"transProd" gorm:"default:0;comment:是否跨省"`
	MoreDialing     int     `json:"moreDialing" gorm:"default:0;comment:是否支持多拨"`
	models.ModelTime
}

func (RsBandwidthFees) TableName() string {
	return "rs_bandwidth_fees"
}

func (e *RsBandwidthFees) GetId() interface{} {
	return e.Id
}
