package do

import (
	"time"
)

type NotifyChannelDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(64)"`
	// 配置
	Config string `json:"config" gorm:"type:text"`
	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;autoUpdateTime"`
}
