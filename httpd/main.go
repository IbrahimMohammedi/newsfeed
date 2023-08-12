package main

import (
	"newsfeeder/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet)

	r.Run()

}
