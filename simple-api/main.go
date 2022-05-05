package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	// Body is the message body
	Body string `json:"body"` // The json tag is the key name in the json object
}

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

	// Post route: Prints POST message body to the console and echoes it back to
	// the client
	r.POST("/receiver", func(c *gin.Context) {
		var msg Message

		// Bind the message body to the `Message`` struct
		if err := c.BindJSON(&msg); err != nil {
			return
		}

		// Print message to the console
		fmt.Printf("POST Request \"message\": \"%v\" \n", msg.Body)

		// Serialize the message body and send it back to the client
		c.IndentedJSON(http.StatusOK, msg)
	})

	r.Run() // listen and serve on localhost:8080
}
