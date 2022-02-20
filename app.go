package main

import (

	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"os"

	"encoding/csv"
)

func main() {

	godotenv.Load()

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"data": 1,
			"name": 2,
		})
	})

	app.Run()
}
