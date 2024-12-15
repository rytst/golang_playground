package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sample API",
			"int":     1,
			"map": map[string]int{
				"key1": 100,
				"key2": 200,
			},
		})
	})

    router.Run(":8080")
}
