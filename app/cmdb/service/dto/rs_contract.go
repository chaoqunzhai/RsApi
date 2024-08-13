package dto

import (
	"database/sql"
	"go-admin/global"
	"time"

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsContractGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string    `form:"name"  search:"type:contains;column:name;table:rs_contract" comment:"合同名称"`
	Number         string    `form:"number"  search:"type:exact;column:number;table:rs_contract" comment:"合同编号"`
	BuId           int       `form:"buId"  search:"type:exact;column:bu_id;table:rs_contract" comment:"商务人员"`
	CustomId       int       `form:"customId"  search:"type:exact;column:custom_id;table:rs_contract" comment:"所属客户ID"`
	SignatoryId    int       `form:"signatoryId"  search:"type:exact;column:signatory_id;table:rs_contract" comment:"签订人"`
	Type           int       `form:"type"  search:"type:exact;column:type;table:rs_contract" comment:"合同类型,contract_type"`
	SettlementType int       `form:"settlementType"  search:"type:exact;column:settlement_type;table:rs_contract" comment:"结算方式,settlement_type"`
	StartTime      time.Time `form:"startTime"  search:"type:exact;column:start_time;table:rs_contract" comment:"合同开始时间"`
	EndTime        time.Time `form:"endTime"  search:"type:exact;column:end_time;table:rs_contract" comment:"合同结束时间"`
	Address        string    `form:"address"  search:"type:exact;column:address;table:rs_contract" comment:"地址"`
	Phone          string    `form:"phone"  search:"type:contains;column:phone;table:rs_contract" comment:"电话"`
	RsContractOrder
}

type RsContractOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:rs_contract"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_contract"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_contract"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_contract"`
	UpdatedAt      string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_contract"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_contract"`
	Desc           string `form:"descOrder"  search:"type:order;column:desc;table:rs_contract"`
	Name           string `form:"nameOrder"  search:"type:order;column:name;table:rs_contract"`
	Number         string `form:"numberOrder"  search:"type:order;column:number;table:rs_contract"`
	BuId           string `form:"buIdOrder"  search:"type:order;column:bu_id;table:rs_contract"`
	CustomId       string `form:"customIdOrder"  search:"type:order;column:custom_id;table:rs_contract"`
	SignatoryId    string `form:"signatoryIdOrder"  search:"type:order;column:signatory_id;table:rs_contract"`
	User           string `form:"userOrder"  search:"type:order;column:user;table:rs_contract"`
	Type           string `form:"typeOrder"  search:"type:order;column:type;table:rs_contract"`
	SettlementType string `form:"settlementTypeOrder"  search:"type:order;column:settlement_type;table:rs_contract"`
	StartTime      string `form:"startTimeOrder"  search:"type:order;column:start_time;table:rs_contract"`
	EndTime        string `form:"endTimeOrder"  search:"type:order;column:end_time;table:rs_contract"`
	AccountName    string `form:"accountNameOrder"  search:"type:order;column:account_name;table:rs_contract"`
	BankAccount    string `form:"bankAccountOrder"  search:"type:order;column:bank_account;table:rs_contract"`
	BankName       string `form:"bankNameOrder"  search:"type:order;column:bank_name;table:rs_contract"`
	IdentifyNumber string `form:"identifyNumberOrder"  search:"type:order;column:identify_number;table:rs_contract"`
	Address        string `form:"addressOrder"  search:"type:order;column:address;table:rs_contract"`
	Phone          string `form:"phoneOrder"  search:"type:order;column:phone;table:rs_contract"`
}

