package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyLogController struct {
	notifyLogDao *dao.NotifyLogDao
	base.CrudController[dao.IBaseDao[do.NotifyLogDO, uint], do.NotifyLogDO, uint]
}

func NewNotifyLogController() *NotifyLogController {
	var notifyLogDao = dao.NotifyLogDao{}
	return &NotifyLogController{
		notifyLogDao: &notifyLogDao,
		CrudController: base.CrudController[dao.IBaseDao[do.NotifyLogDO, uint], do.NotifyLogDO, uint]{
			Dao: &notifyLogDao,
		},
	}
}
