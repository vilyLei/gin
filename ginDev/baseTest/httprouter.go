package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VpostGetting(c *gin.Context) {
	c.String(http.StatusOK, "hello VpostGetting")
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word httprouter")
	})
	r.POST("/vpost", VpostGetting)
	r.PUT("/vput")
	//监听端口默认为8080
	r.Run(":9090")
}
