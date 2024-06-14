package dao

import (
	"server-monitor/pkg/server/dal/do"
	"server-monitor/pkg/server/entity/bo"
)

type ServerDao struct {
	BaseDao[do.ServerDO]
}

func (ServerDao) GetServerInfo() bo.ServerInfoBO {
	return bo.ServerInfoBO{}
}
