package main

import (

	"fmt"

	"github.com/gin-gonic/gin"

	"strings"
	
	)

func main() {

	r := gin.Default()

	data := read()

	fmt.Println("Starting server")

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"ping" : "pong",
		})
	})

	r.GET("/:city", func(c *gin.Context) {

		city := c.Param("city")

		flag := false

		for _, v := range data {

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

	r.GET("/filter", func(c *gin.Context) {

		query := c.DefaultQuery("state", "xxx")

		var cities_data []map[string]string

		for _, v := range data {

			if strings.Contains(v.state, query) {

			current_city_data := map[string]string{
				"City": v.city,
				"State": v.state,
				"id": v.id,
			}

			cities_data = append(cities_data, current_city_data)
			}

		}

		if query == "xxx" {

			c.JSON(400, gin.H{
				"error": "Please Pass State Parameter",
			})
		} else {
		
		c.JSON(400, gin.H{
			"Results": cities_data,
		})
	}
	})

	r.GET("/", func(c *gin.Context) {

		var cities_data []map[string]string

		for _, v := range data {
			
			current_city_data := map[string]string{
				"City": v.city,
				"State": v.state,
				"id": v.id,
			}

			cities_data = append(cities_data, current_city_data)

			}
			c.JSON(400, gin.H{
				"Results": cities_data,
			})

		})

	r.Run("0.0.0.0:9000")

}
