package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/cmdb/service"
	"go-admin/app/cmdb/service/dto"
	"go-admin/costAlg"
	"time"
)

type Crontab struct {
	api.Api
}

func (e Crontab) Algorithm(c *gin.Context) {
	req := dto.RsContractGetPageReq{}
	s := service.RsContract{}
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

	startTime := time.Now()
	costAlgorithm := costAlg.CostAlgorithm{}
	costAlgorithm.SetupDb(sdk.Runtime.GetDb())
	costAlgorithm.StartHostCompute()

	endTime := time.Now()
	result := map[string]interface{}{
		"runTime": endTime.Sub(startTime),
	}
	e.OK(result, "successful")
	return
}
