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

type AssetReturn struct {
	api.Api
}

// GetPage 获取资产退还记录列表
// @Summary 获取资产退还记录列表
// @Description 获取资产退还记录列表
// @Tags 资产退还记录
// @Param assetId query int64 false "资产编码"
// @Param returnPerson query int64 false "退还人编码"
// @Param reason query string false "退还原因"
// @Param returnAt query time.Time false "退还时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetReturn}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-return [get]
// @Security Bearer
func (e AssetReturn) GetPage(c *gin.Context) {
	req := dto.AssetReturnGetPageReq{}
	s := service.AssetReturn{}
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
	list := make([]models.AssetReturn, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产退还记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产退还记录
// @Summary 获取资产退还记录
// @Description 获取资产退还记录
// @Tags 资产退还记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetReturn} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-return/{id} [get]
// @Security Bearer
func (e AssetReturn) Get(c *gin.Context) {
	req := dto.AssetReturnGetReq{}
	s := service.AssetReturn{}
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
	var object models.AssetReturn

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产退还记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产退还记录
// @Summary 创建资产退还记录
// @Description 创建资产退还记录
// @Tags 资产退还记录
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetReturnInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-return [post]
// @Security Bearer
func (e AssetReturn) Insert(c *gin.Context) {
	req := dto.AssetReturnInsertReq{}
	s := service.AssetReturn{}
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
		e.Error(500, err, fmt.Sprintf("创建资产退还记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产退还记录
// @Summary 修改资产退还记录
// @Description 修改资产退还记录
// @Tags 资产退还记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetReturnUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-return/{id} [put]
// @Security Bearer
func (e AssetReturn) Update(c *gin.Context) {
	req := dto.AssetReturnUpdateReq{}
	s := service.AssetReturn{}
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
		e.Error(500, err, fmt.Sprintf("修改资产退还记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产退还记录
// @Summary 删除资产退还记录
// @Description 删除资产退还记录
// @Tags 资产退还记录
// @Param data body dto.AssetReturnDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-return [delete]
// @Security Bearer
func (e AssetReturn) Delete(c *gin.Context) {
	s := service.AssetReturn{}
	req := dto.AssetReturnDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产退还记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
