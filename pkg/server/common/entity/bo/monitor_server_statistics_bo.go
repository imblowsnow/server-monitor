package bo

type MonitorServerStatisticsBO struct {
	// 日期
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	// 在线率
	Rate int `json:"status"`
}
