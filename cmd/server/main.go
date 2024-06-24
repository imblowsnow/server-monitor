package main

import (
	"flag"
	"server-monitor/pkg/server"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 22251, "port")
	flag.Parse()
	server.CreateServer(port)
}
