package main
import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
)
func main()  {
	gin.DisableConsoleColor()
	f,_:=os.Create("gin.log")
	gin.DefaultWriter() = io.MultiWriter("f")
	router:=gin.Default()
	router.GET("/ping",func (c *gin.Context)  {
		c.string(200, "Pong")

	})
	router.Run()
}