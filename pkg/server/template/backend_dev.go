//go:build dev

package template

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func initBackendTemplate(r *gin.Engine) {
	fmt.Println("initBackendTemplate dev")
	r.GET("/admin", func(c *gin.Context) {
		path := "backend/" + BackendTemplate + "/index.html"
		writeHttpFileFromLocal(c, path)
	})
	r.GET("/admin/*path", func(c *gin.Context) {
		path := "backend/" + BackendTemplate + c.Param("path")
		writeHttpFileFromLocal(c, path)
	})
}
