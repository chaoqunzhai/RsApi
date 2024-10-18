package dto

import (
	"fmt"
	"go-admin/app/cmdb/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strings"
)

type RsBusinessGetPageReq struct {
	dto.Pagination `search:"-"`
	TreeTag        int    `form:"treeTag" search:"-"`
	Enable         string `form:"enable"  search:"type:exact;column:enable;table:rs_business" comment:"开关"`
	Status         string `form:"status"  search:"type:exact;column:status;table:rs_business" comment:""`
	Name           string `form:"name"  search:"type:contains;column:name;table:rs_business" comment:"业务云名称"`
	RsBusinessOrder
}

type RsBusinessOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:rs_business"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:rs_business"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:rs_business"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:rs_business"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:rs_business"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:rs_business"`
	Enable    string `form:"enableOrder"  search:"type:order;column:enable;table:rs_business"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:rs_business"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:rs_business"`
	Algorithm string `form:"algorithmOrder"  search:"type:order;column:algorithm;table:rs_business"`
}

func (m *RsBusinessGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RsBusinessInsertReq struct {
	Id            int                  `json:"-" comment:"主键编码"` // 主键编码
	Status        int                  `json:"status" comment:"状态"`
	Desc          string               `json:"desc" comment:"描述信息"`
	Name          string               `json:"name" comment:"业务云名称"`
	EnName        string               `json:"enName" gorm:"index;type:varchar(30);comment:业务别名"`
	BillingMethod int                  `json:"billingMethod" comment:"计费方式"`
	ParentId      int                  `json:"parentId" gorm:"comment:父业务"`
	StartUsage    string               `json:"startUsage" gorm:"index;type:varchar(30);comment:利用率开始时间"`
	EndUsage      string               `json:"endUsage" gorm:"index;type:varchar(30);comment:利用率结束时间"`
	CostCnf       []RsCostCnfInsertReq `json:"costCnf"`
	common.ControlBy
}

type RsCostCnfInsertReq struct {
	BuId           int     `json:"buId"`
	Id             int     `json:"id" comment:"主键编码"` // 主键编码
	Isp            int     `json:"isp"`
	Minimum        float64 `json:"minimum"`
	IpType         int     `json:"ipType"`
	DialType       int     `json:"dialType"`
	BandwidthLower float64 `json:"bandwidthLower"`
	BandwidthLimit float64 `json:"bandwidthLimit"`
	Price          float64 `json:"price"`
	//Start          string  `json:"start"`
	//End            string  `json:"end"`
	RangePrice float64 `json:"rangePrice"`
}

func (s *RsCostCnfInsertReq) Generate(model *models.RsBusinessCostCnf) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}

	model.Isp = s.Isp
	model.BuId = s.BuId
	model.DialType = s.DialType
	model.IpType = s.IpType
	//if s.Start != "" {
	//	if star, err := time.ParseInLocation(time.DateOnly, s.Start, global.LOC); err == nil {
	//
	//		model.Start = common.XTime{
	//			Time: star,
	//		}
	//	}
	//
	//}
	//
	//if s.End != "" {
	//	if end, err := time.ParseInLocation(time.DateOnly, s.End, global.LOC); err == nil {
	//		model.End = common.XTime{
	//			Time: end,
	//		}
	//	}
	//
	//}
	//model.RangePrice = s.RangePrice
	model.Price = s.Price
}
func (s *RsCostCnfInsertReq) GetId() interface{} {
	return s.Id
}

func (s *RsBusinessInsertReq) Generate(model *models.RsBusiness) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.EnName = s.EnName
	model.Desc = s.Desc
	model.Name = s.Name
	model.Status = s.Status
	model.BillingMethod = s.BillingMethod
	model.ParentId = s.ParentId
	model.StartUsage = s.StartUsage
	model.EndUsage = s.EndUsage
}

func (s *RsBusinessInsertReq) GetId() interface{} {
	return s.Id
}

type RsBusinessUpdateReq struct {
	Id            int                  `uri:"id" comment:"主键编码"` // 主键编码
	Status        int                  `json:"status" comment:"状态"`
	Desc          string               `json:"desc" comment:"描述信息"`
	Name          string               `json:"name" comment:"业务云名称"`
	EnName        string               `json:"enName" gorm:"index;type:varchar(30);comment:业务别名"`
	BillingMethod int                  `json:"billingMethod" comment:"计费方式"`
	ParentId      int                  `json:"parentId" gorm:"comment:父业务"`
	StartUsage    string               `json:"startUsage" gorm:"index;type:varchar(30);comment:利用率开始时间"`
	EndUsage      string               `json:"endUsage" gorm:"index;type:varchar(30);comment:利用率结束时间"`
	CostCnf       []RsCostCnfInsertReq `json:"costCnf"`
	common.ControlBy
}

func (s *RsBusinessUpdateReq) Generate(model *models.RsBusiness) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Desc = s.Desc
	model.EnName = s.EnName
	model.Name = s.Name
	model.Status = s.Status
	model.BillingMethod = s.BillingMethod
	model.ParentId = s.ParentId
	model.StartUsage = s.StartUsage
	model.EndUsage = s.EndUsage
}

func (s *RsBusinessUpdateReq) GetId() interface{} {
	return s.Id
}

// RsBusinessGetReq 功能获取请求参数
type RsBusinessGetReq struct {
	Id int `uri:"id"`
}

func (s *RsBusinessGetReq) GetId() interface{} {
	return s.Id
}

// RsBusinessDeleteReq 功能删除请求参数
type RsBusinessDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RsBusinessDeleteReq) GetId() interface{} {
	return s.Ids
}

func (s *RsBusinessDeleteReq) GetIdStr() string {
	var cache []string
	for _, i := range s.Ids {
		if i > 0 {
			cache = append(cache, fmt.Sprintf("%d", i))
		}
	}
	return strings.Join(cache, ",")
}
