package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetIp() string {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Println("获取ip失败", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("获取ip失败", resp.StatusCode, string(body))
		return ""
	}
	// 转换为json
	var ipInfo map[string]interface{}
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		fmt.Println("获取ip失败", err)
		return ""
	}
	return ipInfo["query"].(string)
}
