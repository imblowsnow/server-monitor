package v1

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

var serverDao = dao.ServerDao{}

type ServerController struct {
	base.CrudController[dao.IBaseDao[do.ServerDO, uint], do.ServerDO, uint]
}

func (ServerController) GetServerGroups(context *gin.Context) interface{} {
	return serverDao.GetServerGroups()
}
