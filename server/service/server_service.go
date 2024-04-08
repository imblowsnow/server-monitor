package service

import (
	"server/dao"
	"server/entity/bo"
	"server/entity/do"
	"time"
)

var serverStateDao = dao.GetServerStateDao()
var serverInfoDao = dao.GetServerInfoDao()
var serverDao = dao.GetServerDao()
var serverFaultDao = dao.ServerFaultDao{}

type ServerService struct {
}

// 获取服务器列表（首页展示）
func (s *ServerService) GetIndexServerList() []bo.ServerBO {
	serverList := serverDao.GetServerList(true)
	var serverBOList []bo.ServerBO
	for _, server := range serverList {
		serverBOList = append(serverBOList, bo.ServerBO{
			ID:          server.ID,
			Name:        server.Name,
			Group:       server.Group,
			Version:     server.Version,
			Hide:        server.Hide,
			Index:       server.Index,
			CreatedTime: server.CreatedTime,
			UpdatedTime: server.UpdatedTime,
			Info:        serverInfoDao.GetServerInfoByServerID(server.ID),
			State:       serverStateDao.GetServerStateByServerID(server.ID),
		})
	}

	return serverBOList
}

// 获取服务每分钟统计状态
// func (s *ServerService) GetServerStatistics() map[string]interface{} {
// }
//
// 获取服务故障率 按天
func (s *ServerService) GetServerFaultTotal() []bo.ServerFaultTotalBO {
	// 获取前30天
	today := time.Now()

	times := []string{}
	minDate := today.AddDate(0, 0, -30).Format("2006-01-02")
	maxDate := today.AddDate(0, 0, 1).Format("2006-01-02")
	// 使用循环获取前30天的日期
	for i := 0; i <= 30; i++ {
		// 计算前i天的日期
		previousDate := today.AddDate(0, 0, -i)
		// 格式化日期为字符串
		date := previousDate.Format("2006-01-02")
		times = append(times, date)
	}

	total := make([]bo.ServerFaultTotalBO, 0)
	// 获取服务器列表
	serverList := serverDao.GetServerList(true)
	for _, server := range serverList {
		serverFaultTotal := bo.ServerFaultTotalBO{
			ServerID: server.ID,
			Name:     server.Name,
			Group:    server.Group,
			Items:    make([]*bo.ServerFaultTotalItemBO, 0),
		}
		serverFaultList := serverFaultDao.GetServerFaultListByTime(server.ID, minDate, maxDate)

		for _, startDate := range times {
			startDateTime, _ := time.Parse("2006-01-02", startDate)
			endDate := startDateTime.AddDate(0, 0, 1).Format("2006-01-02")
			endDateTime, _ := time.Parse("2006-01-02", endDate)

			totalTime := 24 * time.Hour
			faultTime := time.Duration(0)

			tmpServerFaultList := make([]do.ServerFault, 0)
			for _, fault := range serverFaultList {
				// 获取交集时间段
				faultStartTime, _ := time.Parse("2006-01-02", fault.StartTime.Format("2006-01-02"))
				var faultEndTime time.Time
				if fault.EndTime.IsZero() {
					faultEndTime = time.Now()
				} else {
					faultEndTime, _ = time.Parse("2006-01-02", fault.EndTime.Format("2006-01-02"))
				}

				// 获取交集时间段
				intersectionStart, intersectionEnd := intersection(startDateTime, endDateTime, faultStartTime, faultEndTime)

				if intersectionStart.IsZero() || intersectionEnd.IsZero() {
					continue
				}

				faultTime += time.Duration(fault.Duration) * time.Second
				tmpServerFaultList = append(tmpServerFaultList, fault)
			}

			serverFaultTotal.Items = append(serverFaultTotal.Items, &bo.ServerFaultTotalItemBO{
				Time:      startDate,
				TotalTime: totalTime,
				FaultTime: faultTime,
				FaultRate: float64(faultTime / totalTime),
				FaultList: tmpServerFaultList,
			})
		}

		total = append(total, serverFaultTotal)
	}

	return total
}

// 获取两个时间段的交集时间段
func intersection(start1, end1, start2, end2 time.Time) (time.Time, time.Time) {
	// 找到最晚的开始时间
	start := start1
	if start2.After(start1) {
		start = start2
	}

	// 找到最早的结束时间
	end := end1
	if end2.Before(end1) {
		end = end2
	}

	// 检查是否存在交集
	if start.After(end) {
		return time.Time{}, time.Time{} // 不存在交集
	}

	return start, end
}
