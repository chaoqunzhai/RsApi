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

type RsCustom struct {
	api.Api
}

// GetPage 获取RsCustom列表
// @Summary 获取RsCustom列表
// @Description 获取RsCustom列表
// @Tags RsCustom
// @Param name query string false "客户名称"
// @Param type query int64 false "客户类型,customer_type"
// @Param cooperation query int64 false "合作状态,work_status"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsCustom}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-custom [get]
// @Security Bearer
func (e RsCustom) GetPage(c *gin.Context) {
    req := dto.RsCustomGetPageReq{}
    s := service.RsCustom{}
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
	list := make([]models.RsCustom, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsCustom失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsCustom
// @Summary 获取RsCustom
// @Description 获取RsCustom
// @Tags RsCustom
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsCustom} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-custom/{id} [get]
// @Security Bearer
func (e RsCustom) Get(c *gin.Context) {
	req := dto.RsCustomGetReq{}
	s := service.RsCustom{}
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
	var object models.RsCustom

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsCustom失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建RsCustom
// @Summary 创建RsCustom
// @Description 创建RsCustom
// @Tags RsCustom
// @Accept application/json
// @Product application/json
// @Param data body dto.RsCustomInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-custom [post]
// @Security Bearer
func (e RsCustom) Insert(c *gin.Context) {
    req := dto.RsCustomInsertReq{}
    s := service.RsCustom{}
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
		e.Error(500, err, fmt.Sprintf("创建RsCustom失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsCustom
// @Summary 修改RsCustom
// @Description 修改RsCustom
// @Tags RsCustom
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsCustomUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-custom/{id} [put]
// @Security Bearer
func (e RsCustom) Update(c *gin.Context) {
    req := dto.RsCustomUpdateReq{}
    s := service.RsCustom{}
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
		e.Error(500, err, fmt.Sprintf("修改RsCustom失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除RsCustom
// @Summary 删除RsCustom
// @Description 删除RsCustom
// @Tags RsCustom
// @Param data body dto.RsCustomDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-custom [delete]
// @Security Bearer
func (e RsCustom) Delete(c *gin.Context) {
    s := service.RsCustom{}
    req := dto.RsCustomDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsCustom失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
