package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/cmdb/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckCMDB = append(routerCheckCMDB, registerRsCustomRouter)
}

// registerRsCustomRouter
func registerRsCustomRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.RsCustom{}
	r := v1.Group("/rs-custom").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		r.POST("/integration",actions.PermissionAction(), api.Integration)
		r.PUT("/integration/:id",actions.PermissionAction(), api.UpdateIntegration)
	}
}
