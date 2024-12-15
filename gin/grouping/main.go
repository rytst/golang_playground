package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("users", func(c *gin.Context) {
		users := []User{
			{Id: "1", Name: "user1"},
			{Id: "2", Name: "user2"},
		}

		c.JSON(200, gin.H{
			"data": users,
		})
	})

	router.Run(":8080")
}
