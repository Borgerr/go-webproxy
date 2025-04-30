package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"flag"
	"fmt"
)

func respond(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", respond)
	return r
}

func main() {
	wordPtr := flag.String("port", "2100", "port to bind proxy to")
	flag.Parse()

	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", *wordPtr))
}

