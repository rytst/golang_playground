package main


import (
    "github.com/gin-gonic/gin"
)


type User struct {
    Id   string `json:"id" binding:"required"`
    Name string `json:"name"`
}


func main() {
    router := gin.Default()

    router.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }

        c.JSON(201, gin.H{
            "message": "created successfully",
            "data": user,
        })
    })

    router.Run(":8080")
}
