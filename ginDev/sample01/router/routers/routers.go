package routers

// go mod init learn.com/routers
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}
func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is the home page.")
}
func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcom to the index page.")
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	fmt.Println("routers::routers::SetupRouter() ...")
	r := gin.Default()
	r.GET("/", homeHandler)
	r.GET("/index", indexHandler)
	return r
}
