package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name`
}

func main() {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		var user User

        // Binding to a struct with json
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{
			"message": "User created successfully",
			"data":    user,
		})
	})

	router.Run(":8080")
}
