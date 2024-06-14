package enum

// 定义常量作为枚举
const (
	/*-------------- 客户端发送服务端消息 --------------*/
	// 服务器初始化消息
	MessageServerInit = 1000
	// 服务器关闭消息
	MessageServerClose = 1001

	// 服务器状态消息
	MessageServerStat = 1002

	/*-------------- 服务端消息 --------------*/
	ServerMessageInitSuccess = 2000
)
