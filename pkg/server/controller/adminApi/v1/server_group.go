package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

var serverGroupDao = dao.NewServerGroupDao()

type ServerGroupController struct {
	base.CrudController[dao.IBaseDao[do.ServerGroupDO, uint], do.ServerGroupDO, uint]
}

func NewServerGroupController() ServerGroupController {
	return ServerGroupController{
		CrudController: base.CrudController[dao.IBaseDao[do.ServerGroupDO, uint], do.ServerGroupDO, uint]{
			Dao: serverGroupDao,
		},
	}
}
