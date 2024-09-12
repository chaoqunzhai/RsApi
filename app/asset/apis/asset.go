package apis

import (
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"
	cDto "go-admin/common/dto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/asset/models"
	"go-admin/app/asset/service"
	"go-admin/app/asset/service/dto"
	"go-admin/common/actions"
)

type AdditionsWarehousing struct {
	api.Api
}

// GetPage 获取AdditionsWarehousing列表
// @Summary 获取AdditionsWarehousing列表
// @Description 获取AdditionsWarehousing列表
// @Tags AdditionsWarehousing
// @Param categoryId query int64 false "关联的资产分类ID"
// @Param supplierId query int64 false "供应商ID"
// @Param wId query int64 false "关联的入库单号"
// @Param name query string false "资产名称"
// @Param spec query string false "规格型号"
// @Param brand query string false "品牌名称"
// @Param sn query string false "资产SN"
// @Param userId query string false "采购人员ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AdditionsWarehousing}} "{"code": 200, "data": [...]}"
// @Router /api/v1/additions-warehousing [get]
// @Security Bearer
func (e AdditionsWarehousing) GetPage(c *gin.Context) {
	req := dto.AdditionsWarehousingGetPageReq{}
	s := service.AdditionsWarehousing{}
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
	list := make([]models.AdditionsWarehousing, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取AdditionsWarehousing失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e AdditionsWarehousing) GetStorePage(c *gin.Context) {
	req := dto.AdditionsOrderGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.AdditionsOrder, 0)
	var count int64

	var data models.AdditionsOrder

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(req.GetNeedSearch()),
			cDto.Paginate(req.GetPageSize(), req.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(&list).Limit(-1).Offset(-1).
		Count(&count).Error

	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取AdditionsWarehousing失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取AdditionsWarehousing
// @Summary 获取AdditionsWarehousing
// @Description 获取AdditionsWarehousing
// @Tags AdditionsWarehousing
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AdditionsWarehousing} "{"code": 200, "data": [...]}"
// @Router /api/v1/additions-warehousing/{id} [get]
// @Security Bearer
func (e AdditionsWarehousing) Get(c *gin.Context) {
	req := dto.AdditionsWarehousingGetReq{}
	s := service.AdditionsWarehousing{}
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
	var object models.AdditionsWarehousing

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取AdditionsWarehousing失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建AdditionsWarehousing
// @Summary 创建AdditionsWarehousing
// @Description 创建AdditionsWarehousing
// @Tags AdditionsWarehousing
// @Accept application/json
// @Product application/json
// @Param data body dto.AdditionsWarehousingInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/additions-warehousing [post]
// @Security Bearer
func (e AdditionsWarehousing) Insert(c *gin.Context) {
	req := dto.AssetInsertReq{}
	s := service.AdditionsWarehousing{}
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
	if len(req.List) == 0 {
		e.Error(500, nil, "资产列表不存在")
		return
	}
	order := models2.AdditionsOrder{
		OrderId:     fmt.Sprintf("%v", time.Now().Unix()),
		StoreRoomId: req.StoreRoomId,
	}
	order.CreateBy = user.GetUserId(c)

	e.Orm.Create(&order)
	for _, row := range req.List {
		err = s.Insert(order.Id, req.StoreRoomId, &row)
		if err != nil {
			fmt.Println("err!", err)
			continue
		}

	}

	e.OK("", "创建成功")
}

func (e AdditionsWarehousing) Update(c *gin.Context) {
	req := dto.AdditionsWarehousingUpdateReq{}
	s := service.AdditionsWarehousing{}
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

	var order models.AdditionsWarehousing

	e.Orm.Model(&order).Where("id = ?", req.Id).Limit(1).Find(&order)
	if order.Id == 0 {
		e.Error(500, nil, "数据不存在")
		return
	}

	p := actions.GetPermissionFromContext(c)
	err = s.Update(&req, p)

	e.OK("", "修改成功")
}

// Update 修改AdditionsWarehousing
// @Summary 修改AdditionsWarehousing
// @Description 修改AdditionsWarehousing
// @Tags AdditionsWarehousing
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AdditionsWarehousingUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/additions-warehousing/{id} [put]
// @Security Bearer
func (e AdditionsWarehousing) UpdateStore(c *gin.Context) {
	req := dto.AssetUpdateReq{}
	s := service.AdditionsWarehousing{}
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

	var order models2.AdditionsOrder

	e.Orm.Model(&order).Where("id = ?", req.Id).Limit(1).Find(&order)
	if order.Id == 0 {
		e.Error(500, nil, "数据不存在")
		return
	}
	order.StoreRoomId = req.StoreRoomId

	e.Orm.Save(&order)

	p := actions.GetPermissionFromContext(c)

	for _, row := range req.List {
		err = s.UpdateStore(req.StoreRoomId, &row, p)
		if err != nil {
			continue
		}
	}

	e.OK("", "修改成功")
}

// Delete 删除AdditionsWarehousing
// @Summary 删除AdditionsWarehousing
// @Description 删除AdditionsWarehousing
// @Tags AdditionsWarehousing
// @Param data body dto.AdditionsWarehousingDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/additions-warehousing [delete]
// @Security Bearer
func (e AdditionsWarehousing) Delete(c *gin.Context) {
	s := service.AdditionsWarehousing{}
	req := dto.AdditionsWarehousingDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除AdditionsWarehousing失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
