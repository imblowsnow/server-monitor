package dao

import (
	"server-monitor/pkg/server/dal/do"
)

type NotifyChannelDao struct {
	BaseDao[do.NotifyChannelDO, uint]
}
