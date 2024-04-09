package service

import (
	"server/entity/do"
)

type ServerStateService struct {
}

func (s ServerStateService) AddServerState(state do.ServerState) error {
	return serverStateDao.AddServerState(state)
}
