package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from Vercel!"})
}

func SetupRouter(r *gin.Engine) {
	r.GET("/hello", HelloHandler)
}