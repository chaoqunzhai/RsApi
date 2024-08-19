/*
*
@Author: chaoqun
* @Date: 2024/7/25 22:35
*/
package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/dial"

	models2 "go-admin/cmd/migrate/migration/models"
	_ "go-admin/common/dial"
	models3 "go-admin/common/models"
	"go-admin/global"
	"net/http"
	"strings"
	"time"
)

type RegisterApi struct {
	api.Api
}

// RegisterData
// @Summary 主机存活注册
// @Description 主动上报
// @Tags 主机上报
// @Success 200 {object} response.Response "{"code": 200, "data": "","msg":"successful"}"
// @Router /api/v1/register/healthy [post]

func (e *RegisterApi) Healthy(c *gin.Context) {
	s := service.RegisterApi{}
	req := dto.RegisterMetrics{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	registerHeaderKey := c.GetHeader("RsRole")

	if strings.TrimSpace(registerHeaderKey) != "rs-sre" {

		e.Error(http.StatusUnauthorized, nil, "you set Header")
		return
	}

	var hostInstance models2.Host
	e.Orm.Model(&hostInstance).Where("sn = ?", req.Sn).First(&hostInstance)

	hostInstance.Sn = req.Sn
	hostInstance.HostName = req.Hostname
	hostInstance.Ip = req.Ip
	hostInstance.CPU = req.CPU
	hostInstance.Kernel = req.Kernel
	hostInstance.Status = global.HostSuccess
	hostInstance.Memory = req.Memory
	hostInstance.Remark = req.Remark

	var IdcId int
	var ispNumber int
	switch strings.TrimSpace(req.Isp) {
	case "移动":
		ispNumber = 1
	case "电信":
		ispNumber = 2
	case "联通":
		ispNumber = 3
	default:
		ispNumber = 4
	}
	if req.Balance > 0 {
		hostInstance.Balance = req.Balance
	}
	if req.BandwidthCnf.Line > 0 {
		hostInstance.AllLine = int(req.BandwidthCnf.Line)
		hostInstance.LineBandwidth = req.BandwidthCnf.Width
	}
	hostInstance.Isp = ispNumber

	if hostInstance.Idc == 0 { //防止已经关联了IDC,被其他原因冲掉
		hostInstance.Idc = IdcId
	}
	hostInstance.HealthyAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	e.Orm.Save(&hostInstance)

	NetDeviceMap := make(map[string]int, 0)
	if req.NetDevice != "" {

		NetDeviceList := strings.Split(req.NetDevice, ",")

		for _, NetDeviceName := range NetDeviceList {
			var (
				DeviceRow models2.HostNetDevice
			)
			e.Orm.Model(&models2.HostNetDevice{}).Where("host_id = ? and `name` = ?",
				hostInstance.Id, NetDeviceName).First(&DeviceRow)
			DeviceRow.HostId = hostInstance.Id
			DeviceRow.Name = NetDeviceName
			DeviceRow.UpdatedAt = models3.XTime{
				Time: time.Now(),
			}
			DeviceRow.Status = 1
			e.Orm.Save(&DeviceRow)
			NetDeviceMap[NetDeviceName] = DeviceRow.Id
		}

	}

	//拨号列表,
	for _, DialRow := range req.Dial {

		if setEvents, ok := dial.MapCnf.Get(DialRow.A); ok {
			if setEvents.Idc > 0 {
				IdcId = setEvents.Idc
			}
		}
		//如果列表存在 全局的缓存中 那就自动归到idc下
		var (
			DialRowModel models.RsDial
			DialCount    int64
		)
		var bindNetDeviceId int
		NetDeviceId, NetDeviceOk := NetDeviceMap[DialRow.I]
		if NetDeviceOk {
			bindNetDeviceId = NetDeviceId
		}
		//对于自动上报数据的数据,做一个特定创建,防止 已经创建了这个账号，被自动创建也冲掉
		e.Orm.Model(&models.RsDial{}).Where("account = ? and source = 1", DialRow.A).Count(&DialCount)
		if DialCount > 0 {
			e.Orm.Model(&models.RsDial{}).Where("account = ? and source = 1", DialRow.A).Updates(map[string]interface{}{
				"host_id":     hostInstance.Id,
				"idc_id":      hostInstance.Idc,
				"account":     DialRow.A,
				"pass":        DialRow.P,
				"status":      DialRow.S,
				"ip":          DialRow.Ip,
				"mac":         DialRow.Mac,
				"source":      1,
				"device_name": DialRow.I,
				"dial_name":   DialRow.D,
				"bu":          DialRow.BU,
				"device_id":   bindNetDeviceId,
			})
		} else {
			DialRowModel.Bu = DialRow.BU
			DialRowModel.HostId = hostInstance.Id
			DialRowModel.IdcId = IdcId
			DialRowModel.Account = DialRow.A
			DialRowModel.Pass = DialRow.P
			DialRowModel.Ip = DialRow.Ip
			DialRowModel.Mac = DialRow.Mac
			DialRowModel.DialName = DialRow.D
			DialRowModel.Source = 1
			DialRowModel.DeviceId = bindNetDeviceId
			DialRowModel.DeviceName = DialRow.I
			DialRowModel.Status = DialRow.S
			e.Orm.Save(&DialRowModel)
		}
	}
	var hostSystem models2.HostSystem

	e.Orm.Model(&models2.HostSystem{}).Where("host_id = ?", hostInstance.Id).First(&hostSystem)
	MemoryData := func() string {

		dat, _ := json.Marshal(req.MemoryMap)

		return string(dat)
	}()
	hostSystem.Disk = req.Disk
	hostSystem.TransmitNumber = req.TransmitNumber
	hostSystem.ReceiveNumber = req.ReceiveNumber
	hostSystem.HostId = hostInstance.Id
	hostSystem.MemoryData = MemoryData
	e.Orm.Save(&hostSystem)

	if len(req.BusinessSn) > 0 {
		for key, val := range req.BusinessSn {
			var (
				snRow models2.HostSoftware
			)
			if val == "" {
				continue
			}
			snKey := fmt.Sprintf("sn_%v", key)
			e.Orm.Model(&models2.HostSoftware{}).Where("host_id = ? and `key` = ?",
				hostInstance.Id, snKey).First(&snRow)
			snRow.HostId = hostInstance.Id
			snRow.Key = snKey
			snRow.Value = val
			e.Orm.Save(&snRow)
		}
	}
	for _, softWare := range req.ExtendMap {
		var (
			softRow models2.HostSoftware
		)
		e.Orm.Model(&models2.HostSoftware{}).Where("host_id = ? and `key`  = ?",
			hostInstance.Id, softWare.Key).First(&softRow)

		softRow.HostId = hostInstance.Id
		softRow.Key = softWare.Key
		softRow.Value = softWare.Value
		softRow.Desc = softWare.Desc
		e.Orm.Save(&softRow)

	}
	e.OK("", "successful")
	return
}
