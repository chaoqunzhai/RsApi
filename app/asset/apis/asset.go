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

type Asset struct {
	api.Api
}

// GetPage 获取资产详情列表
// @Summary 获取资产详情列表
// @Description 获取资产详情列表
// @Tags 资产详情
// @Param assetCode query string false "资产编号"
// @Param snCode query string false "SN编码"
// @Param categoryId query string false "资产类别"
// @Param specification query string false "规格型号"
// @Param brand query string false "品牌"
// @Param unit query string false "计量单位"
// @Param unitPrice query string false "单价"
// @Param status query string false "状态(0=在库, 1=出库, 2=在用, 3=处置)"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Asset}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset [get]
// @Security Bearer
func (e Asset) GetPage(c *gin.Context) {
	req := dto.AssetGetPageReq{}
	s := service.Asset{}
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
	list := make([]models.Asset, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产详情失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产详情
// @Summary 获取资产详情
// @Description 获取资产详情
// @Tags 资产详情
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Asset} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset/{id} [get]
// @Security Bearer
func (e Asset) Get(c *gin.Context) {
	req := dto.AssetGetReq{}
	s := service.Asset{}
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
	var object models.Asset

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产详情失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产详情
// @Summary 创建资产详情
// @Description 创建资产详情
// @Tags 资产详情
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset [post]
// @Security Bearer
func (e Asset) Insert(c *gin.Context) {
	req := dto.AssetInsertReq{}
	s := service.Asset{}
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
		e.Error(500, err, fmt.Sprintf("创建资产详情失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产详情
// @Summary 修改资产详情
// @Description 修改资产详情
// @Tags 资产详情
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset/{id} [put]
// @Security Bearer
func (e Asset) Update(c *gin.Context) {
	req := dto.AssetUpdateReq{}
	s := service.Asset{}
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
		e.Error(500, err, fmt.Sprintf("修改资产详情失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产详情
// @Summary 删除资产详情
// @Description 删除资产详情
// @Tags 资产详情
// @Param data body dto.AssetDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset [delete]
// @Security Bearer
func (e Asset) Delete(c *gin.Context) {
	s := service.Asset{}
	req := dto.AssetDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产详情失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
