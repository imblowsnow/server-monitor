package bo

import "server-monitor/pkg/common/time"

type MonitorServerBO struct {
	ServerId       uint      `json:"server_id"`
	ServerName     string    `json:"server_name"`
	ServerStatus   int       `json:"server_status"`
	LastOnlineTime time.Time `json:"last_online_time"`
	// 在线率
	OnlineRate int `json:"online_rate"`
	// 在线统计列表
	OnlineStatistics []*MonitorServerStatisticsBO `json:"online_statistics"`
}
