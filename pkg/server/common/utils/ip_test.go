package utils

import "testing"

func TestGetIp(t *testing.T) {
	ip := GetIp()
	t.Log(ip)
	if ip == "" {
		t.Error("获取ip失败")
	}
}
