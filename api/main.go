package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	// read .env
	fmt.Println("value:", os.Getenv("key"))

	api := gin.Default()
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongx",
		})
	})
	api.Run()
}
