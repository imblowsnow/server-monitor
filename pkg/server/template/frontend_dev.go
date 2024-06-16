//go:build dev

package template

import (
	"github.com/gin-gonic/gin"
)

func initFrontedTemplate(r *gin.Engine) {
	// 首页html路径
	r.GET("/", func(c *gin.Context) {
		// 从frontedTemplates 读取 index.html 文件
		path := "frontend/" + FrontedTemplate + "/index.html"
		writeHttpFileFromLocal(c, path)
	})
	r.GET("/index.html", func(c *gin.Context) {
		// 从frontedTemplates 读取 index.html 文件
		path := "frontend/" + FrontedTemplate + "/index.html"
		writeHttpFileFromLocal(c, path)
	})
	// 如果 以前缀 /v/ 开头的路径
	r.GET("/v/*path", func(c *gin.Context) {
		path := "frontend/" + FrontedTemplate + c.Param("path")
		writeHttpFileFromLocal(c, path)
	})
}
