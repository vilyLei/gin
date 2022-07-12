package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// go get -u github.com/gin-gonic/gin
/*
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/sqs/goreturns
go get -v github.com/rogpeppe/godef
*/
func main() {
	fmt.Println("t01::main main init...")
	gopath := os.Getenv("GOPARH")
	fmt.Println("gopath: ", gopath)
	adrHome := os.Getenv("ANDROID_HOME")
	fmt.Println("adrHome: ", adrHome)
	//

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
