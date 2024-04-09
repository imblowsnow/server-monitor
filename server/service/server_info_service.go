package service

import "server/entity/do"

type ServerInfoService struct {
}

func (s ServerInfoService) SaveServerInfo(info do.ServerInfo) error {
	return serverInfoDao.SaveServerInfo(info)
}
