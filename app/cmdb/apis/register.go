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
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	models2 "go-admin/cmd/migrate/migration/models"
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
	hostInstance.Disk = req.Disk
	hostInstance.Kernel = req.Kernel
	hostInstance.Status = global.HostSuccess
	hostInstance.Memory = req.Memory
	hostInstance.Remark = req.Remark

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
	hostInstance.Isp = ispNumber
	hostInstance.NetDevice = req.NetDevice
	hostInstance.HealthyAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	e.Orm.Save(&hostInstance)

	var hostSystem models2.HostSystem

	e.Orm.Model(&models2.HostSystem{}).Where("host_id = ?", hostInstance.Id).First(&hostSystem)
	MemoryData := func() string {

		dat, _ := json.Marshal(req.MemoryMap)

		return string(dat)
	}()

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
