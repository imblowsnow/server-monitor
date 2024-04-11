package service

import (
	"fmt"
	_ "server/lib"
	"testing"
)

func TestServerService_GetServerFaultTotal(t *testing.T) {
	serverService := ServerService{}
	total := serverService.GetServerStatus()

	fmt.Println(total)
}
