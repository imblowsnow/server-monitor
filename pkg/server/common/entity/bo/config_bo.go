package bo

type ConfigBo struct {
	// 秘钥
	Secret string `json:"-"`
	// 站点名称
	SiteName string `json:"site_name"`

	// 前端模版
	FrontendTemplate string `json:"frontend_template"`
	// 后端模版
	BackendTemplate string `json:"backend_template"`

	// 管理员账号
	AdminUsername string `json:"admin_username"`
	// 管理员密码
	AdminPassword string `json:"admin_password"`
}
