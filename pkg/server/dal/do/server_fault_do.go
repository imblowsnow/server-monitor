package do

import "server-monitor/pkg/common/time"

type ServerFaultDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 服务器ID
	ServerId uint `json:"server_id"`
	// 异常开始时间
	StartTime time.Time `json:"start_time"`
	// 异常结束时间
	EndTime time.Time `json:"end_time"`
	// 持续时间
	Duration uint64 `json:"duration"`
	// 备注
	Remark string `json:"remark" gorm:"type:varchar(64)"`
}
