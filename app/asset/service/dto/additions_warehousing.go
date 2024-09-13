package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AdditionsOrderGetPageReq struct {
	dto.Pagination `search:"-"`
	StoreRoomId    string `form:"storeRoomId"  search:"type:exact;column:storeroom_Id;table:additions_order" comment:"关联的库房ID"`
	OrderId        string `form:"orderId"  search:"type:contains;column:order_id;table:additions_order" comment:"关联的入库单号"`
	Name           string `form:"name"  search:"-" comment:"资产名称"`
	StartTimeAt    string `form:"startTimeAt"  search:"type:gte;column:created_at;table:additions_order" comment:"入库开始时间"`
	EndTimeAt      string `form:"endTimeAt"  search:"type:lte;column:created_at;table:additions_order" comment:"入库结束时间"`
}

func (m *AdditionsOrderGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AdditionsWarehousingGetPageReq struct {
	dto.Pagination `search:"-"`
	CategoryId     int64  `form:"categoryId"  search:"type:exact;column:category_id;table:additions_warehousing" comment:"关联的资产分类ID"`
	StoreRoomId    int64  `form:"storeRoomId"  search:"type:exact;column:storeroom_Id;table:additions_warehousing" comment:"关联的库房ID"`
	SupplierId     int64  `form:"supplierId"  search:"type:exact;column:supplier_id;table:additions_warehousing" comment:"供应商ID"`
	WId            int64  `form:"wId"  search:"type:exact;column:w_id;table:additions_warehousing" comment:"关联的入库单号"`
	Name           string `form:"name"  search:"type:contains;column:name;table:additions_warehousing" comment:"资产名称"`
	Spec           string `form:"spec"  search:"type:contains;column:spec;table:additions_warehousing" comment:"规格型号"`
	Brand          string `form:"brand"  search:"type:exact;column:brand;table:additions_warehousing" comment:"品牌名称"`
	Sn             string `form:"sn"  search:"type:exact;column:sn;table:additions_warehousing" comment:"资产SN"`
	UserId         string `form:"userId"  search:"type:exact;column:user_id;table:additions_warehousing" comment:"采购人员ID"`
	AdditionsWarehousingOrder
}

type AdditionsWarehousingOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:additions_warehousing"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:additions_warehousing"`
	PurchaseAt string `form:"purchaseAtOrder"  search:"type:order;column:purchase_at;table:additions_warehousing"`
	ExpireAt   string `form:"expireAtOrder"  search:"type:order;column:expire_at;table:additions_warehousing"`
	CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:additions_warehousing"`
	SupplierId string `form:"supplierIdOrder"  search:"type:order;column:supplier_id;table:additions_warehousing"`
	WId        string `form:"wIdOrder"  search:"type:order;column:w_id;table:additions_warehousing"`
	Name       string `form:"nameOrder"  search:"type:order;column:name;table:additions_warehousing"`
	Spec       string `form:"specOrder"  search:"type:order;column:spec;table:additions_warehousing"`
	Brand      string `form:"brandOrder"  search:"type:order;column:brand;table:additions_warehousing"`
	Sn         string `form:"snOrder"  search:"type:order;column:sn;table:additions_warehousing"`
	Unit       string `form:"unitOrder"  search:"type:order;column:unit;table:additions_warehousing"`
	Price      string `form:"priceOrder"  search:"type:order;column:price;table:additions_warehousing"`
	UserId     string `form:"userIdOrder"  search:"type:order;column:user_id;table:additions_warehousing"`
	Desc       string `form:"descOrder"  search:"type:order;column:desc;table:additions_warehousing"`
}

