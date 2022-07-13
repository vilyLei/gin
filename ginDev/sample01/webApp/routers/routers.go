package routers

// go mod init webApp.com/routers
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.Default()
	baseRouter(r)
	for _, opt := range options {
		opt(r)
	}
	return r
}

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.webApp.com",
	})
}
func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is the home page.")
}
func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcom to the index page.")
}

// SetupRouter 配置路由信息
func baseRouter(e *gin.Engine) {
	fmt.Println("routers::routers::SetupRouter() ...")
	e.GET("/", homeHandler)
	e.GET("/index", indexHandler)
}
