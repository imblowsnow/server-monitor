package v1

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/server/common/entity/bo"
	enum2 "server-monitor/pkg/server/common/enum"
	"server-monitor/pkg/server/dal/dao"
)

type DashboardController struct {
	serverDao      *dao.ServerDao
	serverGroupDao *dao.ServerGroupDao
}

func NewDashboardController() *DashboardController {
	return &DashboardController{
		serverDao:      dao.NewServerDao(),
		serverGroupDao: dao.NewServerGroupDao(),
	}
}

func (c DashboardController) Total(context *gin.Context) interface{} {
	// 在线服务器数量
	onlineNum := c.serverDao.GetStatusNum(enum.ServerStatusOnline, false)
	// 离线服务器数量
	offlineNum := c.serverDao.GetStatusNum(enum.ServerStatusOffline, false)

	result := bo.DashboardTotalBO{
		OnlineNum:  onlineNum,
		OfflineNum: offlineNum,
	}

	return result
}

// 获取
func (s DashboardController) MonitorGroups(context *gin.Context) interface{} {
	var monitorGroups []*bo.MonitorGroupBO
	serverGroups := s.serverGroupDao.GetList(nil)
	for i := range serverGroups {

		monitorGroup := bo.MonitorGroupBO{
			GroupId:   serverGroups[i].ID,
			GroupName: serverGroups[i].GroupName,
			Servers:   s.serverDao.GetMonitorServers(serverGroups[i].ID, false, enum2.MONITOR_DURATION_HOUR),
		}

		if monitorGroup.Servers != nil {
			monitorGroups = append(monitorGroups, &monitorGroup)
		}
	}

	return monitorGroups
}
