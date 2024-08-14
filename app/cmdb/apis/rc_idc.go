package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
)

type RcIdc struct {
	api.Api
}

// GetPage 获取RcIdc列表
// @Summary 获取RcIdc列表
// @Description 获取RcIdc列表
// @Tags RcIdc
// @Param enable query string false "开关"
// @Param name query string false "机房名称"
// @Param status query string false "机房状态"
// @Param belong query string false "机房归属"
// @Param typeId query string false "机房类型"
// @Param businessUser query string false "商务人员"
// @Param region query string false "所在区域"
// @Param transProvince query string false "是否跨省"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RcIdc}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-idc [get]
// @Security Bearer
func (e RcIdc) GetPage(c *gin.Context) {
	req := dto.RcIdcGetPageReq{}
	s := service.RcIdc{}
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
	list := make([]models.RcIdc, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RcIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RcIdc
// @Summary 获取RcIdc
// @Description 获取RcIdc
// @Tags RcIdc
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RcIdc} "{"code": 200, "data": [...]}"
// @Router /api/v1/rc-idc/{id} [get]
// @Security Bearer
func (e RcIdc) Get(c *gin.Context) {
	req := dto.RcIdcGetReq{}
	s := service.RcIdc{}
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
	var object models.RcIdc

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RcIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RcIdc
// @Summary 创建RcIdc
// @Description 创建RcIdc
// @Tags RcIdc
// @Accept application/json
// @Product application/json
// @Param data body dto.RcIdcInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rc-idc [post]
// @Security Bearer
func (e RcIdc) Insert(c *gin.Context) {
	req := dto.RcIdcInsertReq{}
	s := service.RcIdc{}
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
		e.Error(500, err, fmt.Sprintf("创建RcIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RcIdc
// @Summary 修改RcIdc
// @Description 修改RcIdc
// @Tags RcIdc
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RcIdcUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rc-idc/{id} [put]
// @Security Bearer
func (e RcIdc) Update(c *gin.Context) {
	req := dto.RcIdcUpdateReq{}
	s := service.RcIdc{}
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
		e.Error(500, err, fmt.Sprintf("修改RcIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RcIdc
// @Summary 删除RcIdc
// @Description 删除RcIdc
// @Tags RcIdc
// @Param data body dto.RcIdcDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rc-idc [delete]
// @Security Bearer
func (e RcIdc) Delete(c *gin.Context) {
	s := service.RcIdc{}
	req := dto.RcIdcDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RcIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
