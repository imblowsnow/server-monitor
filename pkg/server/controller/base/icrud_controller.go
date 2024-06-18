package base

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/dal/dao"
)

type ICrudController[DO any, ID int | uint | uint32 | uint64 | string] interface {
	List(context *gin.Context) interface{}

	Page(context *gin.Context) interface{}

	Get(context *gin.Context) interface{}

	Create(context *gin.Context) interface{}

	Update(context *gin.Context) interface{}

	Delete(context *gin.Context) interface{}

	GetDao() dao.IBaseDao[DO, ID]
}
