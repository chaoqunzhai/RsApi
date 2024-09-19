package apis

import (
	"fmt"

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
	for _, row := range list {
		bindIds = append(bindIds, row.Id)
	}
	var assetList []models.AdditionsWarehousing
	e.Orm.Model(&models.AdditionsWarehousing{}).Select("id,combination_id").Where("combination_id in ?", bindIds).Find(&assetList)

	bindMap := make(map[int]int, 0)
	for _, row := range assetList {
		bindCount, ok := bindMap[row.CombinationId]
		if !ok {
			bindCount = 0
		}
		bindCount += 1
		bindMap[row.CombinationId] = bindCount
	}

	result := make([]interface{}, 0)

	for _, row := range list {

		if AssetCount, ok := bindMap[row.Id]; ok {
			row.AssetCount = AssetCount
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
	e.OK(object, "查询成功")
}

//开机后首次自动注册

func (e Combination) AutoInsert(c *gin.Context) {
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
		e.Orm.Model(&models.Combination{}).Where("id = ? and status = 1", row).Count(&count)
		if count > 0 {
			continue
		}
		newIds = append(newIds, row)
	}
	if len(newIds) == 0 {
		e.Error(500, nil, "所选组合不可删除")
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
