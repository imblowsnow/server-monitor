package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type ServerFaultController struct {
	serverFaultDao *dao.ServerFaultDao
	base.CrudController[dao.IBaseDao[do.ServerFaultDO, uint], do.ServerFaultDO, uint]
}

func NewServerFaultController() *ServerFaultController {
	var serverFaultDao = dao.NewServerFaultDao()
	return &ServerFaultController{
		serverFaultDao: serverFaultDao,
		CrudController: base.CrudController[dao.IBaseDao[do.ServerFaultDO, uint], do.ServerFaultDO, uint]{
			Dao: serverFaultDao,
		},
	}
}
