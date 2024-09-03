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

type RsDial struct {
	api.Api
}

// GetPage 获取RsDial列表
// @Summary 获取RsDial列表
// @Description 获取RsDial列表
// @Tags RsDial
// @Param customId query int64 false "所属客户"
// @Param contractId query int64 false "关联合同"
// @Param broadbandType query int64 false "带宽类型,broadband_type"
// @Param isManager query int64 false "是否管理线"
// @Param ip query string false "IP地址"
// @Param dialName query string false "线路名称"
// @Param networkingStatus query int64 false "拨号状态,1:已联网 0:未联网 -1:联网异常"
// @Param status query int64 false "拨号状态,1:已拨通 0:待使用 -1:拨号异常"
// @Param source query int64 false "拨号状态,0:录入 1:自动上报"
// @Param idcId query int64 false "关联的IDC"
// @Param hostId query int64 false "关联主机ID"
// @Param deviceId query int64 false "关联网卡ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsDial}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-dial [get]
// @Security Bearer
func (e RsDial) GetPage(c *gin.Context) {
	req := dto.RsDialGetPageReq{}
	s := service.RsDial{}
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
	list := make([]models.RsDial, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsDial失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsDial
// @Summary 获取RsDial
// @Description 获取RsDial
// @Tags RsDial
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsDial} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-dial/{id} [get]
// @Security Bearer
func (e RsDial) Get(c *gin.Context) {
	req := dto.RsDialGetReq{}
	s := service.RsDial{}
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
	var object models.RsDial

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsDial失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsDial
// @Summary 创建RsDial
// @Description 创建RsDial
// @Tags RsDial
// @Accept application/json
// @Product application/json
// @Param data body dto.RsDialInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-dial [post]
// @Security Bearer
func (e RsDial) Insert(c *gin.Context) {
	req := dto.RsDialInsertReq{}
	s := service.RsDial{}
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
	if req.Account == "" {
		e.Error(500, nil, "账号名称不可为空")
		return
	}
	var count int64
	e.Orm.Model(&models.RsDial{}).Where("account = ?", req.Account).Count(&count)
	if count > 0 {
		e.Error(500, nil, "拨号已存在")
		return
	}
	modelId, err := s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsDial失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "POST",
		Module:     "rs_dial",
		ObjectId:   modelId,
		TargetId:   modelId,
		Info:       "创建拨号信息",
	})

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsDial
// @Summary 修改RsDial
// @Description 修改RsDial
// @Tags RsDial
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsDialUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-dial/{id} [put]
// @Security Bearer
func (e RsDial) Update(c *gin.Context) {
	req := dto.RsDialUpdateReq{}
	s := service.RsDial{}
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
		e.Error(500, err, fmt.Sprintf("修改RsDial失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "PUT",
		Module:     "rs_dial",
		ObjectId:   req.Id,
		TargetId:   req.Id,
		Info:       "更新拨号",
	})
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsDial
// @Summary 删除RsDial
// @Description 删除RsDial
// @Tags RsDial
// @Param data body dto.RsDialDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-dial [delete]
// @Security Bearer
func (e RsDial) Delete(c *gin.Context) {
	s := service.RsDial{}
	req := dto.RsDialDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsDial失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "DELETE",
		Module:     "rs_dial",
		ObjectId:   req.Ids[0],
		TargetId:   req.Ids[0],
		Info:       "删除拨号",
	})
	e.OK(req.GetId(), "删除成功")
}
