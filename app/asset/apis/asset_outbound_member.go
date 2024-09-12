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

type AssetOutboundMember struct {
	api.Api
}

// GetPage 获取资产出库成员表列表
// @Summary 获取资产出库成员表列表
// @Description 获取资产出库成员表列表
// @Tags 资产出库成员表
// @Param assetOutboundId query string false "资产出库编码"
// @Param assetOutboundCode query string false "资产出库单号"
// @Param assetId query string false "资产编码"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetOutboundMember}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-outbound-member [get]
// @Security Bearer
func (e AssetOutboundMember) GetPage(c *gin.Context) {
	req := dto.AssetOutboundMemberGetPageReq{}
	s := service.AssetOutboundMember{}
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
	list := make([]models.AssetOutboundMember, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产出库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产出库成员表
// @Summary 获取资产出库成员表
// @Description 获取资产出库成员表
// @Tags 资产出库成员表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetOutboundMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-outbound-member/{id} [get]
// @Security Bearer
func (e AssetOutboundMember) Get(c *gin.Context) {
	req := dto.AssetOutboundMemberGetReq{}
	s := service.AssetOutboundMember{}
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
	var object models.AssetOutboundMember

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产出库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产出库成员表
// @Summary 创建资产出库成员表
// @Description 创建资产出库成员表
// @Tags 资产出库成员表
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetOutboundMemberInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-outbound-member [post]
// @Security Bearer
func (e AssetOutboundMember) Insert(c *gin.Context) {
	req := dto.AssetOutboundMemberInsertReq{}
	s := service.AssetOutboundMember{}
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
		e.Error(500, err, fmt.Sprintf("创建资产出库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产出库成员表
// @Summary 修改资产出库成员表
// @Description 修改资产出库成员表
// @Tags 资产出库成员表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetOutboundMemberUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-outbound-member/{id} [put]
// @Security Bearer
func (e AssetOutboundMember) Update(c *gin.Context) {
	req := dto.AssetOutboundMemberUpdateReq{}
	s := service.AssetOutboundMember{}
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
		e.Error(500, err, fmt.Sprintf("修改资产出库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产出库成员表
// @Summary 删除资产出库成员表
// @Description 删除资产出库成员表
// @Tags 资产出库成员表
// @Param data body dto.AssetOutboundMemberDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-outbound-member [delete]
// @Security Bearer
func (e AssetOutboundMember) Delete(c *gin.Context) {
	s := service.AssetOutboundMember{}
	req := dto.AssetOutboundMemberDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产出库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
