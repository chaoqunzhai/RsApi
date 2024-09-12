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

type AssetInboundMember struct {
	api.Api
}

// GetPage 获取资产入库成员表列表
// @Summary 获取资产入库成员表列表
// @Description 获取资产入库成员表列表
// @Tags 资产入库成员表
// @Param assetInboundId query int64 false "资产入库编码"
// @Param assetInboundCode query string false "资产入库单号"
// @Param assetId query int64 false "资产编码"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetInboundMember}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-inbound-member [get]
// @Security Bearer
func (e AssetInboundMember) GetPage(c *gin.Context) {
	req := dto.AssetInboundMemberGetPageReq{}
	s := service.AssetInboundMember{}
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
	list := make([]models.AssetInboundMember, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产入库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产入库成员表
// @Summary 获取资产入库成员表
// @Description 获取资产入库成员表
// @Tags 资产入库成员表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetInboundMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-inbound-member/{id} [get]
// @Security Bearer
func (e AssetInboundMember) Get(c *gin.Context) {
	req := dto.AssetInboundMemberGetReq{}
	s := service.AssetInboundMember{}
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
	var object models.AssetInboundMember

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产入库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产入库成员表
// @Summary 创建资产入库成员表
// @Description 创建资产入库成员表
// @Tags 资产入库成员表
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetInboundMemberInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-inbound-member [post]
// @Security Bearer
func (e AssetInboundMember) Insert(c *gin.Context) {
	req := dto.AssetInboundMemberInsertReq{}
	s := service.AssetInboundMember{}
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
		e.Error(500, err, fmt.Sprintf("创建资产入库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产入库成员表
// @Summary 修改资产入库成员表
// @Description 修改资产入库成员表
// @Tags 资产入库成员表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetInboundMemberUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-inbound-member/{id} [put]
// @Security Bearer
func (e AssetInboundMember) Update(c *gin.Context) {
	req := dto.AssetInboundMemberUpdateReq{}
	s := service.AssetInboundMember{}
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
		e.Error(500, err, fmt.Sprintf("修改资产入库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产入库成员表
// @Summary 删除资产入库成员表
// @Description 删除资产入库成员表
// @Tags 资产入库成员表
// @Param data body dto.AssetInboundMemberDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-inbound-member [delete]
// @Security Bearer
func (e AssetInboundMember) Delete(c *gin.Context) {
	s := service.AssetInboundMember{}
	req := dto.AssetInboundMemberDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产入库成员表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
