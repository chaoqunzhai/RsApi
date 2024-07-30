package apis

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
)

type RsIdc struct {
	api.Api
}

// GetPage 获取RsIdc列表
// @Summary 获取RsIdc列表
// @Description 获取RsIdc列表
// @Tags RsIdc
// @Param number query string false "机房编号"
// @Param name query string false "机房名称"
// @Param customUser query string false "所属客户"
// @Param ipV6 query string false "是否IPV6"
// @Param typeId query string false "机房类型"
// @Param businessUser query string false "商务人员"
// @Param wechatName query string false "企业微信群名称"
// @Param status query string false "机房状态"
// @Param belong query string false "机房归属"
// @Param isp query string false "运营商"
// @Param charging query string false "计费方式"
// @Param transProvince query string false "是否跨省"
// @Param moreDialing query string false "是否支持多拨"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsIdc}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-idc [get]
// @Security Bearer
func (e RsIdc) GetPage(c *gin.Context) {
	req := dto.RsIdcGetPageReq{}
	s := service.RsIdc{}
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
	list := make([]models.RsIdc, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsIdc
// @Summary 获取RsIdc
// @Description 获取RsIdc
// @Tags RsIdc
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsIdc} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-idc/{id} [get]
// @Security Bearer
func (e RsIdc) Get(c *gin.Context) {
	req := dto.RsIdcGetReq{}
	s := service.RsIdc{}
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
	var object models.RsIdc

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsIdc
// @Summary 创建RsIdc
// @Description 创建RsIdc
// @Tags RsIdc
// @Accept application/json
// @Product application/json
// @Param data body dto.RsIdcInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-idc [post]
// @Security Bearer
func (e RsIdc) Insert(c *gin.Context) {
	req := dto.RsIdcInsertReq{}
	s := service.RsIdc{}
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
		e.Error(500, nil, "机房名称为必填")
		return
	}
	var count int64
	e.Orm.Model(&models.RsIdc{}).Where("name = ?", req.Name).Count(&count)
	if count > 0 {
		e.Error(500, nil, "机房已存在")
		return
	}

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsIdc
// @Summary 修改RsIdc
// @Description 修改RsIdc
// @Tags RsIdc
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsIdcUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-idc/{id} [put]
// @Security Bearer
func (e RsIdc) Update(c *gin.Context) {
	req := dto.RsIdcUpdateReq{}
	s := service.RsIdc{}
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
		e.Error(500, err, fmt.Sprintf("修改RsIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsIdc
// @Summary 删除RsIdc
// @Description 删除RsIdc
// @Tags RsIdc
// @Param data body dto.RsIdcDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-idc [delete]
// @Security Bearer
func (e RsIdc) Delete(c *gin.Context) {
	s := service.RsIdc{}
	req := dto.RsIdcDeleteReq{}
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

	//如果关联了很多主机是不能删除的
	cache := make([]string, 0)
	for _, r := range req.Ids {
		cache = append(cache, fmt.Sprintf("%v", r))
	}
	var count int64
	e.Orm.Raw(fmt.Sprintf("select count(*) from host_bind_idc where idc_id in (%v)", strings.Join(cache, ","))).Scan(&count)

	if count > 0 {
		e.Error(500, nil, "IDC下有关联主机不可删除")
		return
	}

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除RsIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
