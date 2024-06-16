//go:build !dev

package template

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
)

//go:embed backend/*
var backendTemplates embed.FS

func initBackendTemplate(r *gin.Engine) {
	fmt.Println("initBackendTemplate prod")
	r.GET("/admin/*path", func(c *gin.Context) {
		urlPath := c.Param("path")
		if urlPath == "/" {
			urlPath = "/index.html"
		}
		path := "backend/" + BackendTemplate + urlPath
		writeHttpFileFromFs(c, path, backendTemplates)
	})
}
