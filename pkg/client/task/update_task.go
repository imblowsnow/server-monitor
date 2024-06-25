package task

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"net/http"
	"server-monitor/pkg/client/constants"
)

func init() {
	fmt.Println("[定时任务]检测更新任务启动")

	go func() {
		CheckUpdate()

		c := cron.New()
		c.AddFunc("0 0 0 * * *", func() {
			fmt.Println("[定时任务]检测更新任务执行")

			// 检查是否是最新版本
			CheckUpdate()
		})
	}()
}

func getLatestVersion() string {
	// 获取最新版本
	githubUrl := "https://api.github.com/repos/imblowsnow/server-monitor/releases/latest"
	// get请求
	resp, err := http.Get(githubUrl)

	if err != nil {
		log.Fatalf("[更新检测任务] http.Get() failed with %s\n", err)
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("[更新检测任务] 请求失败", resp.StatusCode, string(body))
		return ""
	}

	// 转换为json
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("[更新检测任务] 获取失败", err)
		return ""
	}
	return data["tag_name"].(string)
}

func CheckUpdate() {
	// 获取最新版本
	latestVersion := getLatestVersion()
	if latestVersion == "" {
		return
	}

	if latestVersion != constants.ClientVersion {
		fmt.Println("[更新检测任务] 发现新版本", latestVersion)
		doUpdate()
	}
}
