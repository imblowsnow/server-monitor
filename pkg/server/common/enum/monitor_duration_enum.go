package enum

import "time"

type MonitorDurationEnum struct {
	Type     int
	Duration time.Duration
	Num      int
}

// 单位 分
var MONITOR_DURATION_MINUTE = MonitorDurationEnum{Type: 1, Duration: time.Minute, Num: 60}

// 单位 小时
var MONITOR_DURATION_HOUR = MonitorDurationEnum{Type: 2, Duration: time.Hour, Num: 24}

// 单位 天
var MONITOR_DURATION_DAY = MonitorDurationEnum{Type: 3, Duration: time.Hour * 24, Num: 30}

// 监控时长-月
var MONITOR_DURATION_MONTH = MonitorDurationEnum{Type: 4, Duration: time.Hour * 24 * 30, Num: 12}

func FindMonitorDurationEnumByType(typeI int) MonitorDurationEnum {
	switch typeI {
	case MONITOR_DURATION_MINUTE.Type:
		return MONITOR_DURATION_MINUTE
	case MONITOR_DURATION_HOUR.Type:
		return MONITOR_DURATION_HOUR
	case MONITOR_DURATION_DAY.Type:
		return MONITOR_DURATION_DAY
	case MONITOR_DURATION_MONTH.Type:
		return MONITOR_DURATION_MONTH
	default:
		return MONITOR_DURATION_MINUTE
	}
}
