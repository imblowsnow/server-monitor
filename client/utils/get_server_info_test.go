package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetServerInfo(t *testing.T) {

	// 测试获取服务器状态
	serverInfo := GetServerInfo()

	bytes, _ := json.Marshal(serverInfo)
	// 转换为string

	fmt.Println("serverInfo", string(bytes))
}
