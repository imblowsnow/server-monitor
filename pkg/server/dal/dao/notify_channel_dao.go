package dao

import (
	"server-monitor/pkg/server/dal/do"
)

type NotifyChannelDao struct {
	IBaseDao[do.NotifyChannelDO, uint]
}

func NewNotifyChannelDao() *NotifyChannelDao {
	return &NotifyChannelDao{
		IBaseDao: &BaseDao[do.NotifyChannelDO, uint]{},
	}
}
