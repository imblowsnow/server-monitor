package do

import "server-monitor/pkg/common/time"

type NotifyGroupChannelDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 通知组ID
	NotifyGroupId uint `json:"notify_group_id"`
	// 通知渠道ID
	NotifyChannelId uint `json:"notify_channel_id"`
	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;autoUpdateTime"`
}
