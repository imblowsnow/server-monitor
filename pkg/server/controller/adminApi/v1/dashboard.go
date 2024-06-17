package v1

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/entity/bo"
)

type DashboardController struct {
	serverDao *dao.ServerDao
}

func NewDashboardController() *DashboardController {
	return &DashboardController{
		serverDao: dao.NewServerDao(),
	}
}

func (c DashboardController) Total(context *gin.Context) interface{} {
	// 在线服务器数量
	onlineNum := c.serverDao.GetStatusNum(enum.ServerStatusOnline)
	// 离线服务器数量
	offlineNum := c.serverDao.GetStatusNum(enum.ServerStatusOffline)

	result := bo.DashboardTotalBO{
		OnlineNum:  onlineNum,
		OfflineNum: offlineNum,
	}

	return result
}
