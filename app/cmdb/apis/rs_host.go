package apis

import (
	"fmt"
	"github.com/google/uuid"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/prometheus"
	"go-admin/common/remoteCommand"
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
	var idcModel models.RsIdc
	e.Orm.Model(&models.RsIdc{}).Where("id = ?", req.IdcId).Limit(1).Find(&idcModel)

	if idcModel.Id == 0 {
		e.Error(500, nil, "IDC不存在")
		return
	}

	for _, hostId := range req.HostIds {

		e.Orm.Model(&models.RsHost{}).Where("id = ?", hostId).Updates(map[string]interface{}{
			"idc": req.IdcId,
		})
		e.Orm.Create(&models2.OperationLog{
			CreateUser: user.GetUserName(c),
			Action:     "POST",
			Module:     "rs_host",
			ObjectId:   hostId,
			TargetId:   req.IdcId,
			Info:       fmt.Sprintf("绑定 IDC: %s", idcModel.Name),
		})
	}
	e.OK("", "绑定IDC成功")
	return
}

func (e RsHost) BindDial(c *gin.Context) {
	req := dto.HostBindDial{}
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
	var dialModel models.RsDial
	e.Orm.Model(&dialModel).Where("id = ?", req.DialId).Limit(1).Find(&dialModel)
	if dialModel.Id == 0 {
		e.Error(500, nil, "拨号配置不存在")
		return
	}
	if dialModel.HostId != 0 {
		e.Error(500, nil, "拨号已经被关联,无法绑到")
		return
	}

	dialModel.DeviceId = req.DriverId
	dialModel.HostId = req.HostId

	e.Orm.Save(&dialModel)
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "POST",
		Module:     "rs_host",
		ObjectId:   dialModel.HostId,
		TargetId:   dialModel.Id,
		Info:       fmt.Sprintf("绑定拨号:%v", dialModel.Account),
	})
	e.OK("", "successful")
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
	busIds := make([]string, 0)
	for _, i := range req.Business {
		busIds = append(busIds, fmt.Sprintf("%v", i.Id))
	}
	e.Orm.Model(&models.RsBusiness{}).Where("id in ?", busIds).Find(&BusinessList)
	var hostList []models.RsHost
	e.Orm.Model(&models.RsHost{}).Where("id in ?", req.HostIds).Preload("Business").Find(&hostList)

	switchList := make([]map[string]string, 0)
	if len(hostList) == 0 {
		e.Error(500, nil, "主机不存在")
		return
	}
	if len(BusinessList) == 0 {
		e.Error(500, nil, "业务不存在")
		return
	}
	JobId := uuid.New().String()
	for _, host := range hostList {
		//业务的英文列表
		buEnList := make([]string, 0)
		//先获取原来的业务
		sureList := make([]models.RsBusiness, 0)

		for _, business := range host.Business {
			sureList = append(sureList, business)
		}

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

		addSwitchLog := func(name, enName string) {
			for _, bu := range BusinessList {
				//记录下 主机之前的sn列表,需要通过sn去查询监控数据
				event := models.RsHostSwitchLog{
					BuTargetId: bu.Id,
					HostId:     host.Id,
					JobId:      uuid.New().String(),
					CreateBy:   user.GetUserId(c),
					Desc:       req.Desc,
					BuSource:   name,
					BuEnSource: enName,
				}

				if bu.EnName == "" {
					continue
				}
				buEnList = append(buEnList, bu.EnName)
				e.Orm.Create(&event)
			}
		}
		if len(sureList) == 0 { //暂无业务的情况下
			addSwitchLog("", "")
		} else {
			for _, sure := range sureList {
				addSwitchLog(sure.Name, sure.EnName)
			}
		}
		switchList = append(switchList, map[string]string{
			"host": host.HostName,
			"info": "切换成功",
		})

		//同时执行远程shell命令

		command := remoteCommand.Command{
			Orm:        e.Orm,
			CreateBy:   user.GetUserId(c),
			HostId:     host.Id,
			RemotePort: host.RemotePort,
			JobId:      JobId,
		}
		go func() {
			command.BusinessSwitching(strings.Join(buEnList, "-"))
		}()

		e.Orm.Create(&models2.OperationLog{
			CreateUser: user.GetUserName(c),
			Action:     "POST",
			Module:     "rs_host",
			ObjectId:   host.Id,
			Info:       fmt.Sprintf("切换至业务:%v", strings.Join(buEnList, "-")),
		})

	}

	e.OK(JobId, "successful")
	return

}

