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

type AssetGroup struct {
	api.Api
}

// GetPage 获取资产组合列表
// @Summary 获取资产组合列表
// @Description 获取资产组合列表
// @Tags 资产组合
// @Param groupName query string false "资产组合名称"
// @Param mainAssetId query string false "主资产编码"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetGroup}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-group [get]
// @Security Bearer
func (e AssetGroup) GetPage(c *gin.Context) {
	req := dto.AssetGroupGetPageReq{}
	s := service.AssetGroup{}
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
	list := make([]models.AssetGroup, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产组合失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产组合
// @Summary 获取资产组合
// @Description 获取资产组合
// @Tags 资产组合
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetGroup} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-group/{id} [get]
// @Security Bearer
func (e AssetGroup) Get(c *gin.Context) {
	req := dto.AssetGroupGetReq{}
	s := service.AssetGroup{}
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
	var object models.AssetGroup

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产组合失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产组合
// @Summary 创建资产组合
// @Description 创建资产组合
// @Tags 资产组合
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetGroupInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-group [post]
// @Security Bearer
func (e AssetGroup) Insert(c *gin.Context) {
	req := dto.AssetGroupInsertReq{}
	s := service.AssetGroup{}
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
		e.Error(500, err, fmt.Sprintf("创建资产组合失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产组合
// @Summary 修改资产组合
// @Description 修改资产组合
// @Tags 资产组合
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetGroupUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-group/{id} [put]
// @Security Bearer
func (e AssetGroup) Update(c *gin.Context) {
	req := dto.AssetGroupUpdateReq{}
	s := service.AssetGroup{}
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
		e.Error(500, err, fmt.Sprintf("修改资产组合失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产组合
// @Summary 删除资产组合
// @Description 删除资产组合
// @Tags 资产组合
// @Param data body dto.AssetGroupDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-group [delete]
// @Security Bearer
func (e AssetGroup) Delete(c *gin.Context) {
	s := service.AssetGroup{}
	req := dto.AssetGroupDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产组合失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
