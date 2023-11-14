package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/opengo-org/easy-netdisk/api"
)

func RunApplication(port string) {
	server := gin.Default()

	api.LoadApi(server)

	server.Run(port)
}
