package template

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func initBackendTemplate(r *gin.Engine) {
	fmt.Println("initBackendTemplate dev")
	r.GET("/admin/*path", func(c *gin.Context) {
		urlPath := c.Param("path")
		if urlPath == "/" {
			urlPath = "/index.html"
		}
		autoResponseHttpFile(c, "backend", BackendTemplate, urlPath)
	})
}
