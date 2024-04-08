package main

import "server/boot"

func main() {
	httpPort := 8999

	// 启动服务
	boot.Run(httpPort)
}
