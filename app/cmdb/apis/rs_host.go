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

type RsHost struct {
	api.Api
}

// GetPage 获取RsHost列表
// @Summary 获取RsHost列表
// @Description 获取RsHost列表
// @Tags RsHost
// @Param enable query string false "开关"
// @Param hostName query string false "主机名"
// @Param sn query string false "sn"
// @Param ip query string false "ip"
// @Param kernel query string false "内核版本"
// @Param belong query string false "机器归属"
// @Param remark query string false "备注"
// @Param operator query string false "运营商"
// @Param status query string false "主机状态"
// @Param businessSn query string false "业务SN"
// @Param province query string false "省份"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsHost}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host [get]
// @Security Bearer
func (e RsHost) GetPage(c *gin.Context) {
	req := dto.RsHostGetPageReq{}
	s := service.RsHost{}
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
	list := make([]models.RsHost, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsHost
// @Summary 获取RsHost
// @Description 获取RsHost
// @Tags RsHost
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsHost} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host/{id} [get]
// @Security Bearer
func (e RsHost) Get(c *gin.Context) {
	req := dto.RsHostGetReq{}
	s := service.RsHost{}
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
	var object models.RsHost

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsHost
// @Summary 创建RsHost
// @Description 创建RsHost
// @Tags RsHost
// @Accept application/json
// @Product application/json
// @Param data body dto.RsHostInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-host [post]
// @Security Bearer
func (e RsHost) Insert(c *gin.Context) {
	req := dto.RsHostInsertReq{}
	s := service.RsHost{}
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
		e.Error(500, err, fmt.Sprintf("创建RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsHost
// @Summary 修改RsHost
// @Description 修改RsHost
// @Tags RsHost
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsHostUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-host/{id} [put]
// @Security Bearer
func (e RsHost) Update(c *gin.Context) {
	req := dto.RsHostUpdateReq{}
	s := service.RsHost{}
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
		e.Error(500, err, fmt.Sprintf("修改RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsHost
// @Summary 删除RsHost
// @Description 删除RsHost
// @Tags RsHost
// @Param data body dto.RsHostDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-host [delete]
// @Security Bearer
func (e RsHost) Delete(c *gin.Context) {
	s := service.RsHost{}
	req := dto.RsHostDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
