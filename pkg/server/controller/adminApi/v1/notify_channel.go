package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyChannelController struct {
	dao *dao.NotifyChannelDao
	base.CrudController[dao.IBaseDao[do.NotifyChannelDO, uint], do.NotifyChannelDO, uint]
}

func NewNotifyChannelController() *NotifyChannelController {
	var notifyChannelDao = dao.NotifyChannelDao{}
	return &NotifyChannelController{
		dao: &notifyChannelDao,
		CrudController: base.CrudController[dao.IBaseDao[do.NotifyChannelDO, uint], do.NotifyChannelDO, uint]{
			Dao: &notifyChannelDao,
		},
	}
}
