package template

import (
	"embed"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// 默认后台模版
var BackendTemplate = "default"

// 默认前台模版
var FrontedTemplate = "default"

func InitTemplate(r *gin.Engine) {
	initFrontedTemplate(r)
	initBackendTemplate(r)
}

func writeHttpFileFromFs(c *gin.Context, path string, fs embed.FS) {
	// 读取文件内容
	fileBytes, err := fs.ReadFile(path)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte("404 Not Found"))
		return
	}

	writeHttpFile(c, fileBytes)
}

func writeHttpFileFromLocal(c *gin.Context, path string) {
	// 当前路径
	pwd, _ := os.Getwd()
	// 读取文件内容
	fileBytes, err := os.ReadFile(pwd + "/pkg/server/template/" + path)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte("404 Not Found"))
		return
	}
	writeHttpFile(c, fileBytes)
}

func writeHttpFile(c *gin.Context, content []byte) {
	// 根据文件路径获取content-type
	mime := mimetype.Detect(content)
	c.Writer.Header().Set("Content-Type", mime.String())
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(content)
}
