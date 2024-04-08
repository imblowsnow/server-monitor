package do

import (
	"common/model"
	"time"
)

type ServerInfo struct {
	ID       uint `json:"id" gorm:"primarykey"`
	ServerID uint `json:"server_id" gorm:"uniqueIndex"`
	// 关联服务器信息
	Info model.ServerInfo `json:"info" gorm:"embedded"`
	// 创建时间
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime"`
	// 更新时间
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime"`
}
