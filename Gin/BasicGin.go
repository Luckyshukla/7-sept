package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go get -u gorm.io/gorm"
	"go get -u gorm.io/driver/sqlite"

)
/*
func Getting(R *gin.Context)  {
	R.JSON(200,gin.H{
		"name":"Lucky",
	})
}



func main()  {
	router := gin.Default()
	router.GET("/",Getting)	
	router.Run()
}
*/
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
