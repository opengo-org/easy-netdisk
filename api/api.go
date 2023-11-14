package api

import (
	"github.com/gin-gonic/gin"
	"github.com/opengo-org/easy-netdisk/internal/handler"
)

func LoadApi(server *gin.Engine) {
	server.GET("/", handler.HandleRoot)
	user_group := server.Group("/user")
	{
		user_group.POST("/login", handler.HandleLogin)
	}
}
