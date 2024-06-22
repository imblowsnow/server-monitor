package controller

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/controller/adminApi"
	"server-monitor/pkg/server/controller/api"
)

func InitRoute(r *gin.Engine) {
	// 后台接口
	adminApi.InitRoute(r)

	// 前台接口
	api.InitRoute(r)
}
