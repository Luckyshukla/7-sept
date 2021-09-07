package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)
/*
func main()  {
	router:=gin.Default()

	router.GET("/user/:name", func (c *gin.Context)  {
		name:= c.Param("name")
		c.String(200, "Name %s", name)
	})
	router.Run()
}
*/

/*
func main()  {
	router := gin.Default()
	 router.GET("/user/:name/*action", func (r *gin.Context) {
		 name:= r.Param("name")
		 action :=r.Param("action")
		 message := name +"is"+action

		 r.String(200,message)
		 
	 })
	router.Run()
}
*/

// For each matched request Context will hold the route definition
/*

func  main()  {
	router:= gin.Default()
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(200, "%t", b)
	})
	router.Run()
}

*/




// Querystring parameters


/*
func main()  {
	router:= gin.Default()
	router.GET("/welcome", func (c *gin.Context) {
		FistName:= 	c.DefaultQuery("firstnsme", "Guest")
		Lastname:= c.Query("Lastname")

		c.String(200, "Hello %s %s", FistName, Lastname)
	})
	router.Run()
}
*/




//Multipart/Urlencoded Form


func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}


/*
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
*/


