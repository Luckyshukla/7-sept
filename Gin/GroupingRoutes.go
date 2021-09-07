package main 

import (
	"fmt"
	//"net/http"
	"github.com/gin-gonic/gin"
)
 type TestModel struct{
	 Id int `json:"id" binding:"required"`
	 name string `json:"name" binding:"required`
 }
 func main() {
	r:=gin.Default()
	v2:= r.Group("v2")
	v1:=r.Group("v1")
	{
	user:= v1.Group("user")
	{
		user.GET("",func(c *gin.Context)  {
			c.JSON(200,gin.H{
				"data":[]TestModel{
					TestModel{
						Id:1,
						name:"User1",
					},
					TestModel{
						Id:1,
						name:"User2", 
					},
				},
			})
		})
		user.GET(":Id",func(c *gin.Context){
			var id= c.Param("Id")
			fmt.Println("id recive id ", id)
			c.JSON(200,gin.H{
				"Data":TestModel{
					Id:1,
					name:"User1",
				},
			})
		})
	}
	Product:= v1.Group("Product")
	{
		Product.GET("/:Id",func (c *gin.Context)  {
	 	c.JSON(200,gin.H{
	 		"Product":[]TestModel{
				TestModel{
	 					Id:1,
	 					name:"Product1",
	 				},
	 				TestModel{
	 					Id:1,
	 					name:"Product1", 
	 				},
	 			},
	 		})
	 	} )
	}
}
//2nd group
user:= v2.Group("Product")
	{
		user.GET("",func(c *gin.Context)  {
			c.JSON(200,gin.H{
				"datav2":[]TestModel{
					TestModel{
						Id:1,
						name:"User1v2",
					},
					TestModel{
						Id:1,
						name:"User2v2", 
					},
				},
			})
		})
		user.GET(":Id",func(c *gin.Context){
			var id= c.Param("Id")
			fmt.Println("id recive id ", id)
			c.JSON(200,gin.H{
				"Data":TestModel{
					Id:1,
					name:"User1v2",
				},
			})
		})
	Product:= v2.Group("Product")
	{
		Product.GET("/:Id",func (c *gin.Context)  {
	 	c.JSON(200,gin.H{
	 		"Product":[]TestModel{
				TestModel{
	 					Id:1,
	 					name:"Product1v2",
	 				},
	 				TestModel{
	 					Id:1,
	 					name:"Productv2", 
	 				},
	 			},
	 		})
	 	} )
	}
}

	r.Run()
}