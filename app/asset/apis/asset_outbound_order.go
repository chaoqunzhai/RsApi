package apis

import (
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"
	cDto "go-admin/common/dto"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/asset/models"
	"go-admin/app/asset/service"
	"go-admin/app/asset/service/dto"
	"go-admin/common/actions"
)

type AssetOutboundOrder struct {
	api.Api
}

// GetPage 获取AssetOutboundOrder列表
// @Summary 获取AssetOutboundOrder列表
// @Description 获取AssetOutboundOrder列表
// @Tags AssetOutboundOrder
// @Param code query string false "出库编码"
// @Param customId query string false "所属客户ID"
// @Param region query string false "省份城市多ID"
// @Param ems query string false "物流公司"
// @Param trackingNumber query string false "物流单号"
// @Param address query string false "联系地址"
// @Param user query string false "联系人"
// @Param idcId query string false "idcId"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetOutboundOrder}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-outbound-order [get]
// @Security Bearer
func (e AssetOutboundOrder) GetPage(c *gin.Context) {
	req := dto.AssetOutboundOrderGetPageReq{}
	s := service.AssetOutboundOrder{}
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
	list := make([]models.AssetOutboundOrder, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取AssetOutboundOrder失败，\r\n失败信息 %s", err.Error()))
		return
	}
	idS := make([]int, 0)
	for _, v := range list {
		idS = append(idS, v.Id)
	}
	var asset []models.AdditionsWarehousing
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("out_id in ?", idS).Find(&asset)
	assetMap := make(map[int64][]models.AdditionsWarehousing)
	for _, v := range asset {

		assetList, ok := assetMap[v.OutId]
		if !ok {
			assetList = make([]models.AdditionsWarehousing, 0)
		}
		assetList = append(assetList, v)
		assetMap[v.OutId] = assetList
	}
	result := make([]interface{}, 0)
	for _, v := range list {
		assetList, ok := assetMap[int64(v.Id)]

		if ok {
			v.Asset = assetList
		} else {
			v.Asset = make([]models.AdditionsWarehousing, 0)
		}
		result = append(result, v)
	}

	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取AssetOutboundOrder
// @Summary 获取AssetOutboundOrder
// @Description 获取AssetOutboundOrder
// @Tags AssetOutboundOrder
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetOutboundOrder} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-outbound-order/{id} [get]
// @Security Bearer
func (e AssetOutboundOrder) Get(c *gin.Context) {
	req := dto.AssetOutboundOrderGetReq{}
	s := service.AssetOutboundOrder{}
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
	var object models.AssetOutboundOrder

	e.Orm.Model(models.AssetOutboundOrder{}).Where("id = ?", orderId).Limit(1).Find(&object)
	if object.Id == 0 {
		e.Error(500, nil, "出库单不存在")
		return
	}

	var count int64
	var bindAsset []models.AdditionsWarehousing
	e.Orm.Model(models.AdditionsWarehousing{}).Where("out_id = ?", orderId).Scopes(
		cDto.Paginate(pageSizeInt, pageIndexInt),
	).Find(&bindAsset).Limit(-1).Offset(-1).Count(&count)
	object.Asset = bindAsset

	//客户信息
	var CustomModel models2.Custom

	e.Orm.Model(&CustomModel).Where("id = ?", object.CustomId).Limit(1).Find(&CustomModel)
	object.CustomInfo = CustomModel
	//位置信息
	var IdcModel models2.Idc
	e.Orm.Model(&IdcModel).Where("id = ?", object.IdcId).Limit(1).Find(&IdcModel)
	object.RegionInfo = IdcModel
	//联系人

	var UserModel models2.CustomUser
	e.Orm.Model(&UserModel).Where("id = ?", object.UserId).Limit(1).Find(&UserModel)
	object.UserInfo = UserModel
	e.PageOK(object, int(count), pageIndexInt, pageSizeInt, "")
	return
}

// Insert 创建AssetOutboundOrder
// @Summary 创建AssetOutboundOrder
// @Description 创建AssetOutboundOrder
// @Tags AssetOutboundOrder
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetOutboundOrderInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-outbound-order [post]
// @Security Bearer
func (e AssetOutboundOrder) Insert(c *gin.Context) {
	req := dto.AssetOutboundOrderInsertReq{}
	s := service.AssetOutboundOrder{}
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
		e.Error(500, err, fmt.Sprintf("创建AssetOutboundOrder失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改AssetOutboundOrder
// @Summary 修改AssetOutboundOrder
// @Description 修改AssetOutboundOrder
// @Tags AssetOutboundOrder
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetOutboundOrderUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-outbound-order/{id} [put]
// @Security Bearer
func (e AssetOutboundOrder) Update(c *gin.Context) {
	req := dto.AssetOutboundOrderUpdateReq{}
	s := service.AssetOutboundOrder{}
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
		e.Error(500, err, fmt.Sprintf("修改AssetOutboundOrder失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除AssetOutboundOrder
// @Summary 删除AssetOutboundOrder
// @Description 删除AssetOutboundOrder
// @Tags AssetOutboundOrder
// @Param data body dto.AssetOutboundOrderDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-outbound-order [delete]
// @Security Bearer
func (e AssetOutboundOrder) Delete(c *gin.Context) {
	s := service.AssetOutboundOrder{}
	req := dto.AssetOutboundOrderDeleteReq{}
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
	for _, rowId := range req.Ids {
		var model models.AssetOutboundOrder
		e.Orm.Model(&models.AssetOutboundOrder{}).Where("id = ?", rowId).First(&model)
		if model.Id == 0 {
			continue
		}
		//更新对应出库的资产 状态为在库 和 out_id = 0
		comIds := model.CombinationId
		if comIds != "" {
			e.Orm.Model(&models.Combination{}).Where("id in ?", strings.Split(comIds, ",")).Updates(map[string]interface{}{
				"status": 1,
			})
		}
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("out_id = ?", model.Id).Updates(map[string]interface{}{
			"status": 1,
			"out_id": 0,
		})
		e.Orm.Model(&model).Unscoped().Delete(&model)
	}

	e.OK(req.GetId(), "删除成功")
}
