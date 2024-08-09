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

type AssetGroupMember struct {
	api.Api
}

// GetPage 获取资产组合成员列表
// @Summary 获取资产组合成员列表
// @Description 获取资产组合成员列表
// @Tags 资产组合成员
// @Param assetGroupId query string false "资产组合编码"
// @Param assetId query string false "资产编码"
// @Param isMain query string false "是否为主资产"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetGroupMember}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-group-member [get]
// @Security Bearer
func (e AssetGroupMember) GetPage(c *gin.Context) {
	req := dto.AssetGroupMemberGetPageReq{}
	s := service.AssetGroupMember{}
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
	list := make([]models.AssetGroupMember, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产组合成员失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产组合成员
// @Summary 获取资产组合成员
// @Description 获取资产组合成员
// @Tags 资产组合成员
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetGroupMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-group-member/{id} [get]
// @Security Bearer
func (e AssetGroupMember) Get(c *gin.Context) {
	req := dto.AssetGroupMemberGetReq{}
	s := service.AssetGroupMember{}
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
	var object models.AssetGroupMember

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产组合成员失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产组合成员
// @Summary 创建资产组合成员
// @Description 创建资产组合成员
// @Tags 资产组合成员
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetGroupMemberInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-group-member [post]
// @Security Bearer
func (e AssetGroupMember) Insert(c *gin.Context) {
	req := dto.AssetGroupMemberInsertReq{}
	s := service.AssetGroupMember{}
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
		e.Error(500, err, fmt.Sprintf("创建资产组合成员失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产组合成员
// @Summary 修改资产组合成员
// @Description 修改资产组合成员
// @Tags 资产组合成员
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetGroupMemberUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-group-member/{id} [put]
// @Security Bearer
func (e AssetGroupMember) Update(c *gin.Context) {
	req := dto.AssetGroupMemberUpdateReq{}
	s := service.AssetGroupMember{}
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
		e.Error(500, err, fmt.Sprintf("修改资产组合成员失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产组合成员
// @Summary 删除资产组合成员
// @Description 删除资产组合成员
// @Tags 资产组合成员
// @Param data body dto.AssetGroupMemberDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-group-member [delete]
// @Security Bearer
func (e AssetGroupMember) Delete(c *gin.Context) {
	s := service.AssetGroupMember{}
	req := dto.AssetGroupMemberDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产组合成员失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
