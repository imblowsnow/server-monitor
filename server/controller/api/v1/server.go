package v1

import (
	"github.com/gin-gonic/gin"
	"server/entity/result"
	"server/service"
)

var serverService = service.ServerService{}

// 获取服务器信息
func GetServerInfo(c *gin.Context) {
	data := serverService.GetIndexServerList()
	c.JSON(200, result.Success(data))
}

// 获取服务器统计
func GetServerStatus(c *gin.Context) {
	data := serverService.GetServerStatus()
	c.JSON(200, result.Success(data))
}

func GetGroups(c *gin.Context) {
	data := serverService.GetGroups()
	c.JSON(200, result.Success(data))
}
