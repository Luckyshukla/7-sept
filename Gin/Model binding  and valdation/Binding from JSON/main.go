package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	name     string `json:"user"  binding:"required"`
	password string `json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var json user
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.name != "manu" || json.password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.Run()

}
