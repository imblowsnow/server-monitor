//go:build dev
// +build dev

package template

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InitRoute(r *gin.Engine) {
	fmt.Println("InitTemplateRoute dev")
	// 首页模板
	for _, path := range frontPaths {
		r.GET(path, func(c *gin.Context) {
			path := c.Request.URL.Path
			outTemplate(c, "front", frontTemplate, path, embed.FS{})
		})
	}

	// 后台模板
	r.GET("/admin/*filepath", func(c *gin.Context) {
		path := c.Request.URL.Path
		// 删除 /admin前缀
		if strings.HasPrefix(path, "/admin") {
			path = path[6:]
		}
		outTemplate(c, "admin", adminTemplate, path, embed.FS{})
	})
}

func outTemplate(c *gin.Context, prefix, template, path string, files embed.FS) {
	realTemplatePath := "template/" + prefix + "/" + template + path
	realDefaultTemplatePath := "template/" + prefix + "/default" + path
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
	c.File(realDefaultTemplatePath)
}
