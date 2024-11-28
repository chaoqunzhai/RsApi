package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/grafana/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, CardRouter)
}

// registerRsHostRouter
func CardRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Card{}
	r := v1.Group("/grafana").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.CardList)

		r.GET("monitor", actions.PermissionAction(), api.Monitor)

		r.GET("isp", actions.PermissionAction(), api.IspMonitor)

		r.GET("plan-bandwidth", actions.PermissionAction(), api.PlanBandWidth)

	}

}
