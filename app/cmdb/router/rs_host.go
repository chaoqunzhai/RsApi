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
		r.GET("/count", actions.PermissionAction(), api.Count)
		r.POST("/switch", actions.PermissionAction(), api.Switch)
		r.POST("/bindIdc", actions.PermissionAction(), api.BindIdc)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}
