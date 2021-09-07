package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.BindQuery(&person) == nil {
		log.Println("====== Only Bind Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	} else {
		fmt.Println("please pass parameter name and address")
	}
	//c.String(200, "Success")
}
