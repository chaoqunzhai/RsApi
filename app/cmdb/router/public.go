package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/cmdb/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckCMDB = append(routerCheckCMDB, registerPublicApiRouter)
	routerNoCheckCMDB = append(routerNoCheckCMDB, registerNoPublicApiRouter)
}

// registerSysApiRouter
func registerPublicApiRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Public{}
	r := v1.Group("/public").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("/uploadFile", api.UploadFile)
		r.GET("/operation", api.OperationLog)
	}
}

func registerNoPublicApiRouter(v1 *gin.RouterGroup) {
	api := apis.Public{}
	r := v1.Group("/public")
	{
		r.GET("/city_tree", api.CityTree)
	}
}
