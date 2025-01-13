package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/cmdb/apis"
)

func init() {
	routerNoCheckCMDB = append(routerNoCheckCMDB, registerCrontabRouter)
}

// registerRsBusinessRouter
func registerCrontabRouter(v1 *gin.RouterGroup) {
	api := apis.Crontab{}
	r := v1.Group("/crontab")
	{
		r.GET("algorithm", api.Algorithm)
		r.GET("watchOnlineUsage", api.WatchOnlineUsage)
		//Data burning 数据刻录,保存一些每天在变化的数据
		r.GET("dataBurning",api.DataBurning)
	}
}
