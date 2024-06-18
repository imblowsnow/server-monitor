package bo

type MonitorServerStatisticsBO struct {
	// 日期
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	// 在线率
	Rate int `json:"status"`

	Total  int64 `json:"total"`
	Online int64 `json:"online"`
}
