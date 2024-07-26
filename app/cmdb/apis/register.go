/*
*
@Author: chaoqun
* @Date: 2024/7/25 22:35
*/
package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	models2 "go-admin/cmd/migrate/migration/models"
	"net/http"
	"strings"
	"time"
)

type RegisterApi struct {
	api.Api
}

// GetPage
// @Summary 主机存活注册
// @Description 主动上报
// @Tags 主机上报
// @Header RsRole
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
	hostInstance.Status = 1
	hostInstance.Memory = req.Memory
	hostInstance.Remark = req.Remark
	hostInstance.City = req.City
	hostInstance.Isp = req.Isp
	hostInstance.NetDevice = req.NetDevice
	hostInstance.UpdatedAt = time.Now()
	e.Orm.Save(&hostInstance)

	var hostSystem models2.HostSystem

	e.Orm.Model(&models2.HostSystem{}).Where("id = ?", hostInstance.Id).First(&hostSystem)
	MemoryData := func() string {

		dat, _ := json.Marshal(req.MemoryMap)

		return string(dat)
	}()
	if req.Balance > 0 {
		hostSystem.Balance = req.Balance
	}

	hostSystem.TransmitNumber = req.TransmitNumber
	hostSystem.ReceiveNumber = req.ReceiveNumber
	hostSystem.HostId = hostInstance.Id
	hostSystem.MemoryData = MemoryData
	e.Orm.Save(&hostSystem)

	for _, softWare := range req.ExtendMap {
		var (
			softRow models2.HostSoftware
		)
		e.Orm.Model(&models2.HostSoftware{}).Where("host_id = ? and key = ?",
			hostInstance.Id, softWare.Key).First(&softRow)

		softRow.Id = hostInstance.Id
		softRow.Value = softWare.Value
		softRow.Desc = softWare.Desc
		e.Orm.Save(&softRow)

	}
	e.OK("", "successful")
	return
}
