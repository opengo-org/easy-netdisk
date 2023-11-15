package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opengo-org/easy-netdisk/internal/config"
)

func HandleRoot(c *gin.Context) {
	db := config.GetDatabase()
	if db != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Database connected successfully"})
	}
}

func HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "admin" && password == "520.Linux" {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Login Successfully!"})
	}
}
