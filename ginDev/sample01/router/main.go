package main

// go mod init learn.com/sample01/router/main
// thanks: https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/%E8%B7%AF%E7%94%B1%E6%8B%86%E5%88%86%E4%B8%8E%E6%B3%A8%E5%86%8C.html
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

var moduleMap map[string]string

type StrMap map[string]string

func initModuleConfig() {
	moduleMap = make(map[string]string)
	moduleMap["main"] = "./main.go"
	for moduleKey := range moduleMap {
		fmt.Println("module[", moduleKey, "] is ", moduleMap[moduleKey])
	}
}
func showStrMap(pmap StrMap) {

	fmt.Println("showStrMap(), init...")

	for key := range pmap {
		fmt.Println("map[", key, "] = ", pmap[key])
	}
}
func main() {
	fmt.Println("router main init ...")
	initModuleConfig()

	showStrMap(StrMap{"a": "sdfasfa", "vi": "hello, 01"})

	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	if err := r.Run(":9090"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
