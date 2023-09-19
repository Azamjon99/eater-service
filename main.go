package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a route that responds with "pong" when accessed
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Run the Gin web server on port 8080
	r.Run(":8080")
}