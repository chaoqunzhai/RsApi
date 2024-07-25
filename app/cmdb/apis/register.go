/**
@Author: chaoqun
* @Date: 2024/7/25 22:35
*/
package apis

import (
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
//
func (e *RegisterApi) Healthy(c *gin.Context) {
	s := service.RegisterApi{}
	req:=dto.RegisterHost{}
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
	registerHeaderKey :=c.GetHeader("RsRole")

	if strings.TrimSpace(registerHeaderKey) != "rs-sre"{

		e.Error(http.StatusUnauthorized,nil,"you set Header")
		return
	}
	if req.Sn == ""{
		e.Error(-1,nil,"sn uniqueness")
		return
	}

	var hostInstance models2.Host
	e.Orm.Model(&hostInstance).Where("sn = ?",req.Sn).First(&hostInstance)

	if hostInstance.Id == 0 {
		e.Error(-1,nil,"not sn")
		return
	}

	hostInstance.Sn = req.Sn
	hostInstance.HostName = req.HostName
	hostInstance.Ip = req.Ip
	hostInstance.CPU = req.CPU
	hostInstance.Disk = req.Disk
	hostInstance.Memory = req.Memory
	hostInstance.Remark = req.Remark
	hostInstance.UpdatedAt = time.Now()
	e.Orm.Save(&hostInstance)

	for _,softWare:=range req.Software{
		var (
			softRow []models2.HostSoftware
		)
		e.Orm.Model(&models2.HostSoftware{}).Where("host_id = ? and key = ?",
			hostInstance.Id, softWare.Key).Find(&softRow)

		if len(softRow) == 0 {
			softErr:=e.Orm.Create(&models2.HostSoftware{
				HostId: hostInstance.Id,
				Key: softWare.Key,
				Value: softWare.Value,
				Desc: softWare.Desc,
			})
			if softErr!=nil{
				e.Error(http.StatusBadRequest,softErr.Error,"error")
				return
			}
		}else {
			firstRow :=softRow[0]
			firstRow.Value = softWare.Value
			firstRow.Desc = softWare.Desc
			e.Orm.Save(&firstRow)
		}

	}
	e.OK("","successful")
	return
}
