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

type RsContacts struct {
	api.Api
}

// GetPage 获取RsContacts列表
// @Summary 获取RsContacts列表
// @Description 获取RsContacts列表
// @Tags RsContacts
// @Param userName query string false "用户名"
// @Param customerId query int64 false "客户ID"
// @Param buUser query int64 false "商务人员"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsContacts}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-contacts [get]
// @Security Bearer
func (e RsContacts) GetPage(c *gin.Context) {
    req := dto.RsContactsGetPageReq{}
    s := service.RsContacts{}
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
	list := make([]models.RsContacts, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsContacts失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsContacts
// @Summary 获取RsContacts
// @Description 获取RsContacts
// @Tags RsContacts
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsContacts} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-contacts/{id} [get]
// @Security Bearer
func (e RsContacts) Get(c *gin.Context) {
	req := dto.RsContactsGetReq{}
	s := service.RsContacts{}
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
	var object models.RsContacts

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsContacts失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建RsContacts
// @Summary 创建RsContacts
// @Description 创建RsContacts
// @Tags RsContacts
// @Accept application/json
// @Product application/json
// @Param data body dto.RsContactsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-contacts [post]
// @Security Bearer
func (e RsContacts) Insert(c *gin.Context) {
    req := dto.RsContactsInsertReq{}
    s := service.RsContacts{}
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
		e.Error(500, err, fmt.Sprintf("创建RsContacts失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsContacts
// @Summary 修改RsContacts
// @Description 修改RsContacts
// @Tags RsContacts
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsContactsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-contacts/{id} [put]
// @Security Bearer
func (e RsContacts) Update(c *gin.Context) {
    req := dto.RsContactsUpdateReq{}
    s := service.RsContacts{}
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
		e.Error(500, err, fmt.Sprintf("修改RsContacts失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除RsContacts
// @Summary 删除RsContacts
// @Description 删除RsContacts
// @Tags RsContacts
// @Param data body dto.RsContactsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-contacts [delete]
// @Security Bearer
func (e RsContacts) Delete(c *gin.Context) {
    s := service.RsContacts{}
    req := dto.RsContactsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsContacts失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
