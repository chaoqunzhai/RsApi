package apis

import (
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/global"

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
// @Param number query int64 false "机房编号"
// @Param name query string false "机房名称"
// @Param customUser query int64 false "所属客户"
// @Param typeId query int64 false "机房类型"
// @Param businessUser query int64 false "商务人员"
// @Param status query int64 false "机房状态"
// @Param belong query int64 false "机房归属"
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
	idcMap := make(map[int]models.RsIdcCount, 0)

	var idcList []int
	for _, idc := range list {
		idcList = append(idcList, idc.Id)
		idcMap[idc.Id] = models.RsIdcCount{}
	}

	var hostList []models.RsHost
	fmt.Println("req.PageSize", req.PageSize)
	if req.PageSize == -1 || req.PageSize >= 1000 {
		e.Orm.Model(&models.RsHost{}).Select("idc,status").Find(&hostList)
	} else {
		e.Orm.Model(&models.RsHost{}).Select("idc,status").Where("idc in ?", idcList).Find(&hostList)
	}

	for _, host := range hostList {
		RsIdcCount, ok := idcMap[host.Idc]
		if !ok {
			continue
		}

		if host.Status == global.HostSuccess {
			RsIdcCount.OnLine += 1
		}
		if host.Status == global.HostOffline {
			RsIdcCount.Offline += 1
		}
		RsIdcCount.AllHost += 1

		idcMap[host.Idc] = RsIdcCount
	}
	var result []interface{}

	for _, idc := range list {
		RsIdcCount, ok := idcMap[idc.Id]
		if !ok {
			continue
		}
		idc.RsIdcCount = RsIdcCount
		result = append(result, idc)
	}

	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
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
	var count int64
	e.Orm.Model(&models.RsIdc{}).Where("name = ?", req.Name).Count(&count)
	if count > 0 {
		e.Error(500, nil, "机房名称已存在")
		return
	}
	modelId, err := s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "POST",
		Module:     "rs_idc",
		ObjectId:   modelId,
		TargetId:   modelId,
		Info:       "创建IDC",
	})

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
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "PUT",
		Module:     "rs_idc",
		ObjectId:   req.Id,
		TargetId:   req.Id,
		Info:       "更新IDC",
	})
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

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除RsIdc失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "DELETE",
		Module:     "rs_idc",
		ObjectId:   req.Ids[0],
		TargetId:   req.Ids[0],
		Info:       "删除IDC",
	})
	e.OK(req.GetId(), "删除成功")
}
