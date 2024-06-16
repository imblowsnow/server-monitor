package dao

import (
	"server-monitor/pkg/server/dal/do"
)

type ServerGroupDao struct {
	BaseDao[do.ServerGroupDO, uint]
}

func NewServerGroupDao() *ServerGroupDao {
	return &ServerGroupDao{
		BaseDao: BaseDao[do.ServerGroupDO, uint]{
			Order: "sort desc, id asc",
		},
	}
}
