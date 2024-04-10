//go:build !dev
// +build !dev

package template

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//go:embed admin
var adminFiles embed.FS

//go:embed front
var frontFiles embed.FS

func InitRoute(r *gin.Engine) {
	// 首页模板
	for _, path := range frontPaths {
		r.GET(path, func(c *gin.Context) {
			path := c.Request.URL.Path
			outTemplate(c, "front", frontTemplate, path, frontFiles)
		})
	}

	// 后台模板
	r.GET("/admin/*filepath", func(c *gin.Context) {
		path := c.Request.URL.Path
		// 删除 /admin前缀
		if strings.HasPrefix(path, "/admin") {
			path = path[6:]
		}
		outTemplate(c, "admin", adminTemplate, path, adminFiles)
	})
}

func outTemplate(c *gin.Context, prefix, template, path string, files embed.FS) {
	realTemplatePath := prefix + "/" + template + path
	realDefaultTemplatePath := prefix + "/default" + path
	// 判断文件是否存在
	if template != "default" {
		// 首页模板文件
		if path == "/" || path == "" {
			// 输出指定模板文件
			c.FileFromFS(realTemplatePath, http.FS(files))
			return
		}
		// 判断文件是否存在
		_, err := files.Open(realTemplatePath)
		if err == nil {
			// 输出指定模板文件
			c.FileFromFS(realTemplatePath, http.FS(files))
			return
		}
	}
	// 输出默认模板文件
	c.FileFromFS(realDefaultTemplatePath, http.FS(files))
}
