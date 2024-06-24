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

	dashboardController := v1.NewDashboardController()
	apiV1Group.GET("/monitor/total", base.HandleRouteFunc(dashboardController.Total))

	// monitor_groups
	apiV1Group.GET("/monitor/groups", base.HandleRouteFunc(dashboardController.MonitorGroups))

	//--------------------服务器--------------------
	serverController := v1.NewServerController()
	apiV1Group.GET("/server/info/:id", base.HandleRouteFunc(serverController.GetServerInfo))
	apiV1Group.GET("/server/statistics/:id", base.HandleRouteFunc(serverController.GetServerStatisticsList))
}
