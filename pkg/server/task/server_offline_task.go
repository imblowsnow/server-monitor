package task

import (
	"fmt"
	"server-monitor/pkg/server/dal/dao"
	"time"
)

func init() {
	fmt.Println("[定时任务]服务端离线检测任务启动")
	// 3分钟检测一次服务端是否在线， 主要用于客户端异常掉线（服务端未获取到关闭事件）/ 服务端异常掉线重启的情况
	// TODO 查询数据库，获取服务端最后一次心跳时间是否超过3分钟，且未通知过
	serverDao := dao.NewServerDao()
	go func() {
		ticket := time.Tick(time.Minute * 1)
		for {
			select {
			case <-ticket:
				fmt.Println("[定时任务]检测服务端是否在线")
				serverDao.CheckServerOffline()
			}
		}
	}()
}
