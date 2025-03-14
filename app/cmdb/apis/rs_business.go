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

type RsBusiness struct {
	api.Api
}

// GetPage 获取RsBusiness列表
// @Summary 获取RsBusiness列表
// @Description 获取RsBusiness列表
// @Tags RsBusiness
// @Param enable query string false "开关"
// @Param name query string false "业务云名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsBusiness}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-business [get]
// @Security Bearer
func (e RsBusiness) GetPage(c *gin.Context) {
	req := dto.RsBusinessGetPageReq{}
	s := service.RsBusiness{}
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
	list := make([]models.RsBusiness, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsBusiness失败，\r\n失败信息 %s", err.Error()))
		return
	}
	result := make([]interface{}, 0)

	for _, row := range list {
		if req.TreeTag > 0 {
			row.Children = s.GetChildren(row.Id)
		}

		result = append(result, row)

	}

	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsBusiness
// @Summary 获取RsBusiness
// @Description 获取RsBusiness
// @Tags RsBusiness
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsBusiness} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-business/{id} [get]
// @Security Bearer
func (e RsBusiness) Get(c *gin.Context) {
	req := dto.RsBusinessGetReq{}
	s := service.RsBusiness{}
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
	var object models.RsBusiness

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsBusiness失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsBusiness
// @Summary 创建RsBusiness
// @Description 创建RsBusiness
// @Tags RsBusiness
// @Accept application/json
// @Product application/json
// @Param data body dto.RsBusinessInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-business [post]
// @Security Bearer
func (e RsBusiness) Insert(c *gin.Context) {
	req := dto.RsBusinessInsertReq{}
	s := service.RsBusiness{}
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
	e.Orm.Model(&models.RsBusiness{}).Where("name = ?", req.Name).Count(&count)
	if count > 0 {
		e.Error(500, nil, "业务名称已经存在")
		return
	}
	var EnCount int64
	e.Orm.Model(&models.RsBusiness{}).Where("en_name = ?", req.Name).Count(&EnCount)
	if EnCount > 0 {
		e.Error(500, nil, "英文业务名称已经存在")
		return
	}
	modelId, err := s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsBusiness失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "POST",
		Module:     "rs_business",
		ObjectId:   modelId,
		TargetId:   modelId,
		Info:       "创建业务",
	})
	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsBusiness
// @Summary 修改RsBusiness
// @Description 修改RsBusiness
// @Tags RsBusiness
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsBusinessUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-business/{id} [put]
// @Security Bearer
func (e RsBusiness) Update(c *gin.Context) {
	req := dto.RsBusinessUpdateReq{}
	s := service.RsBusiness{}
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

	if req.ParentId == req.Id {
		e.Error(500, nil, "父级业务不可为同一个")
		return
	}
	var count int64
	e.Orm.Model(&models.RsBusiness{}).Where("name = ? ", req.Name).Count(&count)
	if count > 1 {
		e.Error(500, nil, "已经存在")
		return
	}
	userName := user.GetUserName(c)
	err = s.Update(userName, &req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改RsBusiness失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsBusiness
// @Summary 删除RsBusiness
// @Description 删除RsBusiness
// @Tags RsBusiness
// @Param data body dto.RsBusinessDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-business [delete]
// @Security Bearer
func (e RsBusiness) Delete(c *gin.Context) {
	s := service.RsBusiness{}
	req := dto.RsBusinessDeleteReq{}
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

	//需要校验是否已经被主机关联了业务

	var count int64

	e.Orm.Raw(fmt.Sprintf("select count(*) from host_bind_business where `business_id` in (%v)", req.GetIdStr())).Scan(&count)

	if count > 0 {
		e.Error(500, nil, fmt.Sprintf("业务已经关联了主机,无法删除"))
		return
	}
	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除RsBusiness失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Module:     "rs_business",
		Action:     "DELETE",
		ObjectId:   req.Ids[0],
		TargetId:   req.Ids[0],
		Info:       "更新业务信息",
	})
	e.OK(req.GetId(), "删除成功")
}
