package apis

import (
	"fmt"
	"github.com/google/uuid"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/global"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
)

type RsHost struct {
	api.Api
}

// GetPage 进行业务切换
// @Summary 进行业务切换
// @Description 进行业务切换
// @Tags 主机业务切换
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsHost}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host/switch [POST]
// @Security Bearer

func (e RsHost) BindIdc(c *gin.Context) {
	req := dto.HostBindIdc{}
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

	if req.IdcId == 0 || len(req.HostIds) == 0 {

		e.Error(500, nil, "请输入IDC或者主机ID列表")
		return
	}
	var idcCount int64
	e.Orm.Model(&models.RsIdc{}).Where("id = ?", req.IdcId).Count(&idcCount)

	if idcCount == 0 {
		e.Error(500, nil, "IDC不存在")
		return
	}

	for _, host := range req.HostIds {

		e.Orm.Model(&models.RsHost{}).Where("id = ?", host).Updates(map[string]interface{}{
			"idc": req.IdcId,
		})
	}
	e.OK("", "绑定IDC成功")
	return
}

// GetPage 进行业务切换
// @Summary 进行业务切换
// @Description 进行业务切换
// @Tags 主机业务切换
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsHost}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host/switch [POST]
// @Security Bearer

func (e RsHost) Switch(c *gin.Context) {
	req := dto.BusinessSwitch{}
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
	var BusinessList []models.RsBusiness

	if len(req.HostIds) == 0 {

		e.Error(500, nil, "请选择主机")
		return
	}
	if len(req.Business) == 0 {

		e.Error(500, nil, "请选择业务")
		return
	}
	busIds := make([]int, 0)
	idBindSN := make(map[int]string, 0)
	for _, i := range req.Business {
		busIds = append(busIds, i.Id)
		idBindSN[i.Id] = i.Sn
	}
	e.Orm.Model(&models.RsBusiness{}).Select("id,name").Where("id in ?", busIds).Find(&BusinessList)
	var hostList []models.RsHost
	e.Orm.Model(&models.RsHost{}).Where("id in ?", req.HostIds).Find(&hostList)

	switchList := make([]map[string]string, 0)
	if len(hostList) == 0 {
		e.Error(500, nil, "主机不存在")
		return
	}
	if len(BusinessList) == 0 {
		e.Error(500, nil, "业务不存在")
		return
	}

	for _, host := range hostList {

		//插入新的业务记录
		clearErr := e.Orm.Model(&host).Association("Business").Clear()

		if clearErr != nil {
			switchList = append(switchList, map[string]string{
				"host": host.HostName,
				"info": fmt.Sprintf("切换失败:%v", clearErr),
			})
			continue
		}
		host.Business = BusinessList

		e.Orm.Save(&host)

		for _, bu := range BusinessList {

			event := models2.HostSwitchLog{
				BusinessId: bu.Id,
				HostId:     host.Id,
				JobId:      uuid.New().String(),
				CreateBy:   user.GetUserId(c),
				BusinessSn: idBindSN[bu.Id],
			}
			e.Orm.Create(&event)
		}
		switchList = append(switchList, map[string]string{
			"host": host.HostName,
			"info": "切换成功",
		})

	}

	e.OK(switchList, "successful")
	return

}

// CountData
// @Summary 获取服务器数据统计
// @Description 查询在线/离线/机器总数数据
// @Tags 数据统计
// @Success 200 {object} response.Response "{"code": 200, "data": "","msg":"successful"}"
// @Router /api/v1/register/healthy [post]

func (e RsHost) Count(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var allHost int64
	e.Orm.Model(&models.RsHost{}).Count(&allHost)

	var offlineCount int64

	e.Orm.Model(&models.RsHost{}).Where("updated_at < DATE_SUB(NOW(), INTERVAL 5 MINUTE)").Count(&offlineCount)

	//离线-自建机房数量
	var ZjCount int64
	e.Orm.Model(&models.RsHost{}).Where("belong = 1 and updated_at < DATE_SUB(NOW(), INTERVAL 5 MINUTE)").Count(&ZjCount)

	var ZMCount int64
	e.Orm.Model(&models.RsHost{}).Where("belong = 2 and updated_at < DATE_SUB(NOW(), INTERVAL 5 MINUTE)").Count(&ZMCount)

	offlineMap := map[string]int64{
		"zj":  ZjCount,
		"zm":  ZMCount,
		"all": offlineCount,
	}

	//在线

	var onlineCount int64
	e.Orm.Model(&models.RsHost{}).Where(" updated_at >= DATE_SUB(NOW(), INTERVAL 5 MINUTE)").Count(&onlineCount)

	var ZjLineCount int64
	e.Orm.Model(&models.RsHost{}).Where("belong = 1 and updated_at >= DATE_SUB(NOW(), INTERVAL 5 MINUTE)").Count(&ZjLineCount)

	var ZMLineCount int64
	e.Orm.Model(&models.RsHost{}).Where("belong = 2 and updated_at >= DATE_SUB(NOW(), INTERVAL 5 MINUTE)").Count(&ZMLineCount)
	onlineMap := map[string]int64{
		"zj":  ZjLineCount,
		"zm":  ZMLineCount,
		"all": onlineCount,
	}
	result := map[string]interface{}{
		"all":     allHost,
		"offline": offlineMap,
		"online":  onlineMap,
	}

	e.OK(result, "successful")
	return
}

