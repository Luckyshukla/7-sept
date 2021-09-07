package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)
func main() {
    
    gin.DisableConsoleColor()

    
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    // Use the following code if you need to write the logs to file and console at the same time.
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    router:= gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run(":8080")
}
