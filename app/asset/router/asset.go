package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/asset/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerAdditionsWarehousingRouter)
}

// registerAdditionsWarehousingRouter
func registerAdditionsWarehousingRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.AdditionsWarehousing{}
	//资产
	r := v1.Group("/asset").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.DELETE("", api.Delete)
	}
	//入库

	rStore := v1.Group("/asset-store").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		rStore.GET("", actions.PermissionAction(), api.GetStorePage)
		rStore.GET("/:id", actions.PermissionAction(), api.Get)
		rStore.POST("", api.Insert)
		rStore.PUT("/:id", actions.PermissionAction(), api.UpdateStore)
		rStore.DELETE("", api.Delete)
	}

}
