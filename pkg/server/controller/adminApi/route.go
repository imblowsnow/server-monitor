package adminApi

import (
	"github.com/gin-gonic/gin"
	v1 "server-monitor/pkg/server/controller/adminApi/v1"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/controller/middle"
	"server-monitor/pkg/server/dal/do"
)

func InitRoute(r *gin.Engine) {
	adminApiV1Group := r.Group("/admin-api/v1")
	adminApiV1Group.Use(middle.AdminAuthMiddle)

	// ----------------- 用户相关 ---------------------
	userController := v1.NewUserController()
	adminApiV1Group.POST("/user/login", base.HandleRouteFunc(userController.Login))

	// -------------   配置相关 ---------------------
	configController := v1.NewConfigController()
	adminApiV1Group.GET("/config", base.HandleRouteFunc(configController.GetSetting))
	adminApiV1Group.PUT("/config", base.HandleRouteFunc(configController.UpdateSetting))

	// ----------------- home ---------------------
	dashboardController := v1.NewDashboardController()
	adminApiV1Group.GET("/monitor/total", base.HandleRouteFunc(dashboardController.Total))

	// monitor_groups
	adminApiV1Group.GET("/monitor/groups", base.HandleRouteFunc(dashboardController.MonitorGroups))

	// ----------------------------------------------

	// ----------------- 服务相关 ---------------------
	serverController := v1.NewServerController()

	serverV1Group := adminApiV1Group.Group("/server")
	crudApi[do.ServerDO, uint](serverV1Group, serverController)
	serverV1Group.GET("/info/:id", base.HandleRouteFunc(serverController.GetServerInfo))
	serverV1Group.GET("/statistics/:id", base.HandleRouteFunc(serverController.GetServerStatisticsList))

	// 服务器分组
	crudApi[do.ServerGroupDO, uint](adminApiV1Group.Group("/server_group"), v1.NewServerGroupController())
	// 服务器故障
	serverFaultController := v1.NewServerFaultController()
	serverV1Group.GET("/fault/page", base.HandleRouteFunc(serverFaultController.Page))
	serverV1Group.PUT("/fault/:id", base.HandleRouteFunc(serverFaultController.Update))

	// 推送更新
	serverV1Group.PUT("/check_update/:id", base.HandleRouteFunc(serverController.CheckUpdate))
	// --------------------------------------------

	// ---------------  通知相关 -----------------------
	crudApi[do.NotifyGroupDO, uint](adminApiV1Group.Group("/notify_group"), v1.NewNotifyGroupController())
	crudApi[do.NotifyChannelDO, uint](adminApiV1Group.Group("/notify_channel"), v1.NewNotifyChannelController())

	notifyLogGroup := adminApiV1Group.Group("/notify_log")
	notifyLogGroup.GET("/", base.HandleRouteFunc(v1.NewNotifyLogController().List))
	// -------------------------------------------
}

func crudApi[DO any, ID int | uint | uint32](g *gin.RouterGroup, crud base.ICrudController[DO, ID]) {
	g.GET("/", base.HandleRouteFunc(crud.List))
	g.GET("/page", base.HandleRouteFunc(crud.Page))
	g.GET("/:id", base.HandleRouteFunc(crud.Get))
	g.POST("/", base.HandleRouteFunc(crud.Create))
	g.PUT("/:id", base.HandleRouteFunc(crud.Update))
	g.DELETE("/:id", base.HandleRouteFunc(crud.Delete))
}
