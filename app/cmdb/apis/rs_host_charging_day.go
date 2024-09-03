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

type RsHostChargingDay struct {
	api.Api
}

// GetPage 获取RsHostChargingDay列表
// @Summary 获取RsHostChargingDay列表
// @Description 获取RsHostChargingDay列表
// @Tags RsHostChargingDay
// @Param businessId query string false "切换的业务ID"
// @Param hostId query string false "关联的主机ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsHostChargingDay}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host-charging-day [get]
// @Security Bearer
func (e RsHostChargingDay) GetPage(c *gin.Context) {
	req := dto.RsHostChargingDayGetPageReq{}
	s := service.RsHostChargingDay{}
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
	list := make([]models.RsHostChargingDay, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHostChargingDay失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsHostChargingDay
// @Summary 获取RsHostChargingDay
// @Description 获取RsHostChargingDay
// @Tags RsHostChargingDay
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsHostChargingDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host-charging-day/{id} [get]
// @Security Bearer
func (e RsHostChargingDay) Get(c *gin.Context) {
	req := dto.RsHostChargingDayGetReq{}
	s := service.RsHostChargingDay{}
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
	var object models.RsHostChargingDay

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHostChargingDay失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsHostChargingDay
// @Summary 创建RsHostChargingDay
// @Description 创建RsHostChargingDay
// @Tags RsHostChargingDay
// @Accept application/json
// @Product application/json
// @Param data body dto.RsHostChargingDayInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-host-charging-day [post]
// @Security Bearer
func (e RsHostChargingDay) Insert(c *gin.Context) {
	req := dto.RsHostChargingDayInsertReq{}
	s := service.RsHostChargingDay{}
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
		e.Error(500, err, fmt.Sprintf("创建RsHostChargingDay失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsHostChargingDay
// @Summary 修改RsHostChargingDay
// @Description 修改RsHostChargingDay
// @Tags RsHostChargingDay
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsHostChargingDayUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-host-charging-day/{id} [put]
// @Security Bearer
func (e RsHostChargingDay) Update(c *gin.Context) {
	req := dto.RsHostChargingDayUpdateReq{}
	s := service.RsHostChargingDay{}
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
		e.Error(500, err, fmt.Sprintf("修改RsHostChargingDay失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsHostChargingDay
// @Summary 删除RsHostChargingDay
// @Description 删除RsHostChargingDay
// @Tags RsHostChargingDay
// @Param data body dto.RsHostChargingDayDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-host-charging-day [delete]
// @Security Bearer
func (e RsHostChargingDay) Delete(c *gin.Context) {
	s := service.RsHostChargingDay{}
	req := dto.RsHostChargingDayDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsHostChargingDay失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
