package model

type ServerInfo struct {
	HostName string `json:"HostName" gorm:"type:varchar(64)"`
	OS       string `json:"OS" gorm:"type:varchar(64)"`
	// 设备型号
	Platform       string `json:"Platform" gorm:"type:varchar(64)"`
	PlatformFamily string `json:"PlatformFamily" gorm:"type:varchar(64)"`
	// 设备版本
	PlatformVersion string `json:"PlatformVersion" gorm:"type:varchar(64)"`
	// CPU型号列表，多个CPU用逗号分隔
	CPU string `json:"CPU" gorm:"type:varchar(64)"`
	// 内存总大小
	MemTotal uint64 `json:"MemTotal" gorm:"type:bigint"`
	// 磁盘总大小
	DiskTotal uint64 `json:"DiskTotal" gorm:"type:bigint"`
	// 交换分区总大小
	SwapTotal      uint64 `json:"SwapTotal" gorm:"type:bigint"`
	Arch           string `json:"Arch" gorm:"type:varchar(64)"`
	Virtualization string `json:"Virtualization" gorm:"type:varchar(64)"`
	// 启动时间
	BootTime int `json:"BootTime"`
	// 客户端版本
	Version string `json:"Version" gorm:"type:varchar(64)"`
	// 获取系统时区
	Timezone string `json:"Timezone" gorm:"type:varchar(64)"`
}
