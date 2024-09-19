package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/asset/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCombinationRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerCombinationNoAuthRouter)
}

// registerCombinationRouter
func registerCombinationRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Combination{}
	r := v1.Group("/asset-combination").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}
func registerCombinationNoAuthRouter(v1 *gin.RouterGroup) {
	api := apis.Combination{}
	r := v1.Group("/asset-combination")
	{
		r.POST("/auto", api.AutoInsert)
	}
}
