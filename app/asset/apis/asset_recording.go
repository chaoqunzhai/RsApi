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

type AssetRecording struct {
	api.Api
}

// GetPage 获取AssetRecording列表
// @Summary 获取AssetRecording列表
// @Description 获取AssetRecording列表
// @Tags AssetRecording
// @Param assetId query string false "关联资产ID"
// @Param user query string false "操作人"
// @Param type query string false "操作类型"
// @Param info query string false "处理内容"
// @Param bindOrder query string false "关联单据"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AssetRecording}} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-recording [get]
// @Security Bearer
func (e AssetRecording) GetPage(c *gin.Context) {
    req := dto.AssetRecordingGetPageReq{}
    s := service.AssetRecording{}
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
	list := make([]models.AssetRecording, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取AssetRecording失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取AssetRecording
// @Summary 获取AssetRecording
// @Description 获取AssetRecording
// @Tags AssetRecording
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AssetRecording} "{"code": 200, "data": [...]}"
// @Router /api/v1/asset-recording/{id} [get]
// @Security Bearer
func (e AssetRecording) Get(c *gin.Context) {
	req := dto.AssetRecordingGetReq{}
	s := service.AssetRecording{}
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
	var object models.AssetRecording

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取AssetRecording失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建AssetRecording
// @Summary 创建AssetRecording
// @Description 创建AssetRecording
// @Tags AssetRecording
// @Accept application/json
// @Product application/json
// @Param data body dto.AssetRecordingInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/asset-recording [post]
// @Security Bearer
func (e AssetRecording) Insert(c *gin.Context) {
    req := dto.AssetRecordingInsertReq{}
    s := service.AssetRecording{}
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
		e.Error(500, err, fmt.Sprintf("创建AssetRecording失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改AssetRecording
// @Summary 修改AssetRecording
// @Description 修改AssetRecording
// @Tags AssetRecording
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AssetRecordingUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/asset-recording/{id} [put]
// @Security Bearer
func (e AssetRecording) Update(c *gin.Context) {
    req := dto.AssetRecordingUpdateReq{}
    s := service.AssetRecording{}
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
		e.Error(500, err, fmt.Sprintf("修改AssetRecording失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除AssetRecording
// @Summary 删除AssetRecording
// @Description 删除AssetRecording
// @Tags AssetRecording
// @Param data body dto.AssetRecordingDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/asset-recording [delete]
// @Security Bearer
func (e AssetRecording) Delete(c *gin.Context) {
    s := service.AssetRecording{}
    req := dto.AssetRecordingDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除AssetRecording失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
