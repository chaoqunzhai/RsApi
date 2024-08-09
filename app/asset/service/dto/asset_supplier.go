package dto

import (
	"go-admin/app/asset/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AssetSupplierGetPageReq struct {
	dto.Pagination `search:"-"`
	SupplierName   string `form:"supplierName"  search:"type:exact;column:supplier_name;table:asset_supplier" comment:"供应商名称"`
	ContactPerson  string `form:"contactPerson"  search:"type:exact;column:contact_person;table:asset_supplier" comment:"联系人"`
	PhoneNumber    string `form:"phoneNumber"  search:"type:exact;column:phone_number;table:asset_supplier" comment:"联系电话"`
	AssetSupplierOrder
}

type AssetSupplierOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:asset_supplier"`
	SupplierName  string `form:"supplierNameOrder"  search:"type:order;column:supplier_name;table:asset_supplier"`
	ContactPerson string `form:"contactPersonOrder"  search:"type:order;column:contact_person;table:asset_supplier"`
	PhoneNumber   string `form:"phoneNumberOrder"  search:"type:order;column:phone_number;table:asset_supplier"`
	Email         string `form:"emailOrder"  search:"type:order;column:email;table:asset_supplier"`
	Address       string `form:"addressOrder"  search:"type:order;column:address;table:asset_supplier"`
	Remark        string `form:"remarkOrder"  search:"type:order;column:remark;table:asset_supplier"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:asset_supplier"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:asset_supplier"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:asset_supplier"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:asset_supplier"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:asset_supplier"`
}

func (m *AssetSupplierGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AssetSupplierInsertReq struct {
	Id            int    `json:"-" comment:"主键"` // 主键
	SupplierName  string `json:"supplierName" comment:"供应商名称"`
	ContactPerson string `json:"contactPerson" comment:"联系人"`
	PhoneNumber   string `json:"phoneNumber" comment:"联系电话"`
	Email         string `json:"email" comment:"邮箱"`
	Address       string `json:"address" comment:"地址"`
	Remark        string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetSupplierInsertReq) Generate(model *models.AssetSupplier) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SupplierName = s.SupplierName
	model.ContactPerson = s.ContactPerson
	model.PhoneNumber = s.PhoneNumber
	model.Email = s.Email
	model.Address = s.Address
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AssetSupplierInsertReq) GetId() interface{} {
	return s.Id
}

type AssetSupplierUpdateReq struct {
	Id            int    `uri:"id" comment:"主键"` // 主键
	SupplierName  string `json:"supplierName" comment:"供应商名称"`
	ContactPerson string `json:"contactPerson" comment:"联系人"`
	PhoneNumber   string `json:"phoneNumber" comment:"联系电话"`
	Email         string `json:"email" comment:"邮箱"`
	Address       string `json:"address" comment:"地址"`
	Remark        string `json:"remark" comment:"备注"`
	common.ControlBy
}

func (s *AssetSupplierUpdateReq) Generate(model *models.AssetSupplier) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SupplierName = s.SupplierName
	model.ContactPerson = s.ContactPerson
	model.PhoneNumber = s.PhoneNumber
	model.Email = s.Email
	model.Address = s.Address
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AssetSupplierUpdateReq) GetId() interface{} {
	return s.Id
}

// AssetSupplierGetReq 功能获取请求参数
type AssetSupplierGetReq struct {
	Id int `uri:"id"`
}

func (s *AssetSupplierGetReq) GetId() interface{} {
	return s.Id
}

// AssetSupplierDeleteReq 功能删除请求参数
type AssetSupplierDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AssetSupplierDeleteReq) GetId() interface{} {
	return s.Ids
}
