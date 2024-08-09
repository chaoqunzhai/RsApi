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

type AssetOutbound struct {
	api.Api
}

// GetPage 获取资产出库记录列表
// @Summary 获取资产出库记录列表
// @Description 获取资产出库记录列表
// @Tags 资产出库记录
// @Param assetId query string false "资产编码"
// @Param warehouseId query string false "库房编码"
// @Param outboundTo query string false "出库去向(客户编码)"
// @Param outboundBy query string false "出库人编码"
// @Param outboundAt query time.Time false "出库时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetOutbound}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-outbound [get]
// @Security Bearer
func (e AssetOutbound) GetPage(c *gin.Context) {
	req := dto.AssetOutboundGetPageReq{}
	s := service.AssetOutbound{}
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
	list := make([]models.AssetOutbound, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产出库记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产出库记录
// @Summary 获取资产出库记录
// @Description 获取资产出库记录
// @Tags 资产出库记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetOutbound} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-outbound/{id} [get]
// @Security Bearer
func (e AssetOutbound) Get(c *gin.Context) {
	req := dto.AssetOutboundGetReq{}
	s := service.AssetOutbound{}
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
	var object models.AssetOutbound

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产出库记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产出库记录
// @Summary 创建资产出库记录
// @Description 创建资产出库记录
// @Tags 资产出库记录
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetOutboundInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-outbound [post]
// @Security Bearer
func (e AssetOutbound) Insert(c *gin.Context) {
	req := dto.AssetOutboundInsertReq{}
	s := service.AssetOutbound{}
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
		e.Error(500, err, fmt.Sprintf("创建资产出库记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产出库记录
// @Summary 修改资产出库记录
// @Description 修改资产出库记录
// @Tags 资产出库记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetOutboundUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-outbound/{id} [put]
// @Security Bearer
func (e AssetOutbound) Update(c *gin.Context) {
	req := dto.AssetOutboundUpdateReq{}
	s := service.AssetOutbound{}
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
		e.Error(500, err, fmt.Sprintf("修改资产出库记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产出库记录
// @Summary 删除资产出库记录
// @Description 删除资产出库记录
// @Tags 资产出库记录
// @Param data body dto.AssetOutboundDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-outbound [delete]
// @Security Bearer
func (e AssetOutbound) Delete(c *gin.Context) {
	s := service.AssetOutbound{}
	req := dto.AssetOutboundDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产出库记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
