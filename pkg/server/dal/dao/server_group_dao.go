package dao

import "server-monitor/pkg/server/dal/do"

type ServerGroupDao struct {
	BaseDao[do.ServerGroupDO, uint]
}
