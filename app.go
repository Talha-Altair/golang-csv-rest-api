package main

import (

	"fmt"

	"github.com/gin-gonic/gin"

	// "os"
	
	)

func main() {

	r := gin.Default()

	fmt.Println("Starting server")

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"ping" : "pong",
		})
	})

	r.Run("0.0.0.0:9000")

}
