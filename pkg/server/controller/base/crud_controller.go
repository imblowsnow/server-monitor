package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/entity"
	"strconv"
)

type CrudController[DAO dao.IBaseDao[DO, ID], DO any, ID int | uint | uint32 | uint64 | string] struct {
	Dao dao.IBaseDao[DO, ID]
}

func (c CrudController[DAO, DO, ID]) List(context *gin.Context) interface{} {
	return c.Dao.GetList()
}

func (c CrudController[DAO, DO, ID]) Page(context *gin.Context) interface{} {
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
	var pageE entity.Page[DO]
	pageE.Page = currentPage
	pageE.PageSize = pageSize
	err = c.Dao.Page(&pageE)
	if err != nil {
		return err
	}
	return pageE
}

func (c CrudController[DAO, DO, ID]) Get(context *gin.Context) interface{} {
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	idValue := ID(id)
	return c.Dao.GetById(idValue)
}

func (c CrudController[DAO, DO, ID]) Create(context *gin.Context) interface{} {
	var newItem *DO
	if err := context.ShouldBindJSON(&newItem); err != nil {
		return err
	}
	return c.Dao.Add(newItem)
}

func (c CrudController[DAO, DO, ID]) Update(context *gin.Context) interface{} {
	var updatedItem *DO
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	if err := context.ShouldBindJSON(&updatedItem); err != nil {
		return err
	}
	return c.Dao.UpdateById(ID(id), updatedItem)
}

func (c CrudController[DAO, DO, ID]) Delete(context *gin.Context) interface{} {
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	return c.Dao.Delete(ID(id))
}
