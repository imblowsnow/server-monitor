//go:build dev

package template

import (
	"github.com/gin-gonic/gin"
	"os"
)

func autoResponseHttpFile(c *gin.Context, group, template, urlPath string) {
	filepath := group + "/" + template + urlPath
	err := responseHttpFileFromLocal(c, filepath)
	if err != nil {
		c.Writer.WriteHeader(404)
		c.Writer.Write([]byte("404 Not Found"))
	}
}

func responseHttpFileFromLocal(c *gin.Context, path string) error {
	// 当前路径
	pwd, _ := os.Getwd()
	// 读取文件内容
	fileBytes, err := os.ReadFile(pwd + "/pkg/server/template/" + path)
	if err != nil {
		return err
	}
	responseHttpFileContent(c, path, fileBytes)

	return nil
}
