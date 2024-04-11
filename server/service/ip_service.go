package service

import (
	"server/entity/bo"
	"server/lib/ip"
	"server/lib/ip/adapater"
)

type IpService struct {
}

func (s *IpService) QueryIp(query string) bo.QueryIpBO {
	queryIp, err := s.GetIPImpl(ip.IpApi).QueryIp(query)
	if err != nil {
		return bo.QueryIpBO{}
	}
	return queryIp
}
func (s *IpService) QueryIps(ips []string) []bo.QueryIpBO {
	queryIps, err := s.GetIPImpl(ip.IpApi).QueryIps(ips)
	if err != nil {
		return nil
	}
	return queryIps
}

func (s *IpService) GetIPImpl(ip_type int) adapater.IIp {
	return ip.GetIpImpl(ip_type)
}
