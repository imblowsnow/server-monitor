package dao

import "server-monitor/pkg/server/dal/do"

type ServerStateDao struct {
	IBaseDao[do.ServerStateDO, uint]
}
