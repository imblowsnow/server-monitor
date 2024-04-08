package model

type WebsocketMessageInit struct {
	// 验证秘钥
	Key string `json:"key"`
	// 服务器信息
	ServerInfo ServerInfo `json:"server_info"`
}
