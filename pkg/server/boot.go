package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"server-monitor/pkg/common/inner_websocket"
	// 用于注册定时任务
	_ "server-monitor/pkg/server/task"
	"server-monitor/pkg/server/websocket/event"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreateServer(port int) {
	// 启动一个 http 服务
	r := gin.Default()

	// 启动 服务端 websocket
	// 处理 websocket 服务
	r.GET("/ws/client", func(c *gin.Context) {
		handleUpgradeWebsocket(c.Writer, c.Request)
	})

	// TODO 注册接口路由

	r.Run(":" + strconv.Itoa(port))
}

func handleUpgradeWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("客户端ws链接失败", err)
		return
	}
	// 循环处理 websocket 事件
	inner_websocket.LoopWebsocket(conn, &event.ServerWebsocketEvent{})
}
