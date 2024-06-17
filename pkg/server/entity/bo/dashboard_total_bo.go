package bo

type DashboardTotalBO struct {
	// 在线数量
	OnlineNum int64 `json:"onlineNum"`
	// 离线数量
	OfflineNum int64 `json:"offlineNum"`
}
