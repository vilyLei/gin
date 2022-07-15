package main

// thanks: https://andblog.cn/2941
// http://localhost:9090/HelloWordGet
// http://localhost:9090/HelloWordGet?isAbort=tru
// http://localhost:9090/HelloWordGet?isAbort=true
// http://localhost:9090/HelloWordGet?isAbort=false
import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HelloWordGet(c *gin.Context) {
	fmt.Println("exec hello word get func")
	c.String(http.StatusOK, "hello world get")
}

func main() {
	router := gin.Default()
	router.Use(FirstMiddleware(), SecondMiddleware())

	router.GET("/HelloWordGet", Logger(), HelloWordGet)

	router.Run(":9090")

}

func FirstMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first middleware before next()")

		isAbort := c.Query("isAbort")
		bAbort, err := strconv.ParseBool(isAbort)
		if err != nil {
			fmt.Printf("is abort value err, value %s\n", isAbort)
			c.Next() // (2)
		}
		if bAbort {
			fmt.Println("first middleware abort") //(3)
			c.Abort()
			//c.AbortWithStatusJSON(http.StatusOK, "abort is true")
			return
		} else {
			fmt.Println("first middleware doesnot abort") //(4)
			return
		}

		fmt.Println("first middleware after next()")
	}
}

func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("current inside of second middleware")
	}
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		//context.Next()
		fmt.Println(context.Writer.Status())
	}

}