// GetPage 获取RsHost列表
// @Summary 获取RsHost列表
// @Description 获取RsHost列表
// @Tags RsHost
// @Param enable query string false "开关"
// @Param hostName query string false "主机名"
// @Param sn query string false "sn"
// @Param ip query string false "ip"
// @Param kernel query string false "内核版本"
// @Param belong query string false "机器归属"
// @Param remark query string false "备注"
// @Param operator query string false "运营商"
// @Param status query string false "主机状态"
// @Param businessSn query string false "业务SN"
// @Param province query string false "省份"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RsHost}} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host [get]
// @Security Bearer
func (e RsHost) GetPage(c *gin.Context) {
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

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}
	result := make([]map[string]interface{}, 0)
	nowTime := time.Now()
	for _, row := range list {
		customRow := make(map[string]interface{}, 1)
		customRow["updatedAt"] = fmt.Sprintf("%v", row.UpdatedAt.Format(time.DateTime))
		customRow["status"] = global.HostSuccess

		if row.HealthyAt.Valid {
			if int(nowTime.Sub(row.HealthyAt.Time).Minutes()) > 5 {
				customRow["status"] = global.HostOffline
				if row.Status != global.HostOffline {
					e.Orm.Model(&models.RsHost{}).Where("id = ?", row.Id).Updates(map[string]interface{}{
						"status": global.HostOffline,
					})
				}
			}
		}

		customRow["hostname"] = row.HostName
		var businessSnList []models2.HostSoftware
		e.Orm.Model(&models2.HostSoftware{}).Where("host_id = ? ", row.Id).Find(&businessSnList)

		snList := make([]dto.LabelRow, 0)

		snList = append(snList, dto.LabelRow{
			Label: "主机SN",
			Value: row.Sn,
		})
		for _, item := range businessSnList {
			if strings.HasPrefix(item.Key, "sn_") {
				Label := ""
				switch item.Key {
				case "sn_baishan":
					Label = "白山"
				case "sn_jinshan":
					Label = "金山"

				}
				snList = append(snList, dto.LabelRow{
					Label: Label,
					Value: item.Value,
				})
			}
		}
		customRow["sn"] = snList
		customRow["system"] = map[string]interface{}{
			"cpu": row.Cpu,
			"ip":  row.Ip,
			"memory": func() int {
				if row.Memory == 0 {
					return 0
				}

				return int(row.Memory / 1024 / 1024 / 1024)
			}(),
			"disk":   row.Disk,
			"kernel": row.Kernel,
		}
		if row.HealthyAt.Valid {
			customRow["healthyAt"] = row.HealthyAt.Time.Format("2006-01-02 15:04:05")
		}

		customRow["id"] = row.Id
		customRow["transProd"] = row.TransProvince
		customRow["isp"] = row.Isp
		customRow["balance"] = fmt.Sprintf("%vGbps", row.Balance)
		customRow["remark"] = row.Remark
		customRow["belong"] = row.Belong
		customRow["networkType"] = row.NetworkType
		customRow["monitor"] = s.GetMonitorData(row)
		customRow["idc"] = s.GetIdcInfo(row)
		customRow["line_type"] = row.LineType
		customRow["region"] = row.Region
		customRow["business"] = s.GetBusiness(row)
		customRow["address"] = row.Address
		result = append(result, customRow)
	}

	e.PageOK(result, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取RsHost
// @Summary 获取RsHost
// @Description 获取RsHost
// @Tags RsHost
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RsHost} "{"code": 200, "data": [...]}"
// @Router /api/v1/rs-host/{id} [get]
// @Security Bearer
func (e RsHost) Get(c *gin.Context) {
	req := dto.RsHostGetReq{}
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
	var object models.RsHost

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建RsHost
// @Summary 创建RsHost
// @Description 创建RsHost
// @Tags RsHost
// @Accept application/json
// @Product application/json
// @Param data body dto.RsHostInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/rs-host [post]
// @Security Bearer
func (e RsHost) Insert(c *gin.Context) {
	req := dto.RsHostInsertReq{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改RsHost
// @Summary 修改RsHost
// @Description 修改RsHost
// @Tags RsHost
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RsHostUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/rs-host/{id} [put]
// @Security Bearer
func (e RsHost) Update(c *gin.Context) {
	req := dto.RsHostUpdateReq{}
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
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除RsHost
// @Summary 删除RsHost
// @Description 删除RsHost
// @Tags RsHost
// @Param data body dto.RsHostDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/rs-host [delete]
// @Security Bearer
func (e RsHost) Delete(c *gin.Context) {
	s := service.RsHost{}
	req := dto.RsHostDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
