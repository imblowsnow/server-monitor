package api

import (
	"github.com/gin-gonic/gin"
	v1 "server-monitor/pkg/server/controller/api/v1"
	"server-monitor/pkg/server/controller/base"
)

func InitRoute(r *gin.Engine) {
	apiV1Group := r.Group("/api/v1")
	//--------------------配置--------------------
	configController := v1.NewConfigController()
	apiV1Group.GET("/config", base.HandleRouteFunc(configController.GetConfig))
	//--------------------配置--------------------

	//--------------------服务器--------------------

}
