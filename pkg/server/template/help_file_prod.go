//go:build !dev

package template

import (
	"embed"
	"github.com/gin-gonic/gin"
)

//go:embed backend/*
var backendTemplates embed.FS

//go:embed frontend/*
var frontedTemplates embed.FS

func autoResponseHttpFile(c *gin.Context, group, template, urlPath string) {
	filepath := group + "/" + template + urlPath

	var fs embed.FS
	if group == "backend" {
		fs = backendTemplates
	} else {
		fs = frontedTemplates
	}
	err := responseHttpFileFromFs(c, filepath, fs)
	if err != nil {
		c.Writer.WriteHeader(404)
		c.Writer.Write([]byte("404 Not Found"))
	}
}

func responseHttpFileFromFs(c *gin.Context, path string, fs embed.FS) error {
	// 读取文件内容
	fileBytes, err := fs.ReadFile(path)
	if err != nil {
		return err
	}
	responseHttpFileContent(c, path, fileBytes)

	return nil
}
