package service

import (
	"fmt"
	_ "server/lib"
	"testing"
)

func TestServerService_GetServerFaultTotal(t *testing.T) {
	serverService := ServerService{}
	total := serverService.GetServerFaultTotal()

	fmt.Println(total)
}
