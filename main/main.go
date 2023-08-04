package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
