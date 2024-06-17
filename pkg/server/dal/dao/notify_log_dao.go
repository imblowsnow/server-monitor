package dao

import "server-monitor/pkg/server/dal/do"

type NotifyLogDao struct {
	IBaseDao[do.NotifyLogDO, uint]
}

func NewNotifyLogDao() *NotifyLogDao {
	return &NotifyLogDao{
		IBaseDao: &BaseDao[do.NotifyLogDO, uint]{},
	}
}
