package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/cmdb/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckCMDB = append(routerCheckCMDB, registerRsHostSwitchLogRouter)
}

// registerRsHostSwitchLogRouter
func registerRsHostSwitchLogRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.RsHostSwitchLog{}
	r := v1.Group("/rs-host-switch-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		//r.POST("", api.Insert)
		//r.PUT("/:id", actions.PermissionAction(), api.Update)
		//r.DELETE("", api.Delete)
	}
}
