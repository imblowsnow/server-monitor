package model

type ServerState struct {
	// CPU占用
	CPU float64 `json:"CPU" gorm:"type:double"`
	// 内存占用
	MemUsed uint64 `json:"MemUsed" gorm:"type:bigint"`
	// 交换分区占用
	SwapUsed uint64 `json:"SwapUsed" gorm:"type:bigint"`
	// 磁盘占用
	DiskUsed uint64 `json:"DiskUsed" gorm:"type:bigint"`
	// 网络入站流量
	NetInTransfer uint64 `json:"NetInTransfer" gorm:"type:bigint"`
	// 网络出站流量
	NetOutTransfer uint64 `json:"NetOutTransfer" gorm:"type:bigint"`
	// 网络入站速度
	NetInSpeed uint64 `json:"NetInSpeed" gorm:"type:bigint"`
	// 网络出站速度
	NetOutSpeed uint64 `json:"NetOutSpeed" gorm:"type:bigint"`
	// 开机时间
	Uptime uint64 `json:"Uptime" gorm:"type:bigint"`
	// 系统负载
	Load1  float64 `json:"Load1" gorm:"type:double"`
	Load5  float64 `json:"Load5" gorm:"type:double"`
	Load15 float64 `json:"Load15" gorm:"type:double"`
	// TCP连接数
	TcpConnCount int `json:"TcpConnCount" gorm:"type:int"`
	// UDP连接数
	UdpConnCount int `json:"UdpConnCount" gorm:"type:int"`
	// 进程数
	ProcessCount int `json:"ProcessCount" gorm:"type:int"`
}
