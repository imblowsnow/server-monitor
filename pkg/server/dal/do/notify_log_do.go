package do

import "server-monitor/pkg/common/time"

// NotifyLogDO 通知日志
type NotifyLogDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 通知组ID
	NotifyGroupId uint `json:"notify_group_id"`
	// 通知渠道ID
	NotifyChannelId uint `json:"notify_channel_id"`
	// 通知状态 1.成功 0.失败
	Status int `json:"status"`
	// 通知结果
	Result string `json:"result" gorm:"type:text"`
	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;autoUpdateTime"`
}
