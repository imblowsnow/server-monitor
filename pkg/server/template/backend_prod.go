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

	r.GET("/admin", func(c *gin.Context) {
		path := "backend/" + BackendTemplate + "/index.html"
		writeHttpFileFromFs(c, path, backendTemplates)
	})
	r.GET("/admin/*path", func(c *gin.Context) {
		path := "backend/" + BackendTemplate + c.Param("path")
		writeHttpFileFromFs(c, path, backendTemplates)
	})
}
