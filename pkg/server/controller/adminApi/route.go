package adminApi

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/utils"
	v1 "server-monitor/pkg/server/controller/adminApi/v1"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/do"
)

func InitRoute(r *gin.Engine) {
	adminApiV1Group := r.Group("/admin-api/v1")

	// ----------------- home ---------------------
	dashboardController := v1.NewDashboardController()
	adminApiV1Group.GET("/dashboard/total", handleRouteFunc(dashboardController.Total))

	// monitor_groups
	adminApiV1Group.GET("/monitor/groups", handleRouteFunc(dashboardController.MonitorGroups))
	// ----------------------------------------------

	// ----------------- 服务相关 ---------------------
	serverController := v1.NewServerController()

	serverV1Group := adminApiV1Group.Group("/server")
	crudApi[do.ServerDO, uint](serverV1Group, serverController)
	serverV1Group.GET("/info/:id", handleRouteFunc(serverController.GetServerInfo))
	serverV1Group.GET("/statistics/:id", handleRouteFunc(serverController.GetServerStatisticsList))

	// 服务器分组
	crudApi[do.ServerGroupDO, uint](adminApiV1Group.Group("/server_group"), v1.NewServerGroupController())
	// 服务器故障
	serverFaultController := v1.NewServerFaultController()
	serverV1Group.GET("/fault/page", handleRouteFunc(serverFaultController.Page))
	// --------------------------------------------

	// ---------------  通知相关 -----------------------
	crudApi[do.NotifyGroupDO, uint](adminApiV1Group.Group("/notify_group"), v1.NewNotifyGroupController())
	crudApi[do.NotifyChannelDO, uint](adminApiV1Group.Group("/notify_channel"), v1.NewNotifyChannelController())

	notifyLogGroup := adminApiV1Group.Group("/notify_log")
	notifyLogGroup.GET("/", handleRouteFunc(v1.NewNotifyLogController().List))
	// -------------------------------------------

}

func crudApi[DO any, ID int | uint | uint32](g *gin.RouterGroup, crud base.ICrudController[DO, ID]) {
	g.GET("/", handleRouteFunc(crud.List))
	g.GET("/page", handleRouteFunc(crud.Page))
	g.GET("/:id", handleRouteFunc(crud.Get))
	g.POST("/", handleRouteFunc(crud.Create))
	g.PUT("/:id", handleRouteFunc(crud.Update))
	g.DELETE("/:id", handleRouteFunc(crud.Delete))
}

func handleRouteFunc(method func(context *gin.Context) interface{}) func(context *gin.Context) {
	return func(context *gin.Context) {
		result := method(context)
		if result != nil {
			if err, ok := result.(error); ok {
				context.JSON(200, utils.ResultError(err.Error()))
				return
			}
		}
		context.JSON(200, utils.ResultSuccess(result))
	}
}
