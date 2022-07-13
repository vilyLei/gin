package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func blogHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcom to my blog.")
}
func articleHandler(c *gin.Context) {
	c.String(http.StatusOK, "There are some articles here.")
}
