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


    router.PUT("/users/:userID", func(c *gin.Context) {
        userID := c.Param("userID")

        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }

        c.JSON(200, gin.H{
            "message": "User updated successfully",
            "id": userID,
            "data": user,
        })
    })

    router.Run(":8080")
}
