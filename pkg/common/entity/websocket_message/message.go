package websocket_message

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type WebsocketMessage struct {
	// 消息id
	Id string `json:"id"`
	// 消息类型
	MessageType int `json:"message_type"`
	// json数据
	Message string `json:"message"`
	// 发送时间
	Time time.Time `json:"time"`
}

func BuildWebsocketMessage(messageType int, message any) WebsocketMessage {
	jsonByte, _ := json.Marshal(message)
	websocketMessage := WebsocketMessage{
		Id:          uuid.New().String(),
		MessageType: messageType,
		Message:     string(jsonByte),
		Time:        time.Now(),
	}
	return websocketMessage
}

func ToMessageType[T any](websocketMessage WebsocketMessage, message T) (T, error) {
	err := json.Unmarshal([]byte(websocketMessage.Message), &message)
	if err != nil {
		fmt.Println("解析消息失败:", websocketMessage.Message, err)
		return message, err
	}
	return message, nil
}
