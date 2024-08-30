package apis

import (
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
)

type RsCustomUser struct {
	api.Api
}

// GetPage 获取RsCustomUser列表
// @Summary 获取RsCustomUser列表
// @Description 获取RsCustomUser列表
// @Tags RsCustomUser
// @Param userName query string false "姓名"
// @Param customId query string false "所属客户"
// @Param buId query string false "所属商务人员"
// @Param phone query string false "联系号码"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsCustomUser}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-custom-user [get]
// @Security Bearer
func (e RsCustomUser) GetPage(c *gin.Context) {
	req := dto.RsCustomUserGetPageReq{}
	s := service.RsCustomUser{}
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
	list := make([]models.RsCustomUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsCustomUser失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsCustomUser
// @Summary 获取RsCustomUser
// @Description 获取RsCustomUser
// @Tags RsCustomUser
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsCustomUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-custom-user/{id} [get]
// @Security Bearer
func (e RsCustomUser) Get(c *gin.Context) {
	req := dto.RsCustomUserGetReq{}
	s := service.RsCustomUser{}
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
	var object models.RsCustomUser

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsCustomUser失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsCustomUser
// @Summary 创建RsCustomUser
// @Description 创建RsCustomUser
// @Tags RsCustomUser
// @Accept application/json
// @Product application/json
// @Param data body dto.RsCustomUserInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-custom-user [post]
// @Security Bearer
func (e RsCustomUser) Insert(c *gin.Context) {
	req := dto.RsCustomUserInsertReq{}
	s := service.RsCustomUser{}
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

	var count int64

	e.Orm.Model(&models.RsCustomUser{}).Where("custom_id = ? and user_name = ?", req.CustomId, req.UserName).Count(&count)
	if count > 0 {
		e.Error(500, nil, "该客户联系人已经存在")
		return
	}
	modelId, err := s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsCustomUser失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "POST",
		Module:     "rs_custom_user",
		ObjectId:   modelId,
		TargetId:   modelId,
		Info:       "创建联系人信息",
	})
	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsCustomUser
// @Summary 修改RsCustomUser
// @Description 修改RsCustomUser
// @Tags RsCustomUser
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsCustomUserUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-custom-user/{id} [put]
// @Security Bearer
func (e RsCustomUser) Update(c *gin.Context) {
	req := dto.RsCustomUserUpdateReq{}
	s := service.RsCustomUser{}
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
		e.Error(500, err, fmt.Sprintf("修改RsCustomUser失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "PUT",
		Module:     "rs_custom_user",
		ObjectId:   req.Id,
		TargetId:   req.Id,
		Info:       "更新联系人信息",
	})
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsCustomUser
// @Summary 删除RsCustomUser
// @Description 删除RsCustomUser
// @Tags RsCustomUser
// @Param data body dto.RsCustomUserDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-custom-user [delete]
// @Security Bearer
func (e RsCustomUser) Delete(c *gin.Context) {
	s := service.RsCustomUser{}
	req := dto.RsCustomUserDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsCustomUser失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "DELETE",
		Module:     "rs_custom_user",
		ObjectId:   req.Ids[0],
		TargetId:   req.Ids[0],
		Info:       "删除联系人信息",
	})
	e.OK(req.GetId(), "删除成功")
}
