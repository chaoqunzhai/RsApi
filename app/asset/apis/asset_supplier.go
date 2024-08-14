package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/asset/models"
	"go-admin/app/asset/service"
	"go-admin/app/asset/service/dto"
	"go-admin/common/actions"
)

type AssetSupplier struct {
	api.Api
}

// GetPage 获取资产供应商列表
// @Summary 获取资产供应商列表
// @Description 获取资产供应商列表
// @Tags 资产供应商
// @Param supplierName query string false "供应商名称"
// @Param contactPerson query string false "联系人"
// @Param phoneNumber query string false "联系电话"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetSupplier}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-supplier [get]
// @Security Bearer
func (e AssetSupplier) GetPage(c *gin.Context) {
	req := dto.AssetSupplierGetPageReq{}
	s := service.AssetSupplier{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.AssetSupplier, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产供应商失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产供应商
// @Summary 获取资产供应商
// @Description 获取资产供应商
// @Tags 资产供应商
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetSupplier} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-supplier/{id} [get]
// @Security Bearer
func (e AssetSupplier) Get(c *gin.Context) {
	req := dto.AssetSupplierGetReq{}
	s := service.AssetSupplier{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.AssetSupplier

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产供应商失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产供应商
// @Summary 创建资产供应商
// @Description 创建资产供应商
// @Tags 资产供应商
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetSupplierInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-supplier [post]
// @Security Bearer
func (e AssetSupplier) Insert(c *gin.Context) {
	req := dto.AssetSupplierInsertReq{}
	s := service.AssetSupplier{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建资产供应商失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产供应商
// @Summary 修改资产供应商
// @Description 修改资产供应商
// @Tags 资产供应商
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetSupplierUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-supplier/{id} [put]
// @Security Bearer
func (e AssetSupplier) Update(c *gin.Context) {
	req := dto.AssetSupplierUpdateReq{}
	s := service.AssetSupplier{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改资产供应商失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产供应商
// @Summary 删除资产供应商
// @Description 删除资产供应商
// @Tags 资产供应商
// @Param data body dto.AssetSupplierDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-supplier [delete]
// @Security Bearer
func (e AssetSupplier) Delete(c *gin.Context) {
	s := service.AssetSupplier{}
	req := dto.AssetSupplierDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除资产供应商失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
