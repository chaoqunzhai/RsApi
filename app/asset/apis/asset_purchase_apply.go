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

type AssetPurchaseApply struct {
	api.Api
}

// GetPage 获取资产采购申请列表
// @Summary 获取资产采购申请列表
// @Description 获取资产采购申请列表
// @Tags 资产采购申请
// @Param applyCode query string false "申请单编号"
// @Param categoryId query string false "资产类型编码"
// @Param supplierId query string false "供应商编码"
// @Param applyUser query string false "申购人编码"
// @Param specification query string false "规格型号"
// @Param brand query string false "品牌"
// @Param applyAt query string false "申购日期"
// @Param approver query string false "审批人编码"
// @Param approveAt query string false "审批时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetPurchaseApply}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-purchase-apply [get]
// @Security Bearer
func (e AssetPurchaseApply) GetPage(c *gin.Context) {
	req := dto.AssetPurchaseApplyGetPageReq{}
	s := service.AssetPurchaseApply{}
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
	list := make([]models.AssetPurchaseApply, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产采购申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资产采购申请
// @Summary 获取资产采购申请
// @Description 获取资产采购申请
// @Tags 资产采购申请
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetPurchaseApply} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-purchase-apply/{id} [get]
// @Security Bearer
func (e AssetPurchaseApply) Get(c *gin.Context) {
	req := dto.AssetPurchaseApplyGetReq{}
	s := service.AssetPurchaseApply{}
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
	var object models.AssetPurchaseApply

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资产采购申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建资产采购申请
// @Summary 创建资产采购申请
// @Description 创建资产采购申请
// @Tags 资产采购申请
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetPurchaseApplyInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-purchase-apply [post]
// @Security Bearer
func (e AssetPurchaseApply) Insert(c *gin.Context) {
	req := dto.AssetPurchaseApplyInsertReq{}
	s := service.AssetPurchaseApply{}
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
		e.Error(500, err, fmt.Sprintf("创建资产采购申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资产采购申请
// @Summary 修改资产采购申请
// @Description 修改资产采购申请
// @Tags 资产采购申请
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetPurchaseApplyUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-purchase-apply/{id} [put]
// @Security Bearer
func (e AssetPurchaseApply) Update(c *gin.Context) {
	req := dto.AssetPurchaseApplyUpdateReq{}
	s := service.AssetPurchaseApply{}
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
		e.Error(500, err, fmt.Sprintf("修改资产采购申请失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除资产采购申请
// @Summary 删除资产采购申请
// @Description 删除资产采购申请
// @Tags 资产采购申请
// @Param data body dto.AssetPurchaseApplyDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-purchase-apply [delete]
// @Security Bearer
func (e AssetPurchaseApply) Delete(c *gin.Context) {
	s := service.AssetPurchaseApply{}
	req := dto.AssetPurchaseApplyDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资产采购申请失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
