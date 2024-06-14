package utils

import (
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	serverInfo := GetServerInfo()
	t.Log(serverInfo)
}

func TestGetServerState(t *testing.T) {
	serverStat := GetServerState()
	t.Log(serverStat)
}
