/**
@Author: chaoqun
* @Date: 2024/7/25 22:25
*/
package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/cmdb/apis"
)

func init() {
	routerNoCheckCMDB = append(routerNoCheckCMDB, registerHostApiRouter)
}

func registerHostApiRouter(v1 *gin.RouterGroup) {
	api := apis.RegisterApi{}
	r := v1.Group("/register")
	{
		r.POST("/healthy", api.Healthy)
	}
}