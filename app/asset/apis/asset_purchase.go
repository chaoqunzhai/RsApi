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

type AssetPurchase struct {
	api.Api
}

// GetPage 获取资产采购记录列表
// @Summary 获取资产采购记录列表
// @Description 获取资产采购记录列表
// @Tags 资产采购记录
// @Param purchaseCode query string false "采购单编号"
// @Param categoryId query string false "资产类型编码"
// @Param supplierId query string false "供应商编码"
// @Param purchaseUser query string false "采购人编码"
// @Param specification query string false "规格型号"
// @Param brand query string false "品牌"
// @Param purchaseAt query string false "采购日期"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetPurchase}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-purchase [get]
// @Security Bearer
func (e AssetPurchase) GetPage(c *gin.Context) {
	req := dto.AssetPurchaseGetPageReq{}
	s := service.AssetPurchase{}
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
	list := make([]models.AssetPurchase, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产采购记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产采购记录
// @Summary 获取资产采购记录
// @Description 获取资产采购记录
// @Tags 资产采购记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetPurchase} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-purchase/{id} [get]
// @Security Bearer
func (e AssetPurchase) Get(c *gin.Context) {
	req := dto.AssetPurchaseGetReq{}
	s := service.AssetPurchase{}
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
	var object models.AssetPurchase

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产采购记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产采购记录
// @Summary 创建资产采购记录
// @Description 创建资产采购记录
// @Tags 资产采购记录
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetPurchaseInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-purchase [post]
// @Security Bearer
func (e AssetPurchase) Insert(c *gin.Context) {
	req := dto.AssetPurchaseInsertReq{}
	s := service.AssetPurchase{}
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
		e.Error(500, err, fmt.Sprintf("创建资产采购记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产采购记录
// @Summary 修改资产采购记录
// @Description 修改资产采购记录
// @Tags 资产采购记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetPurchaseUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-purchase/{id} [put]
// @Security Bearer
func (e AssetPurchase) Update(c *gin.Context) {
	req := dto.AssetPurchaseUpdateReq{}
	s := service.AssetPurchase{}
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
		e.Error(500, err, fmt.Sprintf("修改资产采购记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产采购记录
// @Summary 删除资产采购记录
// @Description 删除资产采购记录
// @Tags 资产采购记录
// @Param data body dto.AssetPurchaseDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-purchase [delete]
// @Security Bearer
func (e AssetPurchase) Delete(c *gin.Context) {
	s := service.AssetPurchase{}
	req := dto.AssetPurchaseDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产采购记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
