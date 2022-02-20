package main

import (

	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/bxcodec/faker/v3"

	"github.com/joho/godotenv"

	"os"

	"encoding/csv"
)

type faker_struct struct {
	Name     string `faker:"name"`
	Currency string `faker:"currency"`
}

func main() {

	// gin.SetMode(gin.ReleaseMode)

	godotenv.Load()

	altair := os.Getenv("name")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		a := faker_struct{}

		err := faker.FakeData(&a)

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"data": a,
			"name": altair,
		})
	})

	main2()

	r.Run()
}
