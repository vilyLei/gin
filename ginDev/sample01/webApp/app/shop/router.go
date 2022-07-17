package shop

// go mod init webApp.com/shop
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// func shopHandler(c *gin.Context)
// func goodsHandler(c *gin.Context)
// func checkoutHandler(c *gin.Context)

func Routers(c *gin.Engine) {
	fmt.Println("app::router::Routers() ...")
	c.GET("/shop", shopHandler)
	c.GET("/goods", goodsHandler)
	c.GET("/checkout", checkoutHandler)
}
