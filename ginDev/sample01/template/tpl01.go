package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templ/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "templ01.html", gin.H{"title": "templ01 Page", "ce": "[content info]"})
	})
	r.Run(":9090")
}
