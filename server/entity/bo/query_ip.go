package bo

type QueryIpBO struct {
	Ip          string `json:"ip"`
	CountryCode string `json:"country_code"`
	Isp         string `json:"isp"`
}
