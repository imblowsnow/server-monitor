package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetServerState(t *testing.T) {

	// 测试获取服务器状态
	serverState := GetServerState()

	bytes, _ := json.Marshal(serverState)
	// 转换为string

	fmt.Println("serverState", string(bytes))
}
