package boot

import (
	"github.com/gin-gonic/gin"
	"server/websocket_event"
	"strconv"
)

func Run(httpPort int) {
	// 启动一个 http 服务
	r := gin.Default()

	// 启动 服务端 websocket
	// 处理 websocket 服务
	r.GET("/client", func(c *gin.Context) {
		handleUpgradeWebsocket(&websocket_event.ServerWebSocketEvent{}, c.Writer, c.Request)
	})

	// 启动web websocket端
	r.GET("/web", func(c *gin.Context) {
		handleUpgradeWebsocket(&websocket_event.WebWebSocketEvent{}, c.Writer, c.Request)
	})

	r.Run(":" + strconv.Itoa(httpPort))
}
