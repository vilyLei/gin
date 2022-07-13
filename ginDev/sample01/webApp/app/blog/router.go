package blog

// go mod init webApp.com/blog
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routers(c *gin.Engine) {
	fmt.Println("app::router::BlogRouter() ...")
	c.GET("/blog", blogHandler)
	c.GET("/article", articleHandler)
}
