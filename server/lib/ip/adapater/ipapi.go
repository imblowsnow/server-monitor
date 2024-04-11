package adapater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/entity/bo"
)

type IpApi struct {
	AbstractIp
}

func (s *IpApi) QueryIp(ip string) (bo.QueryIpBO, error) {
	// 请求 http://ip-api.com/json/{ip}
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return bo.QueryIpBO{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bo.QueryIpBO{}, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return bo.QueryIpBO{}, err
	}
	if m["status"] == "fail" {
		return bo.QueryIpBO{}, fmt.Errorf(m["message"].(string))
	}
	queryIpBo := bo.QueryIpBO{
		Ip:          m["query"].(string),
		CountryCode: m["countryCode"].(string),
		Isp:         m["isp"].(string),
	}

	return queryIpBo, nil
}

func (s *IpApi) QueryIps(ips []string) ([]bo.QueryIpBO, error) {
	// Convert the ips slice to JSON
	ipJson, err := json.Marshal(ips)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://ip-api.com/batch", "application/json", bytes.NewBuffer(ipJson))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	datas := make([]map[string]interface{}, 0)
	err = json.Unmarshal(body, &datas)
	if err != nil {
		return nil, err
	}
	queryIpBos := make([]bo.QueryIpBO, 0)
	for _, m := range datas {
		if m["status"] == "fail" {
			continue
		}
		queryIpBo := bo.QueryIpBO{
			Ip:          m["query"].(string),
			CountryCode: m["countryCode"].(string),
			Isp:         m["isp"].(string),
		}
		queryIpBos = append(queryIpBos, queryIpBo)
	}

	return queryIpBos, nil
}
