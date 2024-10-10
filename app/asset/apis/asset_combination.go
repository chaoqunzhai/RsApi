package apis

import (
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/utils"
	"go-admin/global"
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
		row.CombinationSN = row.Sn
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
		bindPriceMap[row.CombinationId] = utils.RoundDecimal(bindPrice)

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
			//有被关联到CMDB 那就是CMDB的IDC位置
			row.RegionInfo = hostBindIdcData[row.HostId]
		} else {
			//库房
			var assetRow models.AdditionsWarehousing
			e.Orm.Model(&models.AdditionsWarehousing{}).Where("w_id = ? and category_id = 1", row.Id).Limit(1).Find(&assetRow)

			if assetRow.Id > 0 {
				var storeRoomRow models.AssetWarehouse
				e.Orm.Model(&models.AssetWarehouse{}).Where("id = ?", assetRow.StoreRoomId).Limit(1).Find(&storeRoomRow)
				row.RegionInfo = map[string]interface{}{
					"val": storeRoomRow.WarehouseName,
				}
			}
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
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id = ? ", req.Id).Find(&assetList)
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
	req.Sn = strings.TrimSpace(req.Sn)
	req.Sn = strings.Replace(req.Sn, "\n", "", -1)
	if len(req.Sn) == 0 {
		e.OK("", "successful")
		return
	}

	var SearchSn string
	//SN是否为一个 黑名单.如果是 用主机名做唯一性校验
	isDirty := global.BlackMap[req.Sn]
	if isDirty {
		SearchSn = req.Hostname
	} else {
		SearchSn = req.Sn
	}

	var count int64
	e.Orm.Model(&models.Combination{}).Where("code = ?", SearchSn).Count(&count)
	if count > 0 {
		e.OK("", "successful")
		return
	}
	//主机SN如果 不存在,就创建这么一个组合, 如果存在 不进行操作
	CombinationRow := models.Combination{
		Code:   SearchSn,
		Status: "3",
	}
	e.Orm.Create(&CombinationRow)
	//创建对应的服务器资产

	hostRow := models.AdditionsWarehousing{
		Code:          SearchSn,
		Sn:            req.Sn,
		CategoryId:    1,
		CombinationId: CombinationRow.Id,
		Name:          "服务器",
		Spec:          req.Spec,
		Brand:         req.Brand,
		Status:        3,
	}
	e.Orm.Create(&hostRow)
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ?", hostRow.Id).Updates(map[string]interface{}{
		"code": fmt.Sprintf("ZC%08d", hostRow.Id),
	})

	//创建对应的磁盘资产
	for _, row := range req.DiskSn {
		Status := row.Status
		if row.Status == 1 { //因为自动注册时,磁盘正常状态就是1,如果异常就是0. 因为是自动注册 在入库逻辑 那默认就是在线的
			Status = 3
		}
		if strings.HasPrefix(row.Size, "0B") {
			continue
		}
		assetRow := models.AdditionsWarehousing{
			Sn:            row.Sn,
			Code:          row.Sn,
			CategoryId:    3,
			CombinationId: CombinationRow.Id,
			Name:          row.Name,
			Spec:          row.Size,
			Status:        Status,
			UnitId:        2,
		}
		e.Orm.Create(&assetRow)
		//e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ?", assetRow.Id).Updates(map[string]interface{}{
		//	"code": fmt.Sprintf("ZC%08d", assetRow.Id),
		//})
	}
	//创建对应的内存条
	for sn, size := range req.MemorySn {
		assetRow := models.AdditionsWarehousing{
			Code:          sn,
			Sn:            sn,
			CategoryId:    2,
			CombinationId: CombinationRow.Id,
			Name:          "内存条",
			Spec:          size,
			Status:        3,
			UnitId:        2,
		}
		e.Orm.Create(&assetRow)
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ?", assetRow.Id).Updates(map[string]interface{}{
			"code": fmt.Sprintf("ZC%08d", assetRow.Id),
		})
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
	var bindCount int64
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id in ? and combination_id > 0", req.Asset).Count(&bindCount)
	if bindCount > 0 {
		e.Error(500, nil, "资产有被关联到其他组合中")
		return
	}

	//进行校验 如果主资产没有SN 是不能录入的
	hostSn := ""
	for _, assetId := range req.Asset {
		var row models.AdditionsWarehousing
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ? and category_id = 1", assetId).Limit(1).Find(&row)

		if row.Id == 0 {
			continue
		}
		if row.Sn == "" {
			e.Error(500, nil, "主资产SN不能为空")
			return
		}
		hostSn = row.Sn

	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	var uid int
	uid, err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建Combination失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.Orm.Model(&models.Combination{}).Where("id = ?", uid).Updates(map[string]interface{}{
		"code": hostSn,
	})
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id in ?", req.Asset).Updates(map[string]interface{}{
		"combination_id": uid,
	})
	var userModel models2.SysUser
	e.Orm.Model(&models2.SysUser{}).Where("user_id = ?", user.GetUserId(c)).Limit(1).Find(&userModel)
	e.Orm.Create(&models.AssetRecording{
		User:      userModel.Username,
		Type:      1,
		AssetType: 2,
		Info:      "组合入库",
		AssetId:   uid,
		CreateBy:  user.GetUserId(c),
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

	//获取更新前原来的数据

	var oldAsset []models.AdditionsWarehousing
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id = ?", req.Id).Find(&oldAsset)
	var oldList []int
	for _, v := range oldAsset {
		oldList = append(oldList, v.Id)
	}
	var uid int
	uid, err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Combination失败，\r\n失败信息 %s", err.Error()))
		return
	}

	hostSn := ""

	for _, assetId := range req.Asset {
		var row models.AdditionsWarehousing
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ? and category_id = 1", assetId).Limit(1).Find(&row)

		if row.Id == 0 {
			continue
		}
		if row.Sn == "" {
			e.Error(500, nil, "主资产SN不能为空")
			return
		}
		hostSn = row.Sn
		//newAssetInfo = append(newAssetInfo, fmt.Sprintf("新增 %v:%v", row.Name, row.Sn))
	}

	e.Orm.Model(&models.Combination{}).Where("id = ?", uid).Updates(map[string]interface{}{
		"code": hostSn,
	})
	//把旧的清空

	e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id =  ?", uid).Updates(map[string]interface{}{
		"combination_id": 0,
	})
	//关联新的
	var newAsset []models.AdditionsWarehousing
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id in ?", req.Asset).Updates(map[string]interface{}{
		"combination_id": uid,
	}).Find(&newAsset)
	var newList []int
	for _, v := range newAsset {
		newList = append(newList, v.Id)
	}
	var userModel models2.SysUser
	e.Orm.Model(&models2.SysUser{}).Where("user_id = ?", user.GetUserId(c)).Limit(1).Find(&userModel)

	added, removed := utils.FindDifferences(oldList, newList)

	info := make([]string, 0)
	for _, v := range added {
		var row models.AdditionsWarehousing
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ?", v).Limit(1).Find(&row)
		info = append(info, fmt.Sprintf("新增 %v:%v", row.Name, func() string {
			if len(row.Sn) > 0 {
				return row.Sn
			}
			return "空SN"
		}()))
	}

	for _, v := range removed {
		var row models.AdditionsWarehousing
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ?", v).Limit(1).Find(&row)
		info = append(info, fmt.Sprintf("删除 %v:%v", row.Name, func() string {
			if len(row.Sn) > 0 {
				return row.Sn
			}
			return "空SN"
		}()))
	}

	e.Orm.Create(&models.AssetRecording{
		User:      userModel.Username,
		Type:      1,
		Info:      strings.Join(info, "\n"),
		AssetType: 2,
		AssetId:   uid,
		CreateBy:  user.GetUserId(c),
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

	if len(newIds) == 0 {
		e.Error(500, nil, "在线状态不可删除")
		return
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
