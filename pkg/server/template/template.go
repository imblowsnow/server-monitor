package template

import (
	"github.com/gin-gonic/gin"
)

// 默认后台模版
var BackendTemplate = "default"

// 默认前台模版
var FrontedTemplate = "default"

func InitTemplate(r *gin.Engine) {
	initFrontedTemplate(r)
	initBackendTemplate(r)
}
