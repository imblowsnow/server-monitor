package dao

import "server-monitor/pkg/server/dal/do"

type NotifyLogDao struct {
	BaseDao[do.NotifyLogDO, uint]
}
