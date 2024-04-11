package service

import (
	"server/config"
	"server/dao"
	"server/entity/bo"
	"server/entity/do"
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
func (s *ServerService) GetIndexServerList() []*bo.ServerBO {
	serverList := serverDao.GetServerList(true)
	var serverBOList []*bo.ServerBO
	for _, server := range serverList {
		serverBO := s.GetServerInfoById(server.ID)
		serverBOList = append(serverBOList, serverBO)
	}

	return serverBOList
}

func (s *ServerService) GetServerInfoById(serverId uint) *bo.ServerBO {
	server, err := serverDao.GetServerById(serverId)
	if err != nil {
		return nil
	}

	return s.GetServerInfo(server)
}
func (s *ServerService) GetServerInfo(server do.Server) *bo.ServerBO {
	if server.ID == 0 {
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

func (s *ServerService) GetServerStatus() []bo.ServerFaultTotalBO {
	// 获取前30天
	today := time.Now()

	times := []string{}
	minDate := today.AddDate(0, 0, -30).Unix()
	maxDate := today.AddDate(0, 0, 1).Unix()
	// 使用循环获取前30天的日期
	for i := 0; i < 30; i++ {
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
		serverStateList := serverStateDao.GetServerStateListByTime(server.ID, minDate, maxDate)
		serverStateMaps := make(map[string][]int64)
		for _, state := range serverStateList {
			date := time.Unix(state.CreatedTime, 0).Format("2006-01-02")
			if _, ok := serverStateMaps[date]; !ok {
				serverStateMaps[date] = make([]int64, 0)
			}
			serverStateMaps[date] = append(serverStateMaps[date], state.CreatedTime)
		}

		var totalOnlineRate float64
		for _, startDate := range times {
			startDateTime, _ := time.ParseInLocation(dateLayout, startDate, time.Local)
			date := startDateTime.Format("2006-01-02")

			endDate := startDateTime.AddDate(0, 0, 1).Format(dateLayout)
			endDateTime, _ := time.ParseInLocation(dateLayout, endDate, time.Local)

			if endDateTime.After(time.Now()) {
				endDateTime = time.Now()
			}
			totalTime := endDateTime.Sub(startDateTime)
			faultTime := time.Duration(0) * time.Second

			onlineFlag := false
			if serverStateMaps[date] != nil && len(serverStateList) > 0 {
				serverStateList := serverStateMaps[date]
				// 插入第一个
				serverStateList = append([]int64{startDateTime.Unix()}, serverStateList...)
				serverStateList = append(serverStateList, endDateTime.Unix())
				for i := 0; i < len(serverStateList); i++ {
					// 计算2次状态之间的时间差是否大于5分钟
					if i+1 < len(serverStateList) {
						startTime := serverStateList[i]
						endTime := serverStateList[i+1]
						if endTime-startTime > 300 {
							// 故障时间
							faultTime += time.Duration(endTime-startTime) * time.Second
						}
					}
				}
				onlineFlag = true
			}

			faultRate := faultTime.Seconds() / totalTime.Seconds() * 100
			onlineRate := 100 - faultRate
			if !onlineFlag {
				onlineRate = 0
			}
			serverFaultTotal.Items = append(serverFaultTotal.Items, &bo.ServerFaultTotalItemBO{
				Time:       startDateTime.Format("2006-01-02"),
				TotalTime:  totalTime.Seconds(),
				FaultTime:  faultTime.Seconds(),
				FaultRate:  faultRate,
				OnlineRate: onlineRate,
				Online:     onlineFlag,
			})

			totalOnlineRate += onlineRate
		}
		serverFaultTotal.TotalOnlineRate = totalOnlineRate / float64(len(times))
		total = append(total, serverFaultTotal)
	}

	return total
}

func (s *ServerService) GetServerByKey(key string) (do.Server, error) {
	return serverDao.GetServerByKey(key)
}
