package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"flag"
	"fmt"
	"io"
)

func URLRoute(c *gin.Context) {
	res, err := http.Get(c.Request.URL.String())
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		// TODO: how do we handle this?
	}

	byteString, err := io.ReadAll(res.Body)
	resString := string(byteString)
	//fmt.Println("resString: %s", resString)

	//fmt.Println("status for client: %s", res.Status)
	c.String(res.StatusCode, resString)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false

	//r.GET("/ping", respondPing)

	// catchall route
	r.GET("*proxyPath", URLRoute)
	return r
}

func main() {
	wordPtr := flag.String("port", "2100", "port to bind proxy to")
	flag.Parse()

	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", *wordPtr))
}

