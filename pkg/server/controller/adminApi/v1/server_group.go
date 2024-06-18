package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type ServerGroupController struct {
	serverGroupDao *dao.ServerGroupDao
	base.ICrudController[do.ServerGroupDO, uint]
}

func NewServerGroupController() *ServerGroupController {
	var serverGroupDao = dao.NewServerGroupDao()
	return &ServerGroupController{
		serverGroupDao: serverGroupDao,
		ICrudController: base.CrudController[do.ServerGroupDO, uint]{
			Dao: serverGroupDao,
		},
	}
}
