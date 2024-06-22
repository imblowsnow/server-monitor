package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/utils"
)

var whiteList = []string{"/admin-api/v1/user/login"}

func AdminAuthMiddle(c *gin.Context) {
	// 获取path
	path := c.Request.URL.Path
	// 判断是否在白名单中
	for _, v := range whiteList {
		if path == v {
			c.Next()
			return
		}
	}
	// 验证token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"message": "未登录"})
		c.Abort()
		return
	}
	// 验证token是否有效
	account, err := utils.ParseToken(token)
	if err != nil {
		fmt.Println("token解析失败", err)
		c.JSON(401, gin.H{"message": "未登录"})
		c.Abort()
		return
	}
	// 将account存入上下文
	c.Set("account", account)
	c.Next()
}
