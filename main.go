package main


import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func getHelloWorld(c *gin.Context) {
    c.String(http.StatusOK, "Hello, World!")
}

func main() {
    router := gin.Default()
    router.GET("/", getHelloWorld)

    router.Run("localhost:8080")
}