package shop

// go mod init webApp.com/shop
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routers(c *gin.Engine) {
	fmt.Println("app::router::Routers() ...")
	c.GET("/shop", shopHandler)
	c.GET("/goods", goodsHandler)
	c.GET("/checkout", checkoutHandler)
}
