package model

import "time"

// 定义常量作为枚举
const (
	// 初始化消息
	MessageTypeInit  = 100
	MessageTypeState = 200
)

type WebsocketMessage struct {
	// 消息id
	Id string `json:"id"`
	// 消息类型
	MessageType int `json:"message_type"`
	// json数据
	Message any `json:"message"`
	// 发送时间
	Time time.Time `json:"time"`
}
