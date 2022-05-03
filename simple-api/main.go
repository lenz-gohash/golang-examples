package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Creates a gin router
	r := gin.Default()

	// Authorized users
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"lenz":   "gohash",
		"matt":   "gohash",
		"ryo":    "gohash",
		"johann": "gohash",
		"derek":  "gohash",
	}))

	// Main route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Accessed unprotected endpoint",
		})
	})

	// Protected route. Only autorized users have access.
	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Accessed protected endpoint",
		})
	})

	// Post route
	// TODO: Print message to the console and echo it back to the client
	r.POST("/receiver", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Accessed POST endpoint",
		})
	})

	r.Run() // listen and serve on localhost:8080
}
