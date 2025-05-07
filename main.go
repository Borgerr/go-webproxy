package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"flag"
	"fmt"
	"io"
)

func guy(c *gin.Context) {
	fmt.Println("c.FullPath(): ", c.FullPath())
	fmt.Println("c.request.URL: ", c.Request.URL)
	fmt.Println("c.request.URL.path: ", c.Request.URL.Path)

	res, err := http.Get(c.Request.URL.String())
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		// TODO: how do we handle this?
	}

	byteString, err := io.ReadAll(res.Body)
	resString := string(byteString)
	fmt.Println("resString: %s", resString)

	fmt.Println("status for client: %s", res.Status)
	//c.String(res.String())
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	//r.GET("/ping", respondPing)

	// catchall route
	r.Any("/*proxyPath", guy)
	return r
}

func main() {
	wordPtr := flag.String("port", "2100", "port to bind proxy to")
	flag.Parse()

	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", *wordPtr))
}

