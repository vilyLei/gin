package main

// go mod init webApp.com/main
// go mod edit -replace webApp.com/blog=./app/blog
// go mod edit -replace webApp.com/shop=./app/shop
// go mod edit -replace webApp.com/routers=./routers

import (
	"fmt"

	"webApp.com/blog"
	"webApp.com/routers"
	"webApp.com/shop"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(":9090"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
