package main

import (
	"net/http"
	//"net/url"
	"github.com/gin-gonic/gin"
	"flag"
	"fmt"
)

func respondPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func catchallFunc(c *gin.Context) {
	fmt.Println("c.FullPath(): ", c.FullPath())
	fmt.Println("c.request.URL: ", c.Request.URL)
	fmt.Println("c.request.URL.path: ", c.Request.URL.Path)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	//r.GET("/ping", respondPing)

	// catchall route
	r.Any("/*proxyPath", catchallFunc)
	return r
}

func main() {
	wordPtr := flag.String("port", "2100", "port to bind proxy to")
	flag.Parse()

	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", *wordPtr))
}

