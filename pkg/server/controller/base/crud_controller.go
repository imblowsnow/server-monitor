package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/entity"
	"server-monitor/pkg/server/dal/dao"
	"strconv"
)

type CrudController[DO any, ID int | uint | uint32 | uint64 | string] struct {
	Dao dao.IBaseDao[DO, ID]
}

func (c CrudController[DO, ID]) List(context *gin.Context) interface{} {
	return c.Dao.GetList(nil)
}

func (c CrudController[DO, ID]) Page(context *gin.Context) interface{} {
	currentPageStr := context.Query("page")
	pageSizeStr := context.Query("size")
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
	err = c.Dao.Page(&pageE, nil)
	if err != nil {
		return err
	}
	return pageE
}

func (c CrudController[DO, ID]) Get(context *gin.Context) interface{} {
	id := c.parseId(context)
	return c.Dao.GetById(id)
}

func (c CrudController[DO, ID]) Create(context *gin.Context) interface{} {
	var newItem *DO
	if err := context.ShouldBindJSON(&newItem); err != nil {
		return err
	}
	return c.Dao.Add(newItem)
}

func (c CrudController[DO, ID]) Update(context *gin.Context) interface{} {
	var updatedItem *DO
	if err := context.ShouldBindJSON(&updatedItem); err != nil {
		return err
	}

	id := c.parseId(context)

	return c.Dao.UpdateById(id, updatedItem)
}

func (c CrudController[DO, ID]) Delete(context *gin.Context) interface{} {
	id := c.parseId(context)
	return c.Dao.Delete(id)
}

func (c CrudController[DO, ID]) parseId(context *gin.Context) ID {
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	return ID(id)
}

func (c CrudController[DO, ID]) GetDao() dao.IBaseDao[DO, ID] {
	return c.Dao
}
