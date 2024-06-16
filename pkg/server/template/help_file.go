package template

import (
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"path"
)

func responseHttpFileContent(c *gin.Context, filePath string, content []byte) {
	// 根据文件路径获取content-type
	contentType := mime.TypeByExtension(path.Ext(filePath))

	c.Writer.Header().Set("Content-Type", contentType)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(content)
}
