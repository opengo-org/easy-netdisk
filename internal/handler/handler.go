package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Welcom come to my website"})
}

func HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "admin" && password == "520.Linux" {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Login Successfully!"})
	}
}