func (e RsHost) Driver(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	hosId := c.Param("id")
	var hostList []models2.HostNetDevice
	e.Orm.Model(&models2.HostNetDevice{}).Where("host_id = ?", hosId).Find(&hostList)

	e.PageOK(hostList, len(hostList), 1, -1, "")
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

	e.Orm.Model(&models.RsHost{}).Where("healthy_at <= DATE_SUB(NOW(), INTERVAL 30 MINUTE) OR healthy_at IS NULL").Count(&offlineCount)

	//自建机房数量
	var ZjCount int64

	e.Orm.Model(&models.RsHost{}).Where("belong IN (0,1) and (healthy_at <= DATE_SUB(NOW(), INTERVAL 30 MINUTE) OR healthy_at IS NULL )").Count(&ZjCount)

	var ZMCount int64
	e.Orm.Model(&models.RsHost{}).Where("belong = 2 and (healthy_at <= DATE_SUB(NOW(), INTERVAL 30 MINUTE) OR healthy_at IS NULL )").Count(&ZMCount)

	offlineMap := map[string]int64{
		"zj":  ZjCount,
		"zm":  ZMCount,
		"all": offlineCount,
	}

	//在线

	var onlineCount int64
	e.Orm.Model(&models.RsHost{}).Where(" healthy_at >= DATE_SUB(NOW(), INTERVAL 30 MINUTE)").Count(&onlineCount)

	var ZjLineCount int64
	e.Orm.Model(&models.RsHost{}).Where("belong IN (0,1) and healthy_at >= DATE_SUB(NOW(), INTERVAL 30 MINUTE)").Count(&ZjLineCount)

	var ZMLineCount int64
	e.Orm.Model(&models.RsHost{}).Where("belong = 2 and healthy_at >= DATE_SUB(NOW(), INTERVAL 30 MINUTE)").Count(&ZMLineCount)
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

func (e RsHost) MonitorFlow(c *gin.Context) {
	req := dto.RsHostMonitorFlow{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//获取这个主机的主机名

	var hostInstance models.RsHost
	e.Orm.Model(&hostInstance).Select("host_name,id").Where("id = ?", req.Id).Limit(1).Find(&hostInstance)

	if hostInstance.Id == 0 {
		e.Error(500, nil, "主机不存在")
		return
	}

	if req.Start == "" {
		req.Start = fmt.Sprintf("%v", time.Now().Add(-time.Hour).Unix())
	}
	if req.End == "" {
		req.End = fmt.Sprintf("%v", time.Now().Unix())
	}
	if req.Setup == 0 {
		req.Setup = 60
	}
	HostName := hostInstance.HostName
	result := prometheus.Transmit(HostName, &req)
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

	hostIds := make([]int, 0)
	idcIds := make([]int, 0)
	for _, row := range list {

		hostIds = append(hostIds, row.Id)
		idcIds = append(idcIds, row.Idc)
	}
	HostSoftwareMap := s.GetHostSoftware(hostIds)

	HostMapMonitorData := s.GetMonitorData(hostIds)

	IdcMapData := s.GetIdcList(idcIds)

	BusinessMapData := service.GetHostBindBusinessMap(e.Orm, hostIds)
	for _, row := range list {
		customRow := make(map[string]interface{}, 1)
		customRow["updatedAt"] = fmt.Sprintf("%v", row.UpdatedAt.Format(time.DateTime))

		validStatus := row.Status
		//只做在线数据的检查
		if row.Auth > 0 { //只有主机有权限的时候去检查 + 主机是正常的时候
			if row.Status == global.HostSuccess {

				if row.HealthyAt.Valid {
					if int(nowTime.Sub(row.HealthyAt.Time).Minutes()) > 6 { //如果上报的时间大于5分钟 那就删掉线了

						e.Orm.Model(&models.RsHost{}).Where("id = ?", row.Id).Updates(map[string]interface{}{
							"status": global.HostOffline,
						})

					} else { //在5分钟内
						validStatus = global.HostSuccess
					}
					customRow["healthyAt"] = row.HealthyAt.Time.Format("2006-01-02 15:04:05")
				} else { //是一个没有注册到节点的机器，因为没有健康时间
					validStatus = global.HostOffline
				}

				if validStatus == global.HostOffline {
					e.Orm.Model(&models.RsHost{}).Where("id = ?", row.Id).Updates(map[string]interface{}{
						"status": global.HostOffline,
					})
				}

			}

		} else {
			//主机没有权限.执行ProbeShell的命令去进行异步探测
		}
		customRow["status"] = validStatus
		customRow["hostname"] = row.HostName

		snList := make([]dto.LabelRow, 0)

		snList = append(snList, dto.LabelRow{
			Label: "主机SN",
			Value: row.Sn,
		})
		if businessSnList, ok := HostSoftwareMap[row.Id]; ok {
			for _, item := range businessSnList {
				if strings.HasPrefix(item.Key, "sn_") {
					Label := ""
					switch item.Key {
					case "sn_baishan":
						Label = "白山SN:"
					case "sn_jinshan":
						Label = "金山SN:"

					}
					snList = append(snList, dto.LabelRow{
						Label: Label,
						Value: item.Value,
					})
				}
			}
		}
		customRow["auth"] = row.Auth
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
			"kernel": row.Kernel,
		}
		if row.HealthyAt.Valid {
			customRow["healthyAt"] = row.HealthyAt.Time.Format("2006-01-02 15:04:05")
		}
		customRow["remotePort"] = row.RemotePort
		customRow["ip"] = row.Ip
		customRow["publicIp"] = row.PublicIp
		customRow["id"] = row.Id
		customRow["transProd"] = row.TransProvince
		customRow["isp"] = row.Isp
		customRow["balance"] = fmt.Sprintf("%vGbps", row.Balance)
		customRow["remark"] = row.Remark

		customRow["belong"] = row.Belong
		customRow["networkType"] = row.NetworkType
		if monitorDat, ok := HostMapMonitorData[row.Id]; ok {
			customRow["monitor"] = monitorDat["memory"]
		}
		if idcInfo, ok := IdcMapData[row.Idc]; ok {
			if len(idcInfo) > 0 {
				customRow["idcInfo"] = idcInfo[0]
			}
		}
		customRow["lineType"] = row.LineType
		customRow["region"] = row.Region

		if BusinessDat, ok := BusinessMapData[row.Id]; ok {
			customRow["business"] = BusinessDat
		}
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
	var count int64
	e.Orm.Model(&models.RsHost{}).Where("hostname = ?", req.HostName).Count(&count)
	if count > 0 {
		e.Error(500, nil, "主机名已存在")
		return
	}
	modelId, err := s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建RsHost失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "POST",
		Module:     "rs_host",
		ObjectId:   modelId,
		TargetId:   modelId,
		Info:       "创建主机信息",
	})
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
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "PUT",
		Module:     "rs_host",
		ObjectId:   req.Id,
		TargetId:   req.Id,
		Info:       "更新主机信息",
	})
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
	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "DELETE",
		Module:     "rs_host",
		ObjectId:   req.Ids[0],
		TargetId:   req.Ids[0],
		Info:       "删除主机信息",
	})
	e.OK(req.GetId(), "删除成功")
}
