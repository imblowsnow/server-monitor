package websocket_message

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type WebsocketMessageDTO struct {
	// 消息id
	Id string `json:"id"`
	// 消息类型
	MessageType int `json:"message_type"`
	// json数据
	Message string `json:"message"`
	// 发送时间
	Time time.Time `json:"time"`
	// 回复消息id
	ReplyMessageId string `json:"reply_message_id"`
}

func BuildWebsocketMessage(messageType int, message any) WebsocketMessageDTO {
	jsonByte, _ := json.Marshal(message)
	websocketMessage := WebsocketMessageDTO{
		Id:          uuid.New().String(),
		MessageType: messageType,
		Message:     string(jsonByte),
		Time:        time.Now(),
	}
	return websocketMessage
}

func BuildReplyWebsocketMessage(messageType int, message any, replyMessageId string) WebsocketMessageDTO {
	jsonByte, _ := json.Marshal(message)
	websocketMessage := WebsocketMessageDTO{
		Id:             uuid.New().String(),
		MessageType:    messageType,
		Message:        string(jsonByte),
		Time:           time.Now(),
		ReplyMessageId: replyMessageId,
	}
	return websocketMessage
}

func ToMessageType[T any](websocketMessage WebsocketMessageDTO, message T) (T, error) {
	err := json.Unmarshal([]byte(websocketMessage.Message), &message)
	if err != nil {
		fmt.Println("解析消息失败:", websocketMessage.Message, err)
		return message, err
	}
	return message, nil
}
