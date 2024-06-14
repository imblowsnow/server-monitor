package bo

type ServerStateBO struct {
	// CPU占用
	CPU float64 `json:"cpu" gorm:"type:double"`
	// 内存占用
	MemUsed uint64 `json:"mem_used" gorm:"type:bigint"`
	// 交换分区占用
	SwapUsed uint64 `json:"swap_used" gorm:"type:bigint"`
	// 磁盘占用
	DiskUsed uint64 `json:"disk_used" gorm:"type:bigint"`
	// 网络入站流量
	NetInTransfer uint64 `json:"net_in_transfer" gorm:"type:bigint"`
	// 网络出站流量
	NetOutTransfer uint64 `json:"net_out_transfer" gorm:"type:bigint"`
	// 网络入站速度
	NetInSpeed uint64 `json:"net_in_speed" gorm:"type:bigint"`
	// 网络出站速度
	NetOutSpeed uint64 `json:"net_out_speed" gorm:"type:bigint"`
	// 开机时间
	Uptime uint64 `json:"uptime" gorm:"type:bigint"`
	// 系统负载
	Load1  float64 `json:"load1" gorm:"type:double"`
	Load5  float64 `json:"load5" gorm:"type:double"`
	Load15 float64 `json:"load15" gorm:"type:double"`
	// TCP连接数
	TcpConnCount int `json:"tcp_conn_count" gorm:"type:int"`
	// UDP连接数
	UdpConnCount int `json:"udp_conn_count" gorm:"type:int"`
	// 进程数
	ProcessCount int `json:"process_count" gorm:"type:int"`
}
