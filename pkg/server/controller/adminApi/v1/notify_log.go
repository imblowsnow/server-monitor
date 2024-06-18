package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyLogController struct {
	notifyLogDao *dao.NotifyLogDao
	base.ICrudController[do.NotifyLogDO, uint]
}

func NewNotifyLogController() *NotifyLogController {
	var notifyLogDao = dao.NewNotifyLogDao()
	return &NotifyLogController{
		notifyLogDao: notifyLogDao,
		ICrudController: base.CrudController[do.NotifyLogDO, uint]{
			Dao: notifyLogDao,
		},
	}
}
