package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	fmt.Println("Rota ping")
	c.JSON(200, gin.H{"message": "pong"})
}
