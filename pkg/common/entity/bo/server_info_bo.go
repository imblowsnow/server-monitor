package bo

// ServerInfoBO 服务器信息
type ServerInfoBO struct {
	// 主机名
	Hostname string `json:"host_name" gorm:"type:varchar(64)"`
	// 系统版本
	OS string `json:"os" gorm:"type:varchar(64)"`
	// 设备型号
	Platform string `json:"platform" gorm:"type:varchar(64)"`

	PlatformFamily string `json:"platform_family" gorm:"type:varchar(64)"`
	// 设备版本
	PlatformVersion string `json:"platform_version" gorm:"type:varchar(64)"`
	// CPU型号列表，多个CPU用逗号分隔
	CPU string `json:"cpu" gorm:"type:varchar(64)"`
	// 内存总大小
	MemTotal uint64 `json:"mem_total" gorm:"type:bigint"`
	// 磁盘总大小
	DiskTotal uint64 `json:"disk_total" gorm:"type:bigint"`
	// 交换分区总大小
	SwapTotal uint64 `json:"swap_total" gorm:"type:bigint"`
	// 架构
	KernelArch string `json:"kernel_arch" gorm:"type:varchar(64)"`
	// 内核版本
	KernelVersion string `json:"kernel_version" gorm:"type:varchar(64)"`
	// 虚拟化
	Virtualization string `json:"virtualization" gorm:"type:varchar(64)"`
	// 启动时间
	BootTime uint64 `json:"boot_time"`
	// 客户端版本
	Version string `json:"version" gorm:"type:varchar(64)"`
}
