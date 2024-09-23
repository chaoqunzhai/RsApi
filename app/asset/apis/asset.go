package apis

import (
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"
	cDto "go-admin/common/dto"
	"strconv"
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

	combinationId := c.Query("combinationId")
	fmt.Println("combinationId", combinationId)
	p := actions.GetPermissionFromContext(c)
	list := make([]models.AdditionsWarehousing, 0)
	var count int64

	err = s.GetPage(combinationId, &req, p, &list, &count)
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

	orm := e.Orm.Model(&data)
	if req.Name != "" {
		var assetList []models.AdditionsWarehousing
		e.Orm.Model(&models.AdditionsWarehousing{}).Select("w_id").Where("name like ?", "%"+req.Name+"%").Find(&assetList)

		var assetWinds []int64
		for _, asset := range assetList {
			assetWinds = append(assetWinds, asset.WId)
		}
		orm = orm.Where("id in (?)", assetWinds)
	}
	err = orm.Scopes(
		cDto.MakeCondition(req.GetNeedSearch()),
		cDto.Paginate(req.GetPageSize(), req.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).
		Find(&list).Limit(-1).Offset(-1).
		Count(&count).Error

	result := make([]interface{}, 0)

	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取AdditionsWarehousing失败，\r\n失败信息 %s", err.Error()))
		return
	}
	idS := make([]int, 0)
	for _, v := range list {
		idS = append(idS, v.Id)
	}
	var asset []models.AdditionsWarehousing
	e.Orm.Model(&models.AdditionsWarehousing{}).Select("name,w_id,id").Where("w_id in ?", idS).Find(&asset)
	assetMap := make(map[int64][]string, 0)
	for _, v := range asset {

		assetList, ok := assetMap[v.WId]
		if !ok {
			assetList = make([]string, 0)
		}
		assetList = append(assetList, v.Name)
		assetMap[v.WId] = assetList
	}

	for _, v := range list {
		assetList, ok := assetMap[int64(v.Id)]

		if ok {
			v.Asset = assetList
		} else {
			v.Asset = make([]interface{}, 0)
		}
		result = append(result, v)
	}

	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
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

func (e AdditionsWarehousing) GetStore(c *gin.Context) {
	req := dto.AdditionsOrderGetReq{}
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
	orderId := c.Param("id")

	pageSizeInt, _ := strconv.Atoi(c.Query("pageSize"))
	pageIndexInt, _ := strconv.Atoi(c.Query("pageIndex"))
	var object models.AdditionsOrder

	e.Orm.Model(models.AdditionsOrder{}).Where("id = ?", orderId).Limit(1).Find(&object)
	if object.Id == 0 {
		e.Error(500, nil, "入库单不存在")
		return
	}

	var count int64
	var bindAsset []models.AdditionsWarehousing
	e.Orm.Model(models.AdditionsWarehousing{}).Where("w_id = ?", orderId).Scopes(
		cDto.Paginate(pageSizeInt, pageIndexInt),
	).Find(&bindAsset).Limit(-1).Offset(-1).Count(&count)
	object.Asset = bindAsset
	e.PageOK(object, int(count), pageIndexInt, pageSizeInt, "")
	return
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
	if len(req.Asset) == 0 {
		e.Error(500, nil, "资产列表不存在")
		return
	}
	order := models.AdditionsOrder{
		OrderId:     fmt.Sprintf("%v", time.Now().Unix()),
		StoreRoomId: req.StoreRoomId,
	}
	order.CreateBy = user.GetUserId(c)

	e.Orm.Create(&order)

	var userModel models2.SysUser
	e.Orm.Model(&models2.SysUser{}).Where("user_id = ?", order.CreateBy).Limit(1).Find(&userModel)

	for _, row := range req.Asset {
		err = s.Insert(userModel.Username, order.Id, req.StoreRoomId, &row)
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

	uid := c.Param("id")
	var order models.AdditionsWarehousing

	e.Orm.Model(&order).Where("id = ?", uid).Limit(1).Find(&order)
	if order.Id == 0 {
		e.Error(500, nil, "数据不存在")
		return
	}

	p := actions.GetPermissionFromContext(c)
	err = s.Update(uid, &req, p)

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

	var order models.AdditionsOrder

	e.Orm.Model(&order).Where("id = ?", req.Id).Limit(1).Find(&order)
	if order.Id == 0 {
		e.Error(500, nil, "数据不存在")
		return
	}
	order.StoreRoomId = req.StoreRoomId

	e.Orm.Save(&order)

	p := actions.GetPermissionFromContext(c)

	for _, row := range req.Asset {
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

func (e AdditionsWarehousing) StoreDelete(c *gin.Context) {
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

	p := actions.GetPermissionFromContext(c)

	var data models.AdditionsOrder

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, req.GetId())
	if err = db.Error; err != nil {
		e.Error(500, err, fmt.Sprintf("删除AdditionsWarehousing失败，\r\n失败信息 %s", err.Error()))
		return
	}
	if req.Unscoped == 1 {
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("w_id in ?", req.GetId()).Delete(&models.AdditionsWarehousing{})
	}

	e.OK(req.GetId(), "删除成功")
}
