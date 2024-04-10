package service

import (
	"server/config"
	"server/dao"
	"server/entity/bo"
	"server/entity/do"
	"server/utils"
	"strconv"
	"time"
)

var serverStateDao = dao.GetServerStateDao()
var serverInfoDao = dao.GetServerInfoDao()
var serverDao = dao.GetServerDao()
var serverFaultDao = dao.ServerFaultDao{}

type ServerService struct {
}

// 更新服务器ip
func (s ServerService) UpdateServer(server do.Server) {
	serverDao.UpdateServer(server)
}

// 获取服务器列表（首页展示）
func (s *ServerService) GetIndexServerList() []bo.ServerBO {
	serverList := serverDao.GetServerList(true)
	var serverBOList []bo.ServerBO
	for _, server := range serverList {
		serverInfoDO := serverInfoDao.GetServerInfoByServerID(server.ID)
		serverStateDO := serverStateDao.GetLastServerStateByServerID(server.ID)

		serverBO := bo.ServerBO{
			ID:          server.ID,
			Name:        server.Name,
			Group:       server.Group,
			Version:     server.Version,
			Hide:        server.Hide,
			Index:       server.Index,
			CreatedTime: server.CreatedTime,
			UpdatedTime: server.UpdatedTime,
			Live:        config.SocketClients.CheckExistByKey(strconv.Itoa(int(server.ID))),
		}
		if serverInfoDO != nil {
			serverBO.Info = &serverInfoDO.Info
		}
		if serverStateDO != nil {
			serverBO.State = &serverStateDO.State
		}
		serverBOList = append(serverBOList, serverBO)
	}

	return serverBOList
}

func (s *ServerService) GetServerInfo(serverId uint) *bo.ServerBO {
	server, err := serverDao.GetServerById(serverId)
	if err != nil {
		return nil
	}

	serverInfoDO := serverInfoDao.GetServerInfoByServerID(server.ID)
	serverStateDO := serverStateDao.GetLastServerStateByServerID(server.ID)

	serverBO := bo.ServerBO{
		ID:          server.ID,
		Name:        server.Name,
		Group:       server.Group,
		Version:     server.Version,
		Hide:        server.Hide,
		Index:       server.Index,
		CreatedTime: server.CreatedTime,
		UpdatedTime: server.UpdatedTime,
		Live:        config.SocketClients.CheckExistByKey(strconv.Itoa(int(server.ID))),
	}
	if serverInfoDO != nil {
		serverBO.Info = &serverInfoDO.Info
	}
	if serverStateDO != nil {
		serverBO.State = &serverStateDO.State
	}
	return &serverBO
}

// 获取服务每分钟统计状态
// timeType 1 分钟 2 小时 3 天 4 月
//func (s *ServerService) GetServerStatistics(timeType int) map[string]interface{} {
//	times := []string{}
//
//}

// 获取服务故障率 按天
func (s *ServerService) GetServerFaultTotal() []bo.ServerFaultTotalBO {
	// 获取前30天
	today := time.Now()

	times := []string{}
	minDate := today.AddDate(0, 0, -30).Unix()
	maxDate := today.AddDate(0, 0, 1).Unix()
	// 使用循环获取前30天的日期
	for i := 0; i <= 30; i++ {
		// 计算前i天的日期
		previousDate := today.AddDate(0, 0, -i)
		// 格式化日期为字符串
		date := previousDate.Format("2006-01-02")
		times = append(times, date+" 00:00:00")
	}

	total := make([]bo.ServerFaultTotalBO, 0)
	// 获取服务器列表
	serverList := serverDao.GetServerList(true)
	dateLayout := "2006-01-02 15:04:05"
	for _, server := range serverList {
		serverFaultTotal := bo.ServerFaultTotalBO{
			ServerID: server.ID,
			Name:     server.Name,
			Group:    server.Group,
			Items:    make([]*bo.ServerFaultTotalItemBO, 0),
		}
		serverFaultList := serverFaultDao.GetServerFaultListByTime(server.ID, minDate, maxDate)

		for _, startDate := range times {
			startDateTime, _ := time.ParseInLocation(dateLayout, startDate, time.Local)
			endDate := startDateTime.AddDate(0, 0, 1).Format(dateLayout)
			endDateTime, _ := time.ParseInLocation(dateLayout, endDate, time.Local)

			if endDateTime.After(time.Now()) {
				endDateTime = time.Now()
			}
			totalTime := endDateTime.Sub(startDateTime)
			faultTime := time.Duration(0)

			tmpServerFaultList := make([]do.ServerFault, 0)
			for _, fault := range serverFaultList {
				// 获取交集时间段
				faultStartTime, _ := time.ParseInLocation(dateLayout, time.Unix(fault.StartTime, 0).Format(dateLayout), time.Local)
				var faultEndTime time.Time
				if fault.EndTime == 0 {
					faultEndTime = time.Now()
				} else {
					faultEndTime, _ = time.ParseInLocation(dateLayout, time.Unix(fault.EndTime, 0).Format(dateLayout), time.Local)
				}

				// 获取交集时间段
				intersectionStart, intersectionEnd := utils.Intersection(startDateTime, endDateTime, faultStartTime, faultEndTime)

				if intersectionStart.IsZero() || intersectionEnd.IsZero() {
					continue
				}

				faultTime += intersectionEnd.Sub(intersectionStart)
				tmpServerFaultList = append(tmpServerFaultList, fault)
			}

			serverFaultTotal.Items = append(serverFaultTotal.Items, &bo.ServerFaultTotalItemBO{
				Time:      startDateTime.Format("2006-01-02"),
				TotalTime: totalTime.Seconds(),
				FaultTime: faultTime.Seconds(),
				FaultRate: faultTime.Seconds() / totalTime.Seconds() * 100,
			})
		}

		total = append(total, serverFaultTotal)
	}

	return total
}

func (s *ServerService) GetServerByKey(key string) (do.Server, error) {
	return serverDao.GetServerByKey(key)
}
