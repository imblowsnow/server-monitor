package controller

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/controller/adminApi"
)

func InitRoute(r *gin.Engine) {
	adminApi.InitRoute(r)
}
