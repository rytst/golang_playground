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

    router.GET("/products", func(c *gin.Context) {
        category := c.DefaultQuery("category", "all")

        c.JSON(200, gin.H{
            "category": category,
        })
    })

    router.Run(":8080")
}
