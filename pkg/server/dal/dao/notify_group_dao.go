package dao

import "server-monitor/pkg/server/dal/do"

type NotifyGroupDao struct {
	IBaseDao[do.NotifyGroupDO, uint]
}

func NewNotifyGroupDao() *NotifyGroupDao {
	return &NotifyGroupDao{
		IBaseDao: &BaseDao[do.NotifyGroupDO, uint]{},
	}
}
