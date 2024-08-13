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

type RsContract struct {
	api.Api
}

// GetPage 获取RsContract列表
// @Summary 获取RsContract列表
// @Description 获取RsContract列表
// @Tags RsContract
// @Param name query string false "合同名称"
// @Param number query string false "合同编号"
// @Param buId query int64 false "商务人员"
// @Param customId query int64 false "所属客户ID"
// @Param signatoryId query int64 false "签订人"
// @Param type query int64 false "合同类型,contract_type"
// @Param settlementType query int64 false "结算方式,settlement_type"
// @Param startTime query time.Time false "合同开始时间"
// @Param endTime query time.Time false "合同结束时间"
// @Param address query string false "地址"
// @Param phone query string false "电话"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsContract}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-contract [get]
// @Security Bearer
func (e RsContract) GetPage(c *gin.Context) {
	req := dto.RsContractGetPageReq{}
	s := service.RsContract{}
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
	list := make([]models.RsContract, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsContract失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsContract
// @Summary 获取RsContract
// @Description 获取RsContract
// @Tags RsContract
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsContract} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-contract/{id} [get]
// @Security Bearer
func (e RsContract) Get(c *gin.Context) {
	req := dto.RsContractGetReq{}
	s := service.RsContract{}
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
	var object models.RsContract

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsContract失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsContract
// @Summary 创建RsContract
// @Description 创建RsContract
// @Tags RsContract
// @Accept application/json
// @Product application/json
// @Param data body dto.RsContractInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-contract [post]
// @Security Bearer
func (e RsContract) Insert(c *gin.Context) {
	req := dto.RsContractInsertReq{}
	s := service.RsContract{}
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
	if req.Name == "" {
		e.Error(500, nil, "合同名称不可为空")
		return
	}
	var count int64
	e.Orm.Model(&models.RsContract{}).Where("name = ?", req.Name).Count(&count)
	if count > 0 {
		e.Error(500, nil, "合同名称已经存在")
		return
	}
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsContract失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsContract
// @Summary 修改RsContract
// @Description 修改RsContract
// @Tags RsContract
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsContractUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-contract/{id} [put]
// @Security Bearer
func (e RsContract) Update(c *gin.Context) {
	req := dto.RsContractUpdateReq{}
	s := service.RsContract{}
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
		e.Error(500, err, fmt.Sprintf("修改RsContract失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsContract
// @Summary 删除RsContract
// @Description 删除RsContract
// @Tags RsContract
// @Param data body dto.RsContractDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-contract [delete]
// @Security Bearer
func (e RsContract) Delete(c *gin.Context) {
	s := service.RsContract{}
	req := dto.RsContractDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsContract失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
