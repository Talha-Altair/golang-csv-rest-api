package main

import (

	"fmt"

	"github.com/gin-gonic/gin"
	
	)

func main() {

	r := gin.Default()

	data := read()

	fmt.Println("Starting server", data)

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"ping" : "pong",
		})
	})

	r.GET("/:city", func(c *gin.Context) {

		city := c.Param("city")

		fmt.Println("City:", city)

		flag := false

		for _, v := range data {

			fmt.Println("City:", v.city)

			if v.city == city {

			flag = true

			res := gin.H{
				"City": v.city,
				"State": v.state,
				"id": v.id,
			}

			c.JSON(200, res)

			}
		}

		if !flag {
			
			c.JSON(404, gin.H{
				"error": "City not found",
			})
		}
	})

	r.Run("0.0.0.0:9000")

}
