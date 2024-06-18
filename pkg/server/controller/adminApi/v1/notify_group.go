package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyGroupController struct {
	notifyGroupDao *dao.NotifyGroupDao
	base.ICrudController[do.NotifyGroupDO, uint]
}

func NewNotifyGroupController() *NotifyGroupController {
	var notifyGroupDao = dao.NewNotifyGroupDao()
	return &NotifyGroupController{
		notifyGroupDao: notifyGroupDao,
		ICrudController: base.CrudController[do.NotifyGroupDO, uint]{
			Dao: notifyGroupDao,
		},
	}
}
