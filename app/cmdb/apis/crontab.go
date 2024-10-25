package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/jobs/watch"
	"go-admin/costAlg"
	"time"
)

type Crontab struct {
	api.Api
}

func (e Crontab) WatchOnlineUsage(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	startTime := time.Now()
	watch.WatchOnlineUsage()
	result := map[string]interface{}{
		"runTime": time.Now().Sub(startTime).Seconds(),
	}
	e.OK(result, "successful")
	return
}
func (e Crontab) Algorithm(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
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

	result := map[string]interface{}{
		"runTime": time.Now().Sub(startTime).Seconds(),
	}
	e.OK(result, "successful")
	return
}
