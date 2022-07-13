package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func shopHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcom to my shop.")
}
func goodsHandler(c *gin.Context) {
	c.String(http.StatusOK, "There are many goods here.")
}
func checkoutHandler(c *gin.Context) {
	c.String(http.StatusOK, "Chechout your receipts.")
}
