package do

import "time"

type ServerFault struct {
	ID uint `json:"id" gorm:"primarykey"`
	// 服务器ID
	ServerID uint `json:"server_id" gorm:"index"`
	// 故障开始时间
	StartTime time.Time `json:"start_time"`
	// 故障结束时间
	EndTime time.Time `json:"end_time" gorm:"default:null"`
	// 故障时长
	Duration int `json:"duration"`
}
