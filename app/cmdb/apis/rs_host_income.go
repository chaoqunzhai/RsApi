package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"time"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
)

type RsHostIncome struct {
	api.Api
}



func (e RsHostIncome) Compute(c *gin.Context) {
	req := dto.RsHostGetPageReq{}
	s := service.RsHost{}
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
	list := make([]models.RsHost, 0)
	var count int64
	now := time.Now()

	// 获取当前日期
	currentDate := now.Format("2006-01-02")
	fmt.Println("当前日期:", currentDate)
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}
	result:=map[string]interface{}{
		"month":currentDate,
		"day":now.Day(),
	}
	//实际上是 对主机列表进行分页。 分页的主机 然后在这个列表内查询数据
	//获取当月
	for _,row:=range result{

		fmt.Println("row~",row)
	}


	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// GetPage 获取RsHostIncome列表
// @Summary 获取RsHostIncome列表
// @Description 获取RsHostIncome列表
// @Tags RsHostIncome
// @Param hostId query string false "主机ID"
// @Param isp query string false "运营商ID"
// @Param idcId query string false "IDC ID"
// @Param buId query string false "业务ID"
// @Param usage query string false ""
// @Param settleStatus query string false ""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsHostIncome}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host-income [get]
// @Security Bearer
func (e RsHostIncome) GetPage(c *gin.Context) {
	req := dto.RsHostIncomeGetPageReq{}
	s := service.RsHostIncome{}
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
	list := make([]models.RsHostIncome, 0)
	var count int64

	var buList []models.RsBusiness
	e.Orm.Model(&models.RsBusiness{}).Select("id,name").Find(&buList)

	buMap := make(map[int]string)
	for _, b := range buList {
		buMap[b.Id] = b.Name
	}
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHostIncome失败，\r\n失败信息 %s", err.Error()))
		return
	}

	result := make([]interface{}, 0)

	for _, row := range list {

		row.BuName = buMap[row.BuId]
		result = append(result, row)
	}
	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsHostIncome
// @Summary 获取RsHostIncome
// @Description 获取RsHostIncome
// @Tags RsHostIncome
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsHostIncome} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host-income/{id} [get]
// @Security Bearer
func (e RsHostIncome) Get(c *gin.Context) {
	req := dto.RsHostIncomeGetReq{}
	s := service.RsHostIncome{}
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
	var object models.RsHostIncome

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHostIncome失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsHostIncome
// @Summary 创建RsHostIncome
// @Description 创建RsHostIncome
// @Tags RsHostIncome
// @Accept application/json
// @Product application/json
// @Param data body dto.RsHostIncomeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-host-income [post]
// @Security Bearer
func (e RsHostIncome) Insert(c *gin.Context) {
	req := dto.RsHostIncomeInsertReq{}
	s := service.RsHostIncome{}
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
		e.Error(500, err, fmt.Sprintf("创建RsHostIncome失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsHostIncome
// @Summary 修改RsHostIncome
// @Description 修改RsHostIncome
// @Tags RsHostIncome
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsHostIncomeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-host-income/{id} [put]
// @Security Bearer
func (e RsHostIncome) Update(c *gin.Context) {
	req := dto.RsHostIncomeUpdateReq{}
	s := service.RsHostIncome{}
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
		e.Error(500, err, fmt.Sprintf("修改RsHostIncome失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsHostIncome
// @Summary 删除RsHostIncome
// @Description 删除RsHostIncome
// @Tags RsHostIncome
// @Param data body dto.RsHostIncomeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-host-income [delete]
// @Security Bearer
func (e RsHostIncome) Delete(c *gin.Context) {
	s := service.RsHostIncome{}
	req := dto.RsHostIncomeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsHostIncome失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
