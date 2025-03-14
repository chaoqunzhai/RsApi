package dto

import (
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsCustomGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:rs_custom" comment:"客户名称"`
	Type           int64  `form:"type"  search:"type:exact;column:type;table:rs_custom" comment:"客户类型,customer_type"`
	Cooperation    int64  `form:"cooperation"  search:"type:exact;column:cooperation;table:rs_custom" comment:"合作状态,work_status"`
	RsCustomOrder
}

type RsCustomOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:rs_custom"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_custom"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_custom"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_custom"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_custom"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_custom"`
	Desc        string `form:"descOrder"  search:"type:order;column:desc;table:rs_custom"`
	Name        string `form:"nameOrder"  search:"type:order;column:name;table:rs_custom"`
	Type        string `form:"typeOrder"  search:"type:order;column:type;table:rs_custom"`
	Cooperation string `form:"cooperationOrder"  search:"type:order;column:cooperation;table:rs_custom"`
	Region      string `form:"regionOrder"  search:"type:order;column:region;table:rs_custom"`
	Address     string `form:"addressOrder"  search:"type:order;column:address;table:rs_custom"`
}

func (m *RsCustomGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsCustomInsertReq struct {
	Id          int    `json:"-" comment:"主键编码"` // 主键编码
	Desc        string `json:"desc" comment:"描述信息"`
	Name        string `json:"name" comment:"客户名称"`
	Type        int    `json:"type" comment:"客户类型,customer_type"`
	Cooperation int    `json:"cooperation" comment:"合作状态,work_status"`
	Region      string `json:"region" comment:"所在地区"`
	Address     string `json:"address" comment:"地址"`
	common.ControlBy
}
type RsCustomIntegrationReq struct {
	//客户信息
	Desc        string `json:"desc" comment:"描述信息"`
	Name        string `json:"name" comment:"客户名称"`
	Type        int    `json:"type" comment:"客户类型,customer_type"`
	Cooperation int    `json:"cooperation" comment:"合作状态,work_status"`
	Region      string `json:"region" comment:"所在地区"`
	Address     string `json:"address" comment:"地址"`

	//联系人
	UserName string `json:"userName" comment:"姓名"`
	UserRegion string `json:"user_region"`
	BuId     int    `json:"buId" comment:"所属商务人员"`
	Phone    string `json:"phone" comment:"联系号码"`
	Email    string `json:"email" comment:"联系邮箱"`
	Dept     string `json:"dept" comment:"部门"`
	Duties   string `json:"duties" comment:"职务"`
	UserAddress  string `json:"user_address" comment:"详细地址"`


	//合同

	ContractName           string                     `json:"contract_name" comment:"合同名称"`
	ContractNumber         string                     `json:"contract_number" comment:"合同编号"`
	ContractSignatoryId    int                        `json:"contract_signatoryId" comment:"签订人"`
	UserId         int                        `json:"userId" comment:"联系人名称"`
	ContractType           int                        `json:"contract_type" comment:"合同类型,contract_type"`
	ContractSettlementType int                        `json:"contract_settlementType" comment:"结算方式,settlement_type"`
	ContractStartTimeAt    string                     `json:"contract_startTimeAt" comment:"合同开始时间"`
	ContractEndTimeAt      string                     `json:"contract_endTimeAt" comment:"合同结束时间"`
	ContractAccountName    string                     `json:"contract_accountName" comment:"开户名称"`
	ContractBankAccount    string                     `json:"contract_bankAccount" comment:"银行账号"`
	ContractBankName       string                     `json:"contract_bankName" comment:"开户银行"`
	ContractIdentifyNumber string                     `json:"contract_identifyNumber" comment:"纳税人识别号"`

	BandwidthFees  []RsBandwidthFeesInsertReq `json:"bandwidthFees"`
	common.ControlBy
}

type RsCustomIntegrationUpdateReq struct {
	Id          int    `json:"id" comment:"主键编码"` // 主键编码
	Desc        string `json:"desc" comment:"描述信息"`
	Name        string `json:"name" comment:"客户名称"`
	Type        int    `json:"type" comment:"客户类型,customer_type"`
	Cooperation int    `json:"cooperation" comment:"合作状态,work_status"`
	Region      string `json:"region" comment:"所在地区"`
	Address     string `json:"address" comment:"地址"`

	CustomUserId       int    `json:"custom_user_id" comment:"主键编码"` // 主键编码
	UserName string `json:"userName" comment:"姓名"`
	UserRegion string `json:"user_region"`
	BuId     int    `json:"buId" comment:"所属商务人员"`
	Phone    string `json:"phone" comment:"联系号码"`
	Email    string `json:"email" comment:"联系邮箱"`
	Dept     string `json:"dept" comment:"部门"`
	Duties   string `json:"duties" comment:"职务"`
	UserAddress  string `json:"user_address" comment:"详细地址"`

	ContractId       int    `json:"contract_id" comment:"主键编码"` // 主键编码
	ContractName           string                     `json:"contract_name" comment:"合同名称"`
	ContractNumber         string                     `json:"contract_number" comment:"合同编号"`
	ContractSignatoryId    int                        `json:"contract_signatoryId" comment:"签订人"`
	UserId         int                        `json:"userId" comment:"联系人名称"`
	ContractType           int                        `json:"contract_type" comment:"合同类型,contract_type"`
	ContractSettlementType int                        `json:"contract_settlementType" comment:"结算方式,settlement_type"`
	ContractStartTimeAt    string                     `json:"contract_startTimeAt" comment:"合同开始时间"`
	ContractEndTimeAt      string                     `json:"contract_endTimeAt" comment:"合同结束时间"`
	ContractAccountName    string                     `json:"contract_accountName" comment:"开户名称"`
	ContractBankAccount    string                     `json:"contract_bankAccount" comment:"银行账号"`
	ContractBankName       string                     `json:"contract_bankName" comment:"开户银行"`
	ContractIdentifyNumber string                     `json:"contract_identifyNumber" comment:"纳税人识别号"`
	BandwidthFees  []RsBandwidthFeesInsertReq `json:"bandwidthFees" gorm:"-"`
	common.ControlBy
}


func (s *RsCustomInsertReq) Generate(model *models.RsCustom) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Desc = s.Desc
	model.Name = s.Name
	model.Type = s.Type
	model.Cooperation = s.Cooperation
	model.Region = s.Region
	model.Address = s.Address
}

func (s *RsCustomInsertReq) GetId() interface{} {
	return s.Id
}

type RsCustomUpdateReq struct {
	Id          int    `uri:"id" comment:"主键编码"` // 主键编码
	Desc        string `json:"desc" comment:"描述信息"`
	Name        string `json:"name" comment:"客户名称"`
	Type        int    `json:"type" comment:"客户类型,customer_type"`
	Cooperation int    `json:"cooperation" comment:"合作状态,work_status"`
	Region      string `json:"region" comment:"所在地区"`
	Address     string `json:"address" comment:"地址"`
	common.ControlBy
}

func (s *RsCustomUpdateReq) Generate(model *models.RsCustom) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.Name = s.Name
	model.Type = s.Type
	model.Cooperation = s.Cooperation
	model.Region = s.Region
	model.Address = s.Address
}

func (s *RsCustomUpdateReq) GetId() interface{} {
	return s.Id
}

// RsCustomGetReq 功能获取请求参数
type RsCustomGetReq struct {
	Id int `uri:"id"`
}

func (s *RsCustomGetReq) GetId() interface{} {
	return s.Id
}

// RsCustomDeleteReq 功能删除请求参数
type RsCustomDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsCustomDeleteReq) GetId() interface{} {
	return s.Ids
}
