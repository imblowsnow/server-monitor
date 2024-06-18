package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyChannelController struct {
	dao *dao.NotifyChannelDao
	base.ICrudController[do.NotifyChannelDO, uint]
}

func NewNotifyChannelController() *NotifyChannelController {
	var notifyChannelDao = dao.NewNotifyChannelDao()
	return &NotifyChannelController{
		dao: notifyChannelDao,
		ICrudController: base.CrudController[do.NotifyChannelDO, uint]{
			Dao: notifyChannelDao,
		},
	}
}
