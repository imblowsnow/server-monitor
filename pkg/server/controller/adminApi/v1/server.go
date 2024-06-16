package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

var serverDao = dao.ServerDao{}

type ServerController struct {
	base.CrudController[dao.IBaseDao[do.ServerDO, uint], do.ServerDO, uint]
}

func NewServerController() ServerController {
	return ServerController{
		CrudController: base.CrudController[dao.IBaseDao[do.ServerDO, uint], do.ServerDO, uint]{
			Dao: &serverDao,
		},
	}
}
