package dao

import "server-monitor/pkg/server/dal/do"

type NotifyGroupDao struct {
	BaseDao[do.NotifyGroupDO, uint]
}
