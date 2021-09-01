package main 

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	
	fmt.Println("Hello Kushagra_Maple_Labs")
	r := gin.Default()

// ---------------simple get the request---------------

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello Kushagra_Maple_Labs",
	// 	})
	// })

// ------------------getting the username using get request ---------------

	// r.GET("/user/:name", func(c *gin.Context) {
    //     name := c.Param("name")
    //     c.String(200, "Hello %s", name)
    // })

//  --------------- SUBROUTING IN GIN-GOLANG ---------------

	// v1 := r.Group("/api/v1")
    // {
    //     auth := v1.Group("/auth")
    //     {
    //         auth.GET("/", func(c *gin.Context) {
    //             c.JSON(200, gin.H{
    //                 "v1-auth": true,
    //             })
    //         })
    //     }

	// 	name := v1.Group("/name")
	// 	{
	// 		name.GET("/", func(c *gin.Context) {
	// 			c.JSON(200, gin.H{
	// 				"v1-name":true,
	// 			})
	// 		})
	// 	}

    // }

// STACKOVERFLOW PROBLEM SOLUTION -> SUBROUTING In GOLANG

	resources := r.Group("/resources")
	{
		resources.GET("/:id",func(c*gin.Context){
			res_id := c.Param("id")
			c.JSON(200,gin.H{
				"route":true,
				"resources_id": res_id,
			})
		})

		id := resources.Group("/:id")
		{
			subresource := id.Group("/subresource")
			{
				subresource.GET("/:newid", func(c *gin.Context) {
					c.JSON(200, gin.H{
					"route": true,
					"subresource": "i am working route inside new_id",
				})
			})
			}	
		
		}
	}
    r.Run()
}