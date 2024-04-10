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

// 获取服务器故障统计
func GetServerFaultTotal(c *gin.Context) {
	data := serverService.GetServerFaultTotal()
	c.JSON(200, result.Success(data))
}
