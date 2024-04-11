package ip

import "server/lib/ip/adapater"

const (
	IpApi = 1
)

func GetIpImpl(ip_type int) adapater.IIp {
	switch ip_type {
	case IpApi:
		return &adapater.IpApi{}
	}
	return nil
}
