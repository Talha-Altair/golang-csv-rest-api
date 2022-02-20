package main

import (

	"fmt"

	"github.com/gin-gonic/gin"

	// "csv_data"
	
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

	r.Run("0.0.0.0:9000")

}
