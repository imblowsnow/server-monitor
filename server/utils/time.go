package utils

import "time"

// 获取两个时间段的交集时间段
func Intersection(start1, end1, start2, end2 time.Time) (time.Time, time.Time) {
	// 找到最晚的开始时间
	start := start1
	if start2.After(start1) {
		start = start2
	}

	// 找到最早的结束时间
	end := end1
	if end2.Before(end1) {
		end = end2
	}

	// 检查是否存在交集
	if start.After(end) || start.Equal(end) {
		return time.Time{}, time.Time{} // 不存在交集
	}

	return start, end
}
