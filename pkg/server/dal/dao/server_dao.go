package dao

import (
	"github.com/google/uuid"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/server/common/entity/bo"
	enum2 "server-monitor/pkg/server/common/enum"
	"server-monitor/pkg/server/dal/do"
	"time"
)

type ServerDao struct {
	IBaseDao[do.ServerDO, uint]
}

func NewServerDao() *ServerDao {
	return &ServerDao{
		IBaseDao: &BaseDao[do.ServerDO, uint]{
			Order: "sort desc, id asc",
		},
	}
}

func (dao ServerDao) GetServerInfo(id uint) *bo.ServerInfoBO {
	var Server do.ServerDO
	result := dao.DB().Where("id = ?", id).First(&Server)

	if result.Error != nil {
		return nil
	}

	var serverInfo *do.ServerInfoDO
	dao.DB().Where("server_id = ?", id).First(&serverInfo)
	if serverInfo.ID == 0 {
		serverInfo = nil
	}

	// 获取最后一次状态
	var serverState *do.ServerStateDO
	dao.DB().Where("server_id = ?", id).Order("create_time desc").First(&serverState)
	if serverState.ID == 0 {
		serverState = nil
	}

	return &bo.ServerInfoBO{
		ID:             Server.ID,
		Name:           Server.Name,
		Remark:         Server.Remark,
		Status:         Server.Status,
		Ip:             Server.IP,
		Key:            Server.Key,
		LastOnlineTime: Server.LastOnlineTime,
		ServerInfoDO:   serverInfo,
		ServerStateDO:  serverState,
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
	dao.DB().Model(&do.ServerDO{}).Where("id = ?", serverId).UpdateColumns(map[string]interface{}{
		"last_online_time": time.Now(),
		"status":           enum.ServerStatusOnline,
	})
}

func (dao ServerDao) UpdateStatus(serverId uint, status enum.ServerStatus) {
	dao.DB().Model(&do.ServerDO{}).Where("id = ?", serverId).Update("status", status)
}

func (dao ServerDao) Add(server *do.ServerDO) error {
	server.Key = uuid.New().String()
	return dao.IBaseDao.Add(server)
}

func (dao ServerDao) GetStatusNum(status enum.ServerStatus, fromIndex bool) int64 {
	var count int64
	query := dao.DB().Model(&do.ServerDO{})
	if fromIndex {
		query = query.Where("show_index = ?", 1)
	}
	query.Where("status = ?", status).Count(&count)
	return count
}

func (dao ServerDao) GetMonitorServers(groupId uint, fromIndex bool, monitorDurationEnum enum2.MonitorDurationEnum) []*bo.MonitorServerBO {
	var list []do.ServerDO

	query := dao.DB().Select("id")
	if fromIndex {
		query = query.Where("show_index = ?", 1)
	}
	query.Where("group_id = ?", groupId).Find(&list)

	var result []*bo.MonitorServerBO
	for _, v := range list {
		servers := dao.GetMonitorServer(v.ID, monitorDurationEnum)
		result = append(result, servers)
	}
	return result
}

// 获取服务器监控统计数据
func (dao ServerDao) GetMonitorServer(serverId uint, monitorDurationEnum enum2.MonitorDurationEnum) *bo.MonitorServerBO {
	var server *do.ServerDO
	dao.DB().Where("id = ?", serverId).First(&server)

	if server == nil {
		return nil
	}

	monitorServerStatisticsList := dao.GetMonitorServerStatisticsList(serverId, monitorDurationEnum)

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

// 获取服务器监控统计数据列表
func (dao ServerDao) GetMonitorServerStatisticsList(serverId uint, monitorDurationEnum enum2.MonitorDurationEnum) []*bo.MonitorServerStatisticsBO {
	// 获取最近num条事件数据
	var list []*bo.MonitorServerStatisticsBO
	duration := monitorDurationEnum.Duration
	num := monitorDurationEnum.Num
	for i := 0; i < num; i++ {
		var startTime, endTime string

		subDuration := -duration * time.Duration(i)
		addDuration := subDuration + (time.Duration(1) * duration)
		if monitorDurationEnum == enum2.MONITOR_DURATION_MONTH {
			// 按月
			startTime = time.Now().Add(subDuration).Format("2006-01") + "-01 00:00:00"
			endTime = time.Now().Add(addDuration).Format("2006-01") + "-01 00:00:00"
			if startTime == endTime {
				subDuration = -duration * time.Duration(i+1)
				addDuration = -duration * time.Duration(i)

				startTime = time.Now().Add(subDuration).Format("2006-01") + "-01 00:00:00"
				endTime = time.Now().Add(addDuration).Format("2006-01") + "-01 00:00:00"
			}
		} else if monitorDurationEnum == enum2.MONITOR_DURATION_DAY {
			// 按天
			startTime = time.Now().Add(subDuration).Format("2006-01-02") + " 00:00:00"
			endTime = time.Now().Add(addDuration).Format("2006-01-02") + " 00:00:00"
			if startTime == endTime {
				subDuration = -duration * time.Duration(i+1)
				addDuration = -duration * time.Duration(i)

				startTime = time.Now().Add(subDuration).Format("2006-01-02") + " 00:00:00"
				endTime = time.Now().Add(addDuration).Format("2006-01-02") + " 00:00:00"
			}
		} else if monitorDurationEnum == enum2.MONITOR_DURATION_HOUR {
			// 按小时
			startTime = time.Now().Add(subDuration).Format("2006-01-02 15") + ":00:00"
			endTime = time.Now().Add(addDuration).Format("2006-01-02 15") + ":00:00"

			if startTime == endTime {
				subDuration = -duration * time.Duration(i+1)
				addDuration = -duration * time.Duration(i)

				startTime = time.Now().Add(subDuration).Format("2006-01-02 15") + ":00:00"
				endTime = time.Now().Add(addDuration).Format("2006-01-02 15") + ":00:00"
			}
		} else {
			// 按分钟
			startTime = time.Now().Add(subDuration).Format("2006-01-02 15:04") + ":00"
			endTime = time.Now().Add(addDuration).Format("2006-01-02 15:04") + ":00"

			if startTime == endTime {
				subDuration = -duration * time.Duration(i+1)
				addDuration = -duration * time.Duration(i)

				startTime = time.Now().Add(subDuration).Format("2006-01-02 15:04") + ":00"
				endTime = time.Now().Add(addDuration).Format("2006-01-02 15:04") + ":00"
			}
		}

		serverStatistics := dao.GetMonitorServerStatistics(serverId, startTime, endTime)
		list = append(list, serverStatistics)
	}

	return list
}

func (dao ServerDao) GetMonitorServerStatistics(serverId uint, startTime, endTime string) *bo.MonitorServerStatisticsBO {
	startTimeV, _ := time.Parse("2006-01-02 15:04:05", startTime)
	endTimeV, _ := time.Parse("2006-01-02 15:04:05", endTime)
	nowTimeV, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if endTimeV.After(nowTimeV) {
		endTime = time.Now().Format("2006-01-02 15:04") + ":00"
		endTimeV, _ = time.Parse("2006-01-02 15:04:05", endTime)
	}
	var duration = endTimeV.Sub(startTimeV)
	var total = int64(duration / time.Minute)
	var count int64
	if total == 0 {
		return &bo.MonitorServerStatisticsBO{
			StartTime: startTime,
			EndTime:   endTime,
			Rate:      100,
			Total:     total,
			Online:    count,
		}
	}
	dao.DB().Model(&do.ServerStateDO{}).Where("server_id = ?", serverId).Where("create_time >= ?", startTime).Where("create_time < ?", endTime).Count(&count)
	return &bo.MonitorServerStatisticsBO{
		StartTime: startTime,
		EndTime:   endTime,
		Rate:      int(float64(count) / float64(total) * 100),
		Total:     total,
		Online:    count,
	}
}

// 计算监控服务器在线率
func (dao ServerDao) calcMonitorServerStatisticsRate(statisticsList []*bo.MonitorServerStatisticsBO) int {
	total := 0
	for _, statisticsBO := range statisticsList {
		total += statisticsBO.Rate
	}
	return total / len(statisticsList)
}

// 检查服务端掉线>5分钟
func (dao ServerDao) CheckServerOffline() {
	lastOnlineTime := time.Now().Add(-time.Minute * 5)
	var list []do.ServerDO
	dao.DB().Model(&do.ServerDO{}).Where("last_online_time <= ?", lastOnlineTime).Find(&list)
	if len(list) == 0 {
		return
	}

	serverFaultDao := NewServerFaultDao()

	ids := make([]uint, 0)

	for _, v := range list {
		// 记录故障
		serverFaultDao.RecordFaultAndTime(v.ID, "服务端掉线", time.Time(v.LastOnlineTime))

		ids = append(ids, v.ID)
	}

	// 更新状态 where in ids
	dao.DB().Model(&do.ServerDO{}).Where("id in (?)", ids).Update("status", enum.ServerStatusOffline)
}

func (dao ServerDao) GetServerIdByKey(key string) uint {
	var server do.ServerDO
	dao.DB().Where("key = ?", key).First(&server)
	return server.ID
}
