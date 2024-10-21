package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/cmdb/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckCMDB = append(routerCheckCMDB, registerRsHostRouter)
}

// registerRsHostRouter
func registerRsHostRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.RsHost{}
	r := v1.Group("/rs-host").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)

		r.GET("/monitor_flow/:id", api.MonitorFlow)
		r.GET("/countOnline", actions.PermissionAction(), api.CountOnline)
		r.GET("/countOffline", actions.PermissionAction(), api.CountOffline)
		r.GET("/driver/:id", actions.PermissionAction(), api.Driver)
		r.POST("/switch", actions.PermissionAction(), api.Switch)
		r.POST("/bindIdc", actions.PermissionAction(), api.BindIdc)
		r.POST("/bindDial", actions.PermissionAction(), api.BindDial)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}

	//远程命令执行
	{

		//主机名更改-仅支持单机
		r.POST("/exec/upHostName", actions.PermissionAction(), api.ExecUpHostName)
		//命令执行 可批量操作
		r.POST("/exec/command", actions.PermissionAction(), api.ExecCommand)
		//重启主机 可批量操作
		r.POST("/exec/reboot", actions.PermissionAction(), api.ExecReboot)

		//获取单个任务信息
		r.GET("/exec/log/:jobId", api.GetJobLog)
		//获取主机的所有执行日志
		r.GET("/exec/hostLog/:id", api.GetHostLog)
	}
}
