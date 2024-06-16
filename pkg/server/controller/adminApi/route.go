package adminApi

import (
	"github.com/gin-gonic/gin"
	v1 "server-monitor/pkg/server/controller/adminApi/v1"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/utils"
)

func InitRoute(r *gin.Engine) {
	adminApiV1Group := r.Group("/admin-api/v1")

	// ----------------- 服务相关 ---------------------
	serverController := v1.NewServerController()

	serverV1Group := adminApiV1Group.Group("/server")
	crudApi(serverV1Group, serverController.CrudController)
	serverV1Group.GET("/groups", handleRouteFunc(serverController.GetServerGroups))

	crudApi(adminApiV1Group.Group("/server_group"), v1.NewServerGroupController().CrudController)

	// --------------------------------------------

	// ---------------  通知相关 -----------------------
	crudApi(adminApiV1Group.Group("/notify_group"), v1.NewNotifyGroupController().CrudController)
	crudApi(adminApiV1Group.Group("/notify_channel"), v1.NewNotifyChannelController().CrudController)

	notifyLogGroup := adminApiV1Group.Group("/notify_log")
	notifyLogGroup.GET("/", handleRouteFunc(v1.NewNotifyLogController().List))
	// -------------------------------------------
}

func crudApi[DAO dao.IBaseDao[DO, ID], DO any, ID int | uint | uint32](g *gin.RouterGroup, crud base.CrudController[DAO, DO, ID]) {
	g.GET("/", handleRouteFunc(crud.List))
	g.GET("/page/:currentPage/:size", handleRouteFunc(crud.Page))
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
