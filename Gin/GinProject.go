package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"log"
)


// func getting()  {
	
// }

// func posting(c *gin.Context)  {
// 	c.JSON(200, gin.H{
// 		"ID":0,
// 		"Employee":"Lucky",
// 		"email":"luckyshuklakunda@gmail.com",
// 		"age":21,
// 	})
// }
// func putting()  {
	
// }
// func deleting()  {
	
// }
// func patching()  {
	
// }
// func head()  {
	
// }
// func options()  {
	
func Multipart(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")	
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func Querystring(c *gin.Context)  {
	name:= c.Query("name")
	age:= c.Query("age")

	c.JSON(200,gin.H{
		"name": name,
		"age": age,
	})
}

func PathParameters(c *gin.Context) {
	name:=c.Param("name")
	age:= c.Param("age")

	c.JSON(200,gin.H{
		"name": name,
		"age": age,
	})
	
}
func UploadSinleFile(c *gin.Context) {
	// single file
	file, _ := c.FormFile("dst.txt")
	log.Println(file.Filename)
	c.SaveUploadedFile(file, "/")
}

// func MultiplefileFile(c *gin.Context)  {

// 	form, _ := c.MultipartForm()
// 	files := form.File["upload[]"]

// 	for _, file := range files {

// 		log.Println(file.Filename)

// 		// Upload the file to specific dst.
// 		c.SaveUploadedFile(file, dst)
// 	}
// }



func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/multipart",Multipart)
	router.GET("/Query", Querystring)
	router.GET("/user/:name/:age",PathParameters)
	router.POST("/SingleFileUpload", UploadSinleFile)
	//router.POST("/MultiplefileFile", MultiplefileFile)
	
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}

