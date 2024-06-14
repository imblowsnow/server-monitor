package do

import (
	"server-monitor/pkg/common/entity"
	"time"
)

type ServerInfoDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 服务器ID
	ServerId uint `json:"server_id"`

	// 服务器信息
	entity.ServerInfo `gorm:"embedded"`

	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime"`
}
