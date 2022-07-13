package routers

import (
	"fmt"
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
func ShopRouter(c *gin.Engine) {
	fmt.Println("routers::shop::ShopRouter() ...")
	c.GET("/shop", shopHandler)
	c.GET("/goods", goodsHandler)
	c.GET("/checkout", checkoutHandler)
}
