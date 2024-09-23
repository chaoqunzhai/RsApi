package apis

import (
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/asset/models"
	"go-admin/app/asset/service"
	"go-admin/app/asset/service/dto"
	"go-admin/common/actions"
)

type Combination struct {
	api.Api
}

func (e Combination) BindCustomMap(ids []int) map[int]interface{} {
	dat := make(map[int]interface{}, 0)
	if len(ids) <= 0 {
		return dat
	}

	return dat

}
func (e Combination) HostBindIdcData(ids []int) map[int]interface{} {

	//获取关联的主机列表
	dat := make(map[int]interface{}, 0)
	if len(ids) <= 0 {
		return dat
	}
	var hostList []models2.Host
	idsList := make([]int, 0)
	//查询主机列表
	e.Orm.Model(&models2.Host{}).Select("id,idc").Where("id in ?", ids).Find(&hostList)
	for _, row := range hostList {
		idsList = append(idsList, row.Idc)
	}
	idsList = utils.RemoveRepeatInt(idsList)
	var RowList []models2.Idc
	e.Orm.Model(&models2.Idc{}).Where("id in ? ", idsList).Find(&RowList)
	IdcMapData := make(map[int]models2.Idc, 0)
	//做出idc的MAP
	for _, row := range RowList {
		IdcMapData[row.Id] = row
	}
	//把IDC数据放到对应的主机上

	for _, row := range hostList {
		if idcInfo, ok := IdcMapData[row.Idc]; ok {
			dat[row.Id] = map[string]interface{}{
				"val": idcInfo.Name,
			}
		}
	}

	return dat

}
func (e Combination) StoreRoomMapInfo(ids []int) map[int]interface{} {

	dat := make(map[int]interface{}, 0)

	if len(ids) <= 0 {
		return dat
	}
	var StoreRoomLists []models.AssetWarehouse
	e.Orm.Model(&models.AssetWarehouse{}).Where("id in ?", ids).Find(&StoreRoomLists)
	for _, row := range StoreRoomLists {
		dat[row.Id] = map[string]interface{}{
			"val": row.WarehouseName,
		}
	}

	return dat
}

// GetPage 获取Combination列表
// @Summary 获取Combination列表
// @Description 获取Combination列表
// @Tags Combination
// @Param jobId query string false "组合编号"
// @Param status query string false "资产状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Combination}} "{"code": 200, "data": [...]}"
// @Router /api/v1/combination [get]
// @Security Bearer
func (e Combination) GetPage(c *gin.Context) {
	req := dto.CombinationGetPageReq{}
	s := service.Combination{}
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
	list := make([]models.Combination, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Combination失败，\r\n失败信息 %s", err.Error()))
		return
	}
	bindIds := make([]int, 0)

	bindHostIds := make([]int, 0)
	for _, row := range list {
		if row.HostId > 0 {
			bindHostIds = append(bindHostIds, row.HostId)
		}
		bindIds = append(bindIds, row.Id)
	}

	var assetList []models.AdditionsWarehousing
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id in ?", bindIds).Find(&assetList)

	bindMap := make(map[int]int, 0)
	bindAssetMap := make(map[int][]models.AdditionsWarehousing, 0)
	bindPriceMap := make(map[int]float64, 0)
	StoreRoomIdLists := make([]int, 0)
	for _, row := range assetList {

		assetBindList, assetOk := bindAssetMap[row.CombinationId]
		if !assetOk {
			assetBindList = make([]models.AdditionsWarehousing, 0)
		}
		assetBindList = append(assetBindList, row)
		bindAssetMap[row.CombinationId] = assetBindList
		StoreRoomIdLists = append(StoreRoomIdLists, row.StoreRoomId)
		bindCount, ok := bindMap[row.CombinationId]
		bindPrice, ok := bindPriceMap[row.CombinationId]
		if !ok {
			bindCount = 0
			bindPrice = 0
		}
		bindCount += 1
		bindPrice += row.Price
		bindPriceMap[row.CombinationId] = bindPrice
		bindMap[row.CombinationId] = bindCount
	}

	hostBindIdcData := e.HostBindIdcData(bindHostIds)

	result := make([]interface{}, 0)

	for _, row := range list {

		if AssetCount, ok := bindMap[row.Id]; ok {
			row.AssetCount = AssetCount
		}
		row.Price = bindPriceMap[row.Id]
		if row.HostId > 0 {
			//主机位置
			row.RegionInfo = hostBindIdcData[row.HostId]
		}

		if req.Extend == 2 { //展示扩展的信息,例如资产列表
			row.Asset = bindAssetMap[row.Id]
		}
		result = append(result, row)
	}
	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Combination
// @Summary 获取Combination
// @Description 获取Combination
// @Tags Combination
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Combination} "{"code": 200, "data": [...]}"
// @Router /api/v1/combination/{id} [get]
// @Security Bearer
func (e Combination) Get(c *gin.Context) {
	req := dto.CombinationGetReq{}
	s := service.Combination{}
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
	var object models.Combination

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Combination失败，\r\n失败信息 %s", err.Error()))
		return
	}

	var assetList []models.AdditionsWarehousing
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id = ?", req.Id).Find(&assetList)
	object.Asset = assetList

	if object.HostId > 0 {
		//主机位置
		hostBindIdcData := e.HostBindIdcData([]int{object.HostId})
		object.RegionInfo = hostBindIdcData[object.HostId]
	}

	object.Price = 1
	e.OK(object, "查询成功")
}

