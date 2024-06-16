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

	serverController := v1.ServerController{}

	serverV1Group := adminApiV1Group.Group("/server")
	crudApi(serverV1Group, serverController.CrudController)
	serverV1Group.GET("/groups", handleHandlerFunc(serverController.GetServerGroups))

	crudApi(adminApiV1Group.Group("/server_group"), v1.NewServerGroupController().CrudController)
}

func crudApi[DAO dao.IBaseDao[DO, ID], DO any, ID int | uint | uint32](g *gin.RouterGroup, crud base.CrudController[DAO, DO, ID]) {
	g.GET("/", handleHandlerFunc(crud.List))
	g.GET("/page/:currentPage/:size", handleHandlerFunc(crud.Page))
	g.GET("/:id", handleHandlerFunc(crud.Get))
	g.POST("/", handleHandlerFunc(crud.Create))
	g.PUT("/:id", handleHandlerFunc(crud.Update))
	g.DELETE("/:id", handleHandlerFunc(crud.Delete))

}

func handleHandlerFunc(method func(context *gin.Context) interface{}) func(context *gin.Context) {
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
