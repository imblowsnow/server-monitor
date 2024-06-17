package enum

type ServerStatus int

const (
	// 在线
	ServerStatusOnline ServerStatus = 1
	// 离线
	ServerStatusOffline ServerStatus = 0
)
