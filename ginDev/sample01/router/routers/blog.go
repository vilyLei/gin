package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func blogHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcom to my blog.")
}
func articleHandler(c *gin.Context) {
	c.String(http.StatusOK, "There are some articles here.")
}
func BlogRouter(c *gin.Engine) {
	fmt.Println("routers::blog::BlogRouter() ...")
	c.GET("/blog", blogHandler)
	c.GET("/article", articleHandler)
}
