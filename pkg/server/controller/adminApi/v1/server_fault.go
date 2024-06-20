package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/entity"
	"server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
	"strconv"
)

type ServerFaultController struct {
	serverFaultDao *dao.ServerFaultDao
	base.ICrudController[do.ServerFaultDO, uint]
}

func NewServerFaultController() *ServerFaultController {
	var serverFaultDao = dao.NewServerFaultDao()
	return &ServerFaultController{
		serverFaultDao: serverFaultDao,
		ICrudController: base.CrudController[do.ServerFaultDO, uint]{
			Dao: serverFaultDao,
		},
	}
}

func (c ServerFaultController) Page(context *gin.Context) interface{} {
	currentPageStr := context.Query("page")
	pageSizeStr := context.Query("size")
	currentPage, err := strconv.Atoi(currentPageStr)
	serverId, _ := strconv.Atoi(context.Query("serverId"))
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
	var pageE entity.Page[do.ServerFaultDO]
	querys := map[string]interface{}{}
	if serverId != 0 {
		querys["server_id = ?"] = serverId
	}
	pageE.Page = currentPage
	pageE.PageSize = pageSize
	err = c.GetDao().Page(&pageE, querys)
	if err != nil {
		return err
	}

	var page entity.Page[bo.ServerFaultBO]
	page.Page = pageE.Page
	page.PageSize = pageE.PageSize
	page.Total = pageE.Total

	serverDao := dao.NewServerDao()
	// 转换为BO
	for _, item := range pageE.List {
		serverFaultBO := bo.ServerFaultBO{
			ServerFaultDO: item,
		}
		server := serverDao.GetById(item.ServerId)
		if server != nil {
			serverFaultBO.ServerName = server.Name
		}
		page.List = append(page.List, serverFaultBO)
	}
	return page
}
