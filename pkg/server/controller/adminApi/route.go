package adminApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "server-monitor/pkg/server/controller/adminApi/v1"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"strconv"
)

func InitRoute(r *gin.Engine) {
	adminApiV1Group := r.Group("/admin-api/v1")
	crudApi(adminApiV1Group.Group("/server"), v1.NewServerController().CrudController)
	crudApi(adminApiV1Group.Group("/server_group"), v1.NewServerGroupController().CrudController)
}

func crudApi[DAO dao.IBaseDao[DO, ID], DO any, ID int | uint | uint32](g *gin.RouterGroup, crud base.CrudController[DAO, DO, ID]) {
	g.GET("/", func(context *gin.Context) {
		list := crud.List()
		context.JSON(200, list)
	})
	g.GET("/page/:currentPage/:size", func(context *gin.Context) {
		currentPageStr := context.Param("currentPage")
		pageSizeStr := context.Param("size")
		currentPage, err := strconv.Atoi(currentPageStr)
		if err != nil {
			fmt.Println("页数解析失败", err)
			currentPage = 1
		}
		if currentPage == 0 {
			currentPage = 1
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			fmt.Println("每页数量解析失败", err)
			pageSize = 10
		}
		if pageSize == 0 {
			pageSize = 10
		}

		page, err := crud.Page(currentPage, pageSize)
		context.JSON(200, page)
	})
	g.GET("/:id", func(context *gin.Context) {
		// Assuming you need to use ID here
		idParam := context.Param("id")
		id, _ := strconv.ParseUint(idParam, 10, 64)
		item := crud.Get(ID(id))
		context.JSON(200, item)
	})
	g.POST("/", func(context *gin.Context) {
		var newItem DO
		if err := context.ShouldBindJSON(&newItem); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := crud.Create(&newItem)
		if err != nil {
			fmt.Println("创建失败", err)
			context.JSON(500, gin.H{"error": "创建失败"})
			return
		}
		context.JSON(200, newItem)
	})
	g.PUT("/:id", func(context *gin.Context) {
		var updatedItem DO
		idParam := context.Param("id")
		id, _ := strconv.ParseUint(idParam, 10, 64)
		if err := context.ShouldBindJSON(&updatedItem); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := crud.Update(ID(id), &updatedItem)
		if err != nil {
			fmt.Println("更新失败", id, err)
			context.JSON(500, gin.H{"error": "更新失败"})
			return
		}
		context.JSON(200, updatedItem)
	})
	g.DELETE("/:id", func(context *gin.Context) {
		idParam := context.Param("id")
		id, _ := strconv.ParseUint(idParam, 10, 64)
		err := crud.Delete(ID(id))
		if err != nil {
			fmt.Println("删除失败", id, err)
			context.JSON(500, gin.H{"error": "删除失败"})
			return
		}
		context.Status(204)
	})
}
