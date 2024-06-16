package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyGroupController struct {
	base.CrudController[dao.IBaseDao[do.NotifyGroupDO, uint], do.NotifyGroupDO, uint]
}
