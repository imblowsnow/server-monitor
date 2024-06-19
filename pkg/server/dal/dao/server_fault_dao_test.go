package dao

import (
	"fmt"
	time2 "server-monitor/pkg/common/time"
	"testing"
	"time"
)

func TestServerFaultDao_RecoverFault(t *testing.T) {
	dao := NewServerFaultDao()
	dao.RecordFault(1, "测试故障")
	// 延迟 10 秒
	time.Sleep(20 * time.Second)

	dao.RecoverFault(1)
}

func TestTime(t *testing.T) {
	startTimeV, _ := time.Parse("2006-01-02 15:04:05", "2024-06-19 13:55:12")
	startTime := time2.Time(startTimeV)

	endTimeV, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	endTime := time2.Time(endTimeV)

	fmt.Println(time.Time(endTime).Sub(time.Time(startTime)).Seconds())
}
