package template

import (
	"github.com/gin-gonic/gin"
)

func initFrontedTemplate(r *gin.Engine) {
	var indexHandle = func(c *gin.Context) {
		// 从frontedTemplates 读取 index.html 文件
		path := "/index.html"
		autoResponseHttpFile(c, "frontend", FrontedTemplate, path)
	}
	// 首页html路径
	r.GET("/", indexHandle)
	r.GET("/index.html", indexHandle)

	// 如果 以前缀 /v/ 开头的路径
	r.GET("/v/*path", func(c *gin.Context) {
		path := c.Param("path")
		autoResponseHttpFile(c, "frontend", FrontedTemplate, path)
	})
}
