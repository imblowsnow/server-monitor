package boot

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	_ "server/lib"
	"server/template"
	"server/websocket_event"
	"strconv"
)

func Run(httpPort int) {
	// 启动一个 http 服务
	r := gin.Default()

	// 启动 服务端 websocket
	// 处理 websocket 服务
	r.GET("/ws/client", func(c *gin.Context) {
		handleUpgradeWebsocket(&websocket_event.ServerWebSocketEvent{
			FrontNotifyEvent: &websocket_event.FrontNotifyEvent{},
		}, c.Writer, c.Request)
	})

	// 启动web websocket端
	r.GET("/ws/web", func(c *gin.Context) {
		handleUpgradeWebsocket(&websocket_event.FrontWebSocketEvent{}, c.Writer, c.Request)
	})

	// 初始化控制器路由
	controller.InitRoute(r)

	// 初始化模板路由
	template.InitRoute(r)

	r.Run(":" + strconv.Itoa(httpPort))
}
