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
	Message any `json:"message"`
	// 发送时间
	Time time.Time `json:"time"`
}

func BuildWebsocketMessage(messageType int, message any) WebsocketMessage {
	websocketMessage := WebsocketMessage{
		Id:          uuid.New().String(),
		MessageType: messageType,
		Message:     message,
		Time:        time.Now(),
	}
	return websocketMessage
}

func ToMessageType[T any](websocketMessage WebsocketMessage, message T) (T, error) {
	jsonByte, err := json.Marshal(websocketMessage.Message)
	if err != nil {
		fmt.Println("解析消息失败:", err)
		return message, err
	}
	err = json.Unmarshal(jsonByte, &message)
	if err != nil {
		fmt.Println("解析消息失败:", websocketMessage.Message, err)
		return message, err
	}
	return message, nil
}
