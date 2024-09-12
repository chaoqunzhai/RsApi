package dto

import (
	"time"

	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetDisposalGetPageReq struct {
	dto.Pagination `search:"-"`
	AssetDisposalOrder
}

type AssetDisposalOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:asset_disposal"`
	AssetId        string `form:"assetIdOrder"  search:"type:order;column:asset_id;table:asset_disposal"`
	DisposalPerson string `form:"disposalPersonOrder"  search:"type:order;column:disposal_person;table:asset_disposal"`
	Reason         string `form:"reasonOrder"  search:"type:order;column:reason;table:asset_disposal"`
	DisposalWay    string `form:"disposalWayOrder"  search:"type:order;column:disposal_way;table:asset_disposal"`
	DisposalType   string `form:"disposalTypeOrder"  search:"type:order;column:disposal_type;table:asset_disposal"`
	LocationId     string `form:"locationIdOrder"  search:"type:order;column:location_id;table:asset_disposal"`
	Amount         string `form:"amountOrder"  search:"type:order;column:amount;table:asset_disposal"`
	DisposalAt     string `form:"disposalAtOrder"  search:"type:order;column:disposal_at;table:asset_disposal"`
	Remark         string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_disposal"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_disposal"`
	UpdatedAt      string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_disposal"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_disposal"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_disposal"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_disposal"`
}

func (m *AssetDisposalGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetDisposalInsertReq struct {
	Id             int       `json:"-" comment:"主键"` // 主键
	AssetId        int       `json:"assetId" comment:"资产编码"`
	DisposalPerson int       `json:"disposalPerson" comment:"处置人编码"`
	Reason         string    `json:"reason" comment:"处置原因"`
	DisposalWay    int8      `json:"disposalWay" comment:"处置方式(1=报废, 2=出售, 3=出租, 4=退租, 5=捐赠, 6=其它)"`
	DisposalType   int8      `json:"disposalType" comment:"处置地点类型(1=机房, 2=库房)"`
	LocationId     int       `json:"locationId" comment:"处置地点编码(机房编码/库房编码)"`
	Amount         float64   `json:"amount" comment:"处置金额"`
	DisposalAt     time.Time `json:"disposalAt" comment:"处置时间"`
	Remark         string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetDisposalInsertReq) Generate(model *models.AssetDisposal) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.DisposalPerson = s.DisposalPerson
	model.Reason = s.Reason
	model.DisposalWay = s.DisposalWay
	model.DisposalType = s.DisposalType
	model.LocationId = s.LocationId
	model.Amount = s.Amount
	model.DisposalAt = s.DisposalAt
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetDisposalInsertReq) GetId() interface{} {
	return s.Id
}

type AssetDisposalUpdateReq struct {
	Id             int       `uri:"id" comment:"主键"` // 主键
	AssetId        int       `json:"assetId" comment:"资产编码"`
	DisposalPerson int       `json:"disposalPerson" comment:"处置人编码"`
	Reason         string    `json:"reason" comment:"处置原因"`
	DisposalWay    int8      `json:"disposalWay" comment:"处置方式(1=报废, 2=出售, 3=出租, 4=退租, 5=捐赠, 6=其它)"`
	DisposalType   int8      `json:"disposalType" comment:"处置地点类型(1=机房, 2=库房)"`
	LocationId     int       `json:"locationId" comment:"处置地点编码(机房编码/库房编码)"`
	Amount         float64   `json:"amount" comment:"处置金额"`
	DisposalAt     time.Time `json:"disposalAt" comment:"处置时间"`
	Remark         string    `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetDisposalUpdateReq) Generate(model *models.AssetDisposal) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AssetId = s.AssetId
	model.DisposalPerson = s.DisposalPerson
	model.Reason = s.Reason
	model.DisposalWay = s.DisposalWay
	model.DisposalType = s.DisposalType
	model.LocationId = s.LocationId
	model.Amount = s.Amount
	model.DisposalAt = s.DisposalAt
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetDisposalUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetDisposalGetReq 功能获取请求参数
type AssetDisposalGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetDisposalGetReq) GetId() interface{} {
	return s.Id
}

// AssetDisposalDeleteReq 功能删除请求参数
type AssetDisposalDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetDisposalDeleteReq) GetId() interface{} {
	return s.Ids
}
