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

type RsCustomer struct {
	api.Api
}

// GetPage 获取RsCustomer列表
// @Summary 获取RsCustomer列表
// @Description 获取RsCustomer列表
// @Tags RsCustomer
// @Param name query string false "客户名称"
// @Param level query int64 false "客户等级"
// @Param typeId query int64 false "客户类型"
// @Param workStatus query string false "合作状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsCustomer}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-customer [get]
// @Security Bearer
func (e RsCustomer) GetPage(c *gin.Context) {
    req := dto.RsCustomerGetPageReq{}
    s := service.RsCustomer{}
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
	list := make([]models.RsCustomer, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsCustomer失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsCustomer
// @Summary 获取RsCustomer
// @Description 获取RsCustomer
// @Tags RsCustomer
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsCustomer} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-customer/{id} [get]
// @Security Bearer
func (e RsCustomer) Get(c *gin.Context) {
	req := dto.RsCustomerGetReq{}
	s := service.RsCustomer{}
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
	var object models.RsCustomer

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsCustomer失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建RsCustomer
// @Summary 创建RsCustomer
// @Description 创建RsCustomer
// @Tags RsCustomer
// @Accept application/json
// @Product application/json
// @Param data body dto.RsCustomerInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-customer [post]
// @Security Bearer
func (e RsCustomer) Insert(c *gin.Context) {
    req := dto.RsCustomerInsertReq{}
    s := service.RsCustomer{}
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
		e.Error(500, err, fmt.Sprintf("创建RsCustomer失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsCustomer
// @Summary 修改RsCustomer
// @Description 修改RsCustomer
// @Tags RsCustomer
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsCustomerUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-customer/{id} [put]
// @Security Bearer
func (e RsCustomer) Update(c *gin.Context) {
    req := dto.RsCustomerUpdateReq{}
    s := service.RsCustomer{}
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
		e.Error(500, err, fmt.Sprintf("修改RsCustomer失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除RsCustomer
// @Summary 删除RsCustomer
// @Description 删除RsCustomer
// @Tags RsCustomer
// @Param data body dto.RsCustomerDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-customer [delete]
// @Security Bearer
func (e RsCustomer) Delete(c *gin.Context) {
    s := service.RsCustomer{}
    req := dto.RsCustomerDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsCustomer失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
