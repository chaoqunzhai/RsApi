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

type AssetStock struct {
	api.Api
}

// GetPage 获取资产库存列表
// @Summary 获取资产库存列表
// @Description 获取资产库存列表
// @Tags 资产库存
// @Param warehouseId query string false "库房编码"
// @Param categoryId query string false "资产类别编码"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetStock}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-stock [get]
// @Security Bearer
func (e AssetStock) GetPage(c *gin.Context) {
	req := dto.AssetStockGetPageReq{}
	s := service.AssetStock{}
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
	list := make([]models.AssetStock, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产库存失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产库存
// @Summary 获取资产库存
// @Description 获取资产库存
// @Tags 资产库存
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetStock} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-stock/{id} [get]
// @Security Bearer
func (e AssetStock) Get(c *gin.Context) {
	req := dto.AssetStockGetReq{}
	s := service.AssetStock{}
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
	var object models.AssetStock

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产库存失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产库存
// @Summary 创建资产库存
// @Description 创建资产库存
// @Tags 资产库存
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetStockInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-stock [post]
// @Security Bearer
func (e AssetStock) Insert(c *gin.Context) {
	req := dto.AssetStockInsertReq{}
	s := service.AssetStock{}
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
		e.Error(500, err, fmt.Sprintf("创建资产库存失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产库存
// @Summary 修改资产库存
// @Description 修改资产库存
// @Tags 资产库存
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetStockUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-stock/{id} [put]
// @Security Bearer
func (e AssetStock) Update(c *gin.Context) {
	req := dto.AssetStockUpdateReq{}
	s := service.AssetStock{}
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
		e.Error(500, err, fmt.Sprintf("修改资产库存失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产库存
// @Summary 删除资产库存
// @Description 删除资产库存
// @Tags 资产库存
// @Param data body dto.AssetStockDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-stock [delete]
// @Security Bearer
func (e AssetStock) Delete(c *gin.Context) {
	s := service.AssetStock{}
	req := dto.AssetStockDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产库存失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
