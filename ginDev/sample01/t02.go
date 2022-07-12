package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func helloFunc(c *gin.Context) {
	c.String(http.StatusOK, "T02, Hello World!")
}
func main() {
	port := os.Getenv("GO_LOCAL_PORT")

	fmt.Println("port: ", port)
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define handlers
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello World!")
	// })
	r.GET("/", helloFunc)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
