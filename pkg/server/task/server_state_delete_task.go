package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"server-monitor/pkg/server/dal/dao"
)

func init() {
	// 自动删除30天前的任务
	fmt.Println("初始化自动删除30天前的服务记录数据任务")

	c := cron.New()

	serverStateDao := dao.NewServerStateDao()
	// 每天凌晨1点执行
	c.AddFunc("0 0 1 * * *", func() {
		fmt.Println("自动删除30天前的服务记录数据任务")
		// 删除30天前的服务记录数据
		serverStateDao.DeleteThirtyDaysAgo()
	})
	c.Start()
}
