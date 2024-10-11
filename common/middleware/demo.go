package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"net/http"
)

func DemoEvn() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if config.ApplicationConfig.Mode == "demo" {
			if method == "GET" ||
				method == "OPTIONS" ||
				c.Request.RequestURI == "/api/v1/login" ||
				c.Request.RequestURI == "/api/v1/logout" {
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
