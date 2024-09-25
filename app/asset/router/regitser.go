package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/asset/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerHostApiRouter)
}

func registerHostApiRouter(v1 *gin.RouterGroup) {
	api := apis.Combination{}
	r := v1.Group("/asset/register")
	{
		r.POST("", api.AutoInsert)
	}
}
