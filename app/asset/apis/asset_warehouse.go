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

type AssetWarehouse struct {
	api.Api
}

// GetPage 获取库房列表
// @Summary 获取库房列表
// @Description 获取库房列表
// @Tags 库房
// @Param warehouseName query string false "库房名称"
// @Param administratorId query string false "管理员编码"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetWarehouse}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-warehouse [get]
// @Security Bearer
func (e AssetWarehouse) GetPage(c *gin.Context) {
	req := dto.AssetWarehouseGetPageReq{}
	s := service.AssetWarehouse{}
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
	list := make([]models.AssetWarehouse, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取库房失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取库房
// @Summary 获取库房
// @Description 获取库房
// @Tags 库房
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetWarehouse} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-warehouse/{id} [get]
// @Security Bearer
func (e AssetWarehouse) Get(c *gin.Context) {
	req := dto.AssetWarehouseGetReq{}
	s := service.AssetWarehouse{}
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
	var object models.AssetWarehouse

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取库房失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建库房
// @Summary 创建库房
// @Description 创建库房
// @Tags 库房
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetWarehouseInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-warehouse [post]
// @Security Bearer
func (e AssetWarehouse) Insert(c *gin.Context) {
	req := dto.AssetWarehouseInsertReq{}
	s := service.AssetWarehouse{}
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
		e.Error(500, err, fmt.Sprintf("创建库房失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改库房
// @Summary 修改库房
// @Description 修改库房
// @Tags 库房
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetWarehouseUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-warehouse/{id} [put]
// @Security Bearer
func (e AssetWarehouse) Update(c *gin.Context) {
	req := dto.AssetWarehouseUpdateReq{}
	s := service.AssetWarehouse{}
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
		e.Error(500, err, fmt.Sprintf("修改库房失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除库房
// @Summary 删除库房
// @Description 删除库房
// @Tags 库房
// @Param data body dto.AssetWarehouseDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-warehouse [delete]
// @Security Bearer
func (e AssetWarehouse) Delete(c *gin.Context) {
	s := service.AssetWarehouse{}
	req := dto.AssetWarehouseDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除库房失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
