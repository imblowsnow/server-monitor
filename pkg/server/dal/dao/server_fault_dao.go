package dao

import "server-monitor/pkg/server/dal/do"

type ServerFaultDao struct {
	IBaseDao[do.ServerFaultDO, uint]
}

func NewServerFaultDao() *ServerFaultDao {
	return &ServerFaultDao{
		IBaseDao: &BaseDao[do.ServerFaultDO, uint]{
			Order: "start_time desc",
		},
	}
}
