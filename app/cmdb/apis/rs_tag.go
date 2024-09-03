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

type RsTag struct {
	api.Api
}

// GetPage 获取RsTag列表
// @Summary 获取RsTag列表
// @Description 获取RsTag列表
// @Tags RsTag
// @Param enable query string false "开关"
// @Param name query string false "业务云名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsTag}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-tag [get]
// @Security Bearer
func (e RsTag) GetPage(c *gin.Context) {
	req := dto.RsTagGetPageReq{}
	s := service.RsTag{}
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
	list := make([]models.RsTag, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsTag失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsTag
// @Summary 获取RsTag
// @Description 获取RsTag
// @Tags RsTag
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsTag} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-tag/{id} [get]
// @Security Bearer
func (e RsTag) Get(c *gin.Context) {
	req := dto.RsTagGetReq{}
	s := service.RsTag{}
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
	var object models.RsTag

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsTag失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsTag
// @Summary 创建RsTag
// @Description 创建RsTag
// @Tags RsTag
// @Accept application/json
// @Product application/json
// @Param data body dto.RsTagInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-tag [post]
// @Security Bearer
func (e RsTag) Insert(c *gin.Context) {
	req := dto.RsTagInsertReq{}
	s := service.RsTag{}
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
		e.Error(500, err, fmt.Sprintf("创建RsTag失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsTag
// @Summary 修改RsTag
// @Description 修改RsTag
// @Tags RsTag
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsTagUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-tag/{id} [put]
// @Security Bearer
func (e RsTag) Update(c *gin.Context) {
	req := dto.RsTagUpdateReq{}
	s := service.RsTag{}
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
		e.Error(500, err, fmt.Sprintf("修改RsTag失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsTag
// @Summary 删除RsTag
// @Description 删除RsTag
// @Tags RsTag
// @Param data body dto.RsTagDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-tag [delete]
// @Security Bearer
func (e RsTag) Delete(c *gin.Context) {
	s := service.RsTag{}
	req := dto.RsTagDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsTag失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
