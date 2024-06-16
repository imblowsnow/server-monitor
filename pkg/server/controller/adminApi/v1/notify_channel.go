package v1

import (
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type NotifyChannelController struct {
	base.CrudController[dao.IBaseDao[do.NotifyChannelDO, uint], do.NotifyChannelDO, uint]
}
