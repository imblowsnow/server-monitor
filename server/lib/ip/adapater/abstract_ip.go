package adapater

import "server/entity/bo"

type AbstractIp struct {
	IIp
}

func (*AbstractIp) QueryIp(ip string) bo.QueryIpBO {
	panic("not implemented") // TODO: Implement
}

func (s *AbstractIp) QueryIps(ips []string) []bo.QueryIpBO {
	list := []bo.QueryIpBO{}
	for _, ip := range ips {
		list = append(list, s.QueryIp(ip))
	}
	return nil
}
