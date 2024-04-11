package controller

import (
	"github.com/gin-gonic/gin"
	v1 "server/controller/api/v1"
)

func InitRoute(r *gin.Engine) {
	v1Group := r.Group("/api/v1")
	v1Group.GET("/server", v1.GetServerInfo)
	v1Group.GET("/server/status", v1.GetServerStatus)

	adminGroup := r.Group("/api/admin")
	adminGroup.GET("/server", v1.GetServerInfo)
}
