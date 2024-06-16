package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyLogController struct {
	base.CrudController[dao.IBaseDao[do.NotifyLogDO, uint], do.NotifyLogDO, uint]
}