func (m *RsContractGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsContractInsertReq struct {
	Id             int    `json:"-" comment:"主键编码"` // 主键编码
	Desc           string `json:"desc" comment:"描述信息"`
	Name           string `json:"name" comment:"合同名称"`
	Number         string `json:"number" comment:"合同编号"`
	BuId           int    `json:"buId" comment:"商务人员"`
	CustomId       int    `json:"customId" comment:"所属客户ID"`
	SignatoryId    int    `json:"signatoryId" comment:"签订人"`
	User           string `json:"user" comment:"联系人名称"`
	Type           int    `json:"type" comment:"合同类型,contract_type"`
	SettlementType int    `json:"settlementType" comment:"结算方式,settlement_type"`
	StartTime      string `json:"startTime" comment:"合同开始时间"`
	EndTime        string `json:"endTime" comment:"合同结束时间"`
	AccountName    string `json:"accountName" comment:"开户名称"`
	BankAccount    string `json:"bankAccount" comment:"银行账号"`
	BankName       string `json:"bankName" comment:"开户银行"`
	IdentifyNumber string `json:"identifyNumber" comment:"纳税人识别号"`
	Address        string `json:"address" comment:"地址"`
	Phone          string `json:"phone" comment:"电话"`
	common.ControlBy
}

func (s *RsContractInsertReq) Generate(model *models.RsContract) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc
	model.Name = s.Name
	model.Number = s.Number
	model.BuId = s.BuId
	model.CustomId = s.CustomId
	model.SignatoryId = s.SignatoryId
	model.User = s.User
	model.Type = s.Type
	model.SettlementType = s.SettlementType

	if s.StartTime != "" {
		if star, err := time.ParseInLocation(time.DateTime, s.StartTime, global.LOC); err == nil {
			model.StartTime = sql.NullTime{
				Time:  star,
				Valid: true,
			}
		}

	} else {
		model.StartTime = sql.NullTime{}
	}

	if s.EndTime != "" {
		if end, err := time.ParseInLocation(time.DateTime, s.StartTime, global.LOC); err == nil {
			model.EndTime = sql.NullTime{
				Time:  end,
				Valid: true,
			}
		}

	} else {
		model.StartTime = sql.NullTime{}
	}
	model.AccountName = s.AccountName
	model.BankAccount = s.BankAccount
	model.BankName = s.BankName
	model.IdentifyNumber = s.IdentifyNumber
	model.Address = s.Address
	model.Phone = s.Phone
}

func (s *RsContractInsertReq) GetId() interface{} {
	return s.Id
}

type RsContractUpdateReq struct {
	Id             int    `uri:"id" comment:"主键编码"` // 主键编码
	Desc           string `json:"desc" comment:"描述信息"`
	Name           string `json:"name" comment:"合同名称"`
	Number         string `json:"number" comment:"合同编号"`
	BuId           int    `json:"buId" comment:"商务人员"`
	CustomId       int    `json:"customId" comment:"所属客户ID"`
	SignatoryId    int    `json:"signatoryId" comment:"签订人"`
	User           string `json:"user" comment:"联系人名称"`
	Type           int    `json:"type" comment:"合同类型,contract_type"`
	SettlementType int    `json:"settlementType" comment:"结算方式,settlement_type"`
	StartTime      string `json:"startTime" comment:"合同开始时间"`
	EndTime        string `json:"endTime" comment:"合同结束时间"`
	AccountName    string `json:"accountName" comment:"开户名称"`
	BankAccount    string `json:"bankAccount" comment:"银行账号"`
	BankName       string `json:"bankName" comment:"开户银行"`
	IdentifyNumber string `json:"identifyNumber" comment:"纳税人识别号"`
	Address        string `json:"address" comment:"地址"`
	Phone          string `json:"phone" comment:"电话"`
	common.ControlBy
}

func (s *RsContractUpdateReq) Generate(model *models.RsContract) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.Name = s.Name
	model.Number = s.Number
	model.BuId = s.BuId
	model.CustomId = s.CustomId
	model.SignatoryId = s.SignatoryId
	model.User = s.User
	model.Type = s.Type
	model.SettlementType = s.SettlementType

	if s.StartTime != "" {
		if star, err := time.ParseInLocation(time.DateTime, s.StartTime, global.LOC); err == nil {
			model.StartTime = sql.NullTime{
				Time:  star,
				Valid: true,
			}
		}

	} else {
		model.StartTime = sql.NullTime{}
	}

	if s.EndTime != "" {
		if end, err := time.ParseInLocation(time.DateTime, s.StartTime, global.LOC); err == nil {
			model.EndTime = sql.NullTime{
				Time:  end,
				Valid: true,
			}
		}

	} else {
		model.StartTime = sql.NullTime{}
	}
	model.AccountName = s.AccountName
	model.BankAccount = s.BankAccount
	model.BankName = s.BankName
	model.IdentifyNumber = s.IdentifyNumber
	model.Address = s.Address
	model.Phone = s.Phone
}

func (s *RsContractUpdateReq) GetId() interface{} {
	return s.Id
}

// RsContractGetReq 功能获取请求参数
type RsContractGetReq struct {
	Id int `uri:"id"`
}

func (s *RsContractGetReq) GetId() interface{} {
	return s.Id
}

// RsContractDeleteReq 功能删除请求参数
type RsContractDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsContractDeleteReq) GetId() interface{} {
	return s.Ids
}
