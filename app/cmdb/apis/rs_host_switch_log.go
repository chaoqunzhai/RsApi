package apis

import (
	"fmt"
	"go-admin/common/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
)

type RsHostSwitchLog struct {
	api.Api
}

// GetPage 获取RsHostSwitchLog列表
// @Summary 获取RsHostSwitchLog列表
// @Description 获取RsHostSwitchLog列表
// @Tags RsHostSwitchLog
// @Param hostId query string false "切换的主机ID"
// @Param businessId query string false "切换的新业务ID"
// @Param businessSn query string false "原来的业务SN"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsHostSwitchLog}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host-switch-log [get]
// @Security Bearer
func (e RsHostSwitchLog) GetPage(c *gin.Context) {
	req := dto.RsHostSwitchLogGetPageReq{}
	s := service.RsHostSwitchLog{}
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
	list := make([]models.RsHostSwitchLog, 0)
	var count int64

	var businessLis []int

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHostSwitchLog失败，\r\n失败信息 %s", err.Error()))
		return
	}

	for _, row := range list {
		if row.BuTargetId > 0 {
			businessLis = append(businessLis, row.BuTargetId)
		}
	}
	businessLis = utils.RemoveRepeatInt(businessLis)
	BusinessMapData := service.GetBusinessMap(e.Orm, businessLis)
	result := make([]interface{}, 0)
	for _, row := range list {

		buInfo, ok := BusinessMapData[row.BuTargetId]
		if ok {
			row.BusinessInfo = buInfo
		}
		result = append(result, row)

	}

	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsHostSwitchLog
// @Summary 获取RsHostSwitchLog
// @Description 获取RsHostSwitchLog
// @Tags RsHostSwitchLog
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsHostSwitchLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host-switch-log/{id} [get]
// @Security Bearer
func (e RsHostSwitchLog) Get(c *gin.Context) {
	req := dto.RsHostSwitchLogGetReq{}
	s := service.RsHostSwitchLog{}
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
	var object models.RsHostSwitchLog

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHostSwitchLog失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Delete 删除RsHostSwitchLog
// @Summary 删除RsHostSwitchLog
// @Description 删除RsHostSwitchLog
// @Tags RsHostSwitchLog
// @Param data body dto.RsHostSwitchLogDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-host-switch-log [delete]
// @Security Bearer
func (e RsHostSwitchLog) Delete(c *gin.Context) {
	s := service.RsHostSwitchLog{}
	req := dto.RsHostSwitchLogDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsHostSwitchLog失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