//开机后首次自动注册

func (e Combination) AutoInsert(c *gin.Context) {
	req := dto.CombinationAutoReq{}
	s := service.Combination{}
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

	e.OK("", "successful")
	return

}

// Insert 创建Combination
// @Summary 创建Combination
// @Description 创建Combination
// @Tags Combination
// @Accept application/json
// @Product application/json
// @Param data body dto.CombinationInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/combination [post]
// @Security Bearer
func (e Combination) Insert(c *gin.Context) {
	req := dto.CombinationInsertReq{}
	s := service.Combination{}
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

	var uid int
	uid, err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建Combination失败，\r\n失败信息 %s", err.Error()))
		return
	}

	Code := fmt.Sprintf("RS%08d", uid)
	e.Orm.Model(&models.Combination{}).Where("id = ?", uid).Updates(map[string]interface{}{
		"code": Code,
	})
	var bindCount int64
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id in ? and combination_id > 0", req.Asset).Count(&bindCount)
	if bindCount > 0 {
		e.Error(500, nil, "资产有被关联到其他组合中")
		return
	}
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id in ?", req.Asset).Updates(map[string]interface{}{
		"combination_id": uid,
	})
	e.OK(req.GetId(), "创建成功")
}

// Update 修改Combination
// @Summary 修改Combination
// @Description 修改Combination
// @Tags Combination
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CombinationUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/combination/{id} [put]
// @Security Bearer
func (e Combination) Update(c *gin.Context) {
	req := dto.CombinationUpdateReq{}
	s := service.Combination{}
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

	var uid int
	uid, err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Combination失败，\r\n失败信息 %s", err.Error()))
		return
	}
	//把旧的清空

	e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id =  ?", uid).Updates(map[string]interface{}{
		"combination_id": 0,
	})
	//关联新的
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id in ?", req.Asset).Updates(map[string]interface{}{
		"combination_id": uid,
	})
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Combination
// @Summary 删除Combination
// @Description 删除Combination
// @Tags Combination
// @Param data body dto.CombinationDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/combination [delete]
// @Security Bearer
func (e Combination) Delete(c *gin.Context) {
	s := service.Combination{}
	req := dto.CombinationDeleteReq{}
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

	newIds := make([]int, 0)

	for _, row := range req.Ids {
		var count int64
		//只有在库的时候 才能删除
		e.Orm.Model(&models.Combination{}).Where("id = ? and status = 1", row).Count(&count)
		if count > 0 {
			newIds = append(newIds, row)
		}
	}

	err = s.Remove(newIds, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除Combination失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id in ?", newIds).Updates(map[string]interface{}{
		"combination_id": 0,
	})
	e.OK(req.GetId(), "删除成功")
}
