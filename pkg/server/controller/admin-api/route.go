package admin_api

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {
	adminApiV1Group := r.Group("/admin-api/v1")
	adminApiV1Group.GET("/server/list", Ping)
}
