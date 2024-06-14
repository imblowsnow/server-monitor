package dao

import "server-monitor/pkg/server/dal/do"

type ServerStateDao struct {
	BaseDao[do.ServerStateDO, uint]
}
