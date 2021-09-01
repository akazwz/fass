package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	err := server.Run("0.0.0.0:9000")
	if err != nil {
		log.Fatalln("server run error")
	}
}
