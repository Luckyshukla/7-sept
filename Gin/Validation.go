package main
import(
	
	"net/http"
	"github.com/gin-gonic/gin"
)
type Login struct{
	name string `form:"user" json:"user" binding required`
	password string `form:"password" json:"password" binding required`
} 
func main() {
	router := gin.Default()

	
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}