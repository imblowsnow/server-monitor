package dao

import (
	"fmt"
	"github.com/google/uuid"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/dal/do"
	"time"
)

type ServerDao struct {
	IBaseDao[do.ServerDO, uint]
}

func NewServerDao() *ServerDao {
	return &ServerDao{
		IBaseDao: &BaseDao[do.ServerDO, uint]{},
	}
}

func (dao ServerDao) GetServerInfo(id uint) *bo.ServerInfoBO {
	var Server do.ServerDO
	result := dao.DB().Where("id = ?", id).First(&Server)

	if result.Error != nil {
		return nil
	}

	var serverInfo do.ServerInfoDO
	dao.DB().Where("server_id = ?", id).First(&serverInfo)

	// 获取最后一次状态
	var serverState do.ServerStateDO
	dao.DB().Where("server_id = ?", id).Order("create_time desc").First(&serverState)

	return &bo.ServerInfoBO{
		ID:             Server.ID,
		Name:           Server.Name,
		Status:         Server.Status,
		Ip:             Server.IP,
		LastOnlineTime: Server.LastOnlineTime,
		ServerInfoDO:   &serverInfo,
		ServerStateDO:  &serverState,
	}
}

func (dao ServerDao) GetServerInfoList() []*bo.ServerInfoBO {
	list := dao.GetList(nil)
	if list == nil {
		return nil
	}
	var result []*bo.ServerInfoBO
	for _, v := range list {
		result = append(result, dao.GetServerInfo(v.ID))
	}
	return result
}

func (dao ServerDao) GetServerInfoListByGroupId(groupId uint) []*bo.ServerInfoBO {
	var list []do.ServerDO
	dao.DB().Where("group_id = ?", groupId).Find(&list)

	var result []*bo.ServerInfoBO
	for _, v := range list {
		result = append(result, dao.GetServerInfo(v.ID))
	}
	return result
}

func (dao ServerDao) GetServerGroups() []*bo.ServerGroupBO {
	// 获取分组列表
	var serverGroupDao = NewServerGroupDao()
	groups := serverGroupDao.GetList(nil)

	var list []*bo.ServerGroupBO

	for _, v := range groups {
		var result bo.ServerGroupBO

		serverInfos := dao.GetServerInfoListByGroupId(v.ID)
		result.GroupName = v.GroupName
		result.Servers = serverInfos

		list = append(list, &result)
	}

	return list
}

func (dao ServerDao) UpdateLastHeartbeatTime(serverId uint) {
	dao.DB().Model(&do.ServerDO{}).Where("id = ?", serverId).Update("last_online_time", time.Now())
}

func (dao ServerDao) Add(server *do.ServerDO) error {
	server.Key = uuid.New().String()
	return dao.IBaseDao.Add(server)
}

func (dao ServerDao) GetStatusNum(status enum.ServerStatus) int64 {
	var count int64
	dao.DB().Model(&do.ServerDO{}).Where("status = ?", status).Count(&count)
	return count
}

func (dao ServerDao) GetMonitorServers(groupId uint) []*bo.MonitorServerBO {
	var list []do.ServerDO
	dao.DB().Select("id").Where("group_id = ?", groupId).Find(&list)

	var result []*bo.MonitorServerBO
	for _, v := range list {
		servers := dao.GetMonitorServer(v.ID)
		result = append(result, servers)
	}
	return result
}

func (dao ServerDao) GetMonitorServer(serverId uint) *bo.MonitorServerBO {
	var server *do.ServerDO
	dao.DB().Where("id = ?", serverId).First(&server)

	if server == nil {
		return nil
	}

	monitorServerStatisticsList := dao.GetMonitorServerStatisticsList(serverId, time.Hour, 24)

	monitorServer := bo.MonitorServerBO{
		ServerId:         server.ID,
		ServerName:       server.Name,
		ServerStatus:     server.Status,
		LastOnlineTime:   server.LastOnlineTime,
		OnlineRate:       dao.calcMonitorServerStatisticsRate(monitorServerStatisticsList),
		OnlineStatistics: monitorServerStatisticsList,
	}

	return &monitorServer
}

func (dao ServerDao) GetMonitorServerStatisticsList(serverId uint, duration time.Duration, num int) []*bo.MonitorServerStatisticsBO {
	fmt.Println(dao.DB())
	// 获取最近num条事件数据
	var list []*bo.MonitorServerStatisticsBO

	for i := 0; i < num; i++ {
		var startTime, endTime string

		subDuration := -duration * time.Duration(i)
		addDuration := subDuration + (time.Duration(1) * duration)
		if duration >= time.Hour*24*30 {
			// 按月
			startTime = time.Now().Add(subDuration).Format("2006-01") + "-01 00:00:00"
			endTime = time.Now().Add(addDuration).Format("2006-01") + "-30 23:59:59"
		} else if duration >= time.Hour*24 {
			// 按天
			startTime = time.Now().Add(subDuration).Format("2006-01-02") + " 00:00:00"
			endTime = time.Now().Add(addDuration).Format("2006-01-02") + " 00:00:00"
		} else if duration >= time.Hour {
			// 按小时
			startTime = time.Now().Add(subDuration).Format("2006-01-02 15") + ":00:00"
			endTime = time.Now().Add(addDuration).Format("2006-01-02 15") + ":00:00"
		} else {
			// 按分钟
			startTime = time.Now().Add(subDuration).Format("2006-01-02 15:04") + ":00"
			endTime = time.Now().Add(addDuration).Format("2006-01-02 15:04") + ":00"
		}
		fmt.Println(startTime, endTime)
		list = append(list, dao.GetMonitorServerStatistics(serverId, duration, startTime, endTime))
	}

	return list
}

func (dao ServerDao) GetMonitorServerStatistics(serverId uint, duration time.Duration, startTime, endTime string) *bo.MonitorServerStatisticsBO {
	var total = int64(duration / time.Minute)
	var count int64
	dao.DB().Model(&do.ServerStateDO{}).Where("server_id = ?", serverId).Where("create_time >= ?", startTime).Where("create_time < ?", endTime).Count(&count)
	return &bo.MonitorServerStatisticsBO{
		StartTime: startTime,
		EndTime:   endTime,
		Rate:      int(float64(count) / float64(total) * 100),
	}
}

func (dao ServerDao) calcMonitorServerStatisticsRate(statisticsList []*bo.MonitorServerStatisticsBO) int {
	total := 0
	for _, statisticsBO := range statisticsList {
		total += statisticsBO.Rate
	}
	return total / len(statisticsList)
}
