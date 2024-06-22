package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/entity/vo"
	"server-monitor/pkg/server/common/utils"
	"server-monitor/pkg/server/config"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (controller UserController) Login(context *gin.Context) interface{} {
	var userLoginVo vo.UserLoginVo
	err := context.ShouldBindJSON(&userLoginVo)
	if err != nil {
		return fmt.Errorf("登录失败:%v", err)
	}
	if userLoginVo.Username != config.Config.AdminUsername ||
		userLoginVo.Password != config.Config.AdminPassword {
		return fmt.Errorf("登录失败:用户名或密码错误")
	}

	// 登录成功，生成jwt
	token, err := utils.GenerateToken(map[string]interface{}{
		"username": userLoginVo.Username,
	})
	if err != nil {
		fmt.Println("jwt生成失败", err)
		return fmt.Errorf("登录失败:%v", err)
	}
	return token
}
