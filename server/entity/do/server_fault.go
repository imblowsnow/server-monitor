package do

type ServerFault struct {
	ID uint `json:"id" gorm:"primarykey"`
	// 服务器ID
	ServerID uint `json:"server_id" gorm:"index"`
	// 故障开始时间
	StartTime int64 `json:"start_time"`
	// 故障结束时间
	EndTime int64 `json:"end_time" gorm:"default:null"`
}