func (m *AdditionsWarehousingGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetInsertReq struct {
	StoreRoomId int                             `json:"storeRoomId" comment:"存放位置"`
	Desc        string                          `json:"desc" comment:"备注"`
	Asset       []AdditionsWarehousingInsertReq `json:"asset" comment:"资产列表"`
}
type AdditionsWarehousingInsertReq struct {
	PurchaseAt string  `json:"purchaseAt" comment:"采购日期"`
	ExpireAt   string  `json:"expireAt" comment:"维保到期日"`
	CategoryId int64   `json:"categoryId" comment:"关联的资产分类ID"`
	SupplierId int64   `json:"supplierId" comment:"供应商ID"`
	WId        int64   `json:"wId" comment:"关联的入库单号"`
	Name       string  `json:"name" comment:"资产名称"`
	Spec       string  `json:"spec" comment:"规格型号"`
	Brand      string  `json:"brand" comment:"品牌名称"`
	Sn         string  `json:"sn" comment:"资产SN"`
	UnitId     int64   `json:"unitId" comment:"单位"`
	Price      float64 `json:"price" comment:"价格"`
	UserId     int64   `json:"userId" comment:"采购人员ID"`
	Desc       string  `json:"desc" comment:"备注"`
	common.ControlBy
}

type AssetUpdateReq struct {
	Id          int                             `uri:"id" comment:"主键编码"` // 主键编码
	StoreRoomId int                             `json:"storeRoomId" comment:"存放位置"`
	Desc        string                          `json:"desc" comment:"备注"`
	Asset       []AdditionsWarehousingUpdateReq `json:"asset" comment:"资产列表"`
}

func (s *AdditionsWarehousingInsertReq) Generate(model *models.AdditionsWarehousing) {

	model.CategoryId = s.CategoryId
	model.SupplierId = s.SupplierId
	model.WId = s.WId
	model.Name = s.Name
	model.Spec = s.Spec
	model.Brand = s.Brand
	model.Sn = s.Sn
	model.UnitId = s.UnitId
	model.Price = s.Price
	model.UserId = s.UserId
	model.Desc = s.Desc
}

type AdditionsWarehousingUpdateReq struct {
	Id         int     `json:"id" comment:"主键编码"` // 主键编码
	PurchaseAt string  `json:"purchaseAt" comment:"采购日期"`
	ExpireAt   string  `json:"expireAt" comment:"维保到期日"`
	CategoryId int64   `json:"categoryId" comment:"关联的资产分类ID"`
	SupplierId int64   `json:"supplierId" comment:"供应商ID"`
	WId        int64   `json:"wId" comment:"关联的入库单号"`
	Name       string  `json:"name" comment:"资产名称"`
	Spec       string  `json:"spec" comment:"规格型号"`
	Brand      string  `json:"brand" comment:"品牌名称"`
	Sn         string  `json:"sn" comment:"资产SN"`
	UnitId     int64   `json:"unitId" comment:"单位"`
	Price      float64 `json:"price" comment:"价格"`
	UserId     int64   `json:"userId" comment:"采购人员ID"`
	Desc       string  `json:"desc" comment:"备注"`
	common.ControlBy
}

func (s *AdditionsWarehousingUpdateReq) Generate(model *models.AdditionsWarehousing) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}

	model.CategoryId = s.CategoryId
	model.SupplierId = s.SupplierId
	model.WId = s.WId
	model.Name = s.Name
	model.Spec = s.Spec
	model.Brand = s.Brand
	model.Sn = s.Sn
	model.UnitId = s.UnitId
	model.Price = s.Price
	model.UserId = s.UserId
	model.Desc = s.Desc
}

func (s *AdditionsWarehousingUpdateReq) GetId() interface{} {
	return s.Id
}

type AdditionsOrderGetReq struct {
	dto.Pagination `search:"-"`
}

// AdditionsWarehousingGetReq 功能获取请求参数
type AdditionsWarehousingGetReq struct {
	Id int `uri:"id"`
}

func (s *AdditionsWarehousingGetReq) GetId() interface{} {
	return s.Id
}

// AdditionsWarehousingDeleteReq 功能删除请求参数
type AdditionsWarehousingDeleteReq struct {
	Ids      []int `json:"ids"`
	Unscoped int   `json:"unscoped"`
}

func (s *AdditionsWarehousingDeleteReq) GetId() interface{} {
	return s.Ids
}
