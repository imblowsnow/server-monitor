package adapater

import "server/entity/bo"

type IIp interface {
	QueryIp(ip string) (bo.QueryIpBO, error)
	QueryIps(ips []string) ([]bo.QueryIpBO, error)
}
