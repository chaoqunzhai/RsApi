package dto

import (
	"time"

	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RsHostIncomeGetPageReq struct {
	dto.Pagination `search:"-"`
	HostId         string `form:"hostId"  search:"type:exact;column:host_id;table:rs_host_income" comment:"主机ID"`
	Isp            string `form:"isp"  search:"type:exact;column:isp;table:rs_host_income" comment:"运营商ID"`
	IdcId          string `form:"idcId"  search:"type:exact;column:idc_id;table:rs_host_income" comment:"IDC ID"`
	BusinessId     string `form:"businessId"  search:"-" comment:"业务ID"`
	Usage          string `form:"usage"  search:"type:exact;column:usage;table:rs_host_income" comment:""`
	SettleStatus   string `form:"settleStatus"  search:"type:exact;column:settle_status;table:rs_host_income" comment:""`
	StartTimeAt    string `form:"startTime"  search:"type:gte;column:created_at;table:rs_host_income" comment:"开始时间"`
	EndTimeAt      string `form:"endTime"  search:"type:lte;column:created_at;table:rs_host_income" comment:"结束时间"`
	RsHostIncomeOrder
}

type RsHostIncomeOrder struct {
	Id                string `form:"idOrder"  search:"type:order;column:id;table:rs_host_income"`
	CreatedAt         string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_host_income"`
	HostId            string `form:"hostIdOrder"  search:"type:order;column:host_id;table:rs_host_income"`
	Isp               string `form:"ispOrder"  search:"type:order;column:isp;table:rs_host_income"`
	IdcId             string `form:"idcIdOrder"  search:"type:order;column:idc_id;table:rs_host_income"`
	BuId              string `form:"buIdOrder"  search:"type:order;column:bu_id;table:rs_host_income"`
	Income            string `form:"incomeOrder"  search:"type:order;column:income;table:rs_host_income"`
	Usage             string `form:"usageOrder"  search:"type:order;column:usage;table:rs_host_income"`
	Bandwidth95       string `form:"bandwidth95Order"  search:"type:order;column:bandwidth95;table:rs_host_income"`
	BandwidthIncome   string `form:"bandwidthIncomeOrder"  search:"type:order;column:bandwidth_income;table:rs_host_income"`
	Estimate          string `form:"estimateOrder"  search:"type:order;column:estimate;table:rs_host_income"`
	Actual            string `form:"actualOrder"  search:"type:order;column:actual;table:rs_host_income"`
	SlaPrice          string `form:"slaPriceOrder"  search:"type:order;column:sla_price;table:rs_host_income"`
	SlaInfo           string `form:"slaInfoOrder"  search:"type:order;column:sla_info;table:rs_host_income"`
	SettleStatus      string `form:"settleStatusOrder"  search:"type:order;column:settle_status;table:rs_host_income"`
	SettleTime        string `form:"settleTimeOrder"  search:"type:order;column:settle_time;table:rs_host_income"`
	SettleBandwidth   string `form:"settleBandwidthOrder"  search:"type:order;column:settle_bandwidth;table:rs_host_income"`
	TotalBandwidth    string `form:"totalBandwidthOrder"  search:"type:order;column:total_bandwidth;table:rs_host_income"`
	HeartbeatNum      string `form:"heartbeatNumOrder"  search:"type:order;column:heartbeat_num;table:rs_host_income"`
	NightHeartbeatNum string `form:"nightHeartbeatNumOrder"  search:"type:order;column:night_heartbeat_num;table:rs_host_income"`
}

func (m *RsHostIncomeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsHostIncomeInsertReq struct {
	Id                int       `json:"-" comment:"主键编码"` // 主键编码
	HostId            string    `json:"hostId" comment:"主机ID"`
	Isp               string    `json:"isp" comment:"运营商ID"`
	IdcId             string    `json:"idcId" comment:"IDC ID"`
	BuId              string    `json:"buId" comment:"业务ID"`
	Income            string    `json:"income" comment:""`
	Usage             string    `json:"usage" comment:""`
	Bandwidth95       string    `json:"bandwidth95" comment:""`
	BandwidthIncome   string    `json:"bandwidthIncome" comment:""`
	Estimate          string    `json:"estimate" comment:""`
	Actual            string    `json:"actual" comment:""`
	SlaPrice          string    `json:"slaPrice" comment:""`
	SlaInfo           string    `json:"slaInfo" comment:""`
	SettleStatus      string    `json:"settleStatus" comment:""`
	SettleTime        time.Time `json:"settleTime" comment:""`
	SettleBandwidth   string    `json:"settleBandwidth" comment:""`
	TotalBandwidth    string    `json:"totalBandwidth" comment:""`
	HeartbeatNum      string    `json:"heartbeatNum" comment:""`
	NightHeartbeatNum string    `json:"nightHeartbeatNum" comment:""`
	common.ControlBy
}

func (s *RsHostIncomeInsertReq) Generate(model *models.RsHostIncome) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.HostId = s.HostId
	model.Isp = s.Isp
	model.IdcId = s.IdcId
	model.BuId = s.BuId
	model.Income = s.Income
	model.Usage = s.Usage
	model.Bandwidth95 = s.Bandwidth95
	model.BandwidthIncome = s.BandwidthIncome
	model.Estimate = s.Estimate

	model.SlaPrice = s.SlaPrice
	model.SlaInfo = s.SlaInfo
	model.SettleStatus = s.SettleStatus
	model.SettleTime = s.SettleTime
	model.SettleBandwidth = s.SettleBandwidth
	model.TotalBandwidth = s.TotalBandwidth
	model.HeartbeatNum = s.HeartbeatNum
	model.NightHeartbeatNum = s.NightHeartbeatNum
}

func (s *RsHostIncomeInsertReq) GetId() interface{} {
	return s.Id
}

type RsHostIncomeUpdateReq struct {
	Id                int       `uri:"id" comment:"主键编码"` // 主键编码
	HostId            string    `json:"hostId" comment:"主机ID"`
	Isp               string    `json:"isp" comment:"运营商ID"`
	IdcId             string    `json:"idcId" comment:"IDC ID"`
	BuId              string    `json:"buId" comment:"业务ID"`
	Income            string    `json:"income" comment:""`
	Usage             string    `json:"usage" comment:""`
	Bandwidth95       string    `json:"bandwidth95" comment:""`
	BandwidthIncome   string    `json:"bandwidthIncome" comment:""`
	Estimate          string    `json:"estimate" comment:""`
	Actual            string    `json:"actual" comment:""`
	SlaPrice          string    `json:"slaPrice" comment:""`
	SlaInfo           string    `json:"slaInfo" comment:""`
	SettleStatus      string    `json:"settleStatus" comment:""`
	SettleTime        time.Time `json:"settleTime" comment:""`
	SettleBandwidth   string    `json:"settleBandwidth" comment:""`
	TotalBandwidth    string    `json:"totalBandwidth" comment:""`
	HeartbeatNum      string    `json:"heartbeatNum" comment:""`
	NightHeartbeatNum string    `json:"nightHeartbeatNum" comment:""`
	common.ControlBy
}

func (s *RsHostIncomeUpdateReq) Generate(model *models.RsHostIncome) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.HostId = s.HostId
	model.Isp = s.Isp
	model.IdcId = s.IdcId
	model.BuId = s.BuId
	model.Income = s.Income
	model.Usage = s.Usage
	model.Bandwidth95 = s.Bandwidth95
	model.BandwidthIncome = s.BandwidthIncome
	model.Estimate = s.Estimate
	model.SlaPrice = s.SlaPrice
	model.SlaInfo = s.SlaInfo
	model.SettleStatus = s.SettleStatus
	model.SettleTime = s.SettleTime
	model.SettleBandwidth = s.SettleBandwidth
	model.TotalBandwidth = s.TotalBandwidth
	model.HeartbeatNum = s.HeartbeatNum
	model.NightHeartbeatNum = s.NightHeartbeatNum
}

func (s *RsHostIncomeUpdateReq) GetId() interface{} {
	return s.Id
}

// RsHostIncomeGetReq 功能获取请求参数
type RsHostIncomeGetReq struct {
	Id int `uri:"id"`
}

func (s *RsHostIncomeGetReq) GetId() interface{} {
	return s.Id
}

// RsHostIncomeDeleteReq 功能删除请求参数
type RsHostIncomeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsHostIncomeDeleteReq) GetId() interface{} {
	return s.Ids
}
