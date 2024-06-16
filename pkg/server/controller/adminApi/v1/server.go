package v1

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type ServerController struct {
	serverDao *dao.ServerDao
	base.CrudController[dao.IBaseDao[do.ServerDO, uint], do.ServerDO, uint]
}

func NewServerController() *ServerController {
	var serverDao = dao.ServerDao{}

	return &ServerController{
		serverDao: &serverDao,
		CrudController: base.CrudController[dao.IBaseDao[do.ServerDO, uint], do.ServerDO, uint]{
			Dao: &serverDao,
		},
	}
}

func (s ServerController) GetServerGroups(context *gin.Context) interface{} {
	return s.serverDao.GetServerGroups()
}
