package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ymzuiku/w/pkg/execx"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	app := gin.New()

	go func() {
		execx.Run(context.Background(), nil, "gopherjs", "serve", "example/client/main.go")
	}()

	app.Use(func(ctx *gin.Context) {
		remotePath := ctx.Request.URL.String()
		if !strings.Contains(remotePath, "/v1") {
			remote, err := url.Parse("http://127.0.0.1:9100" + remotePath)
			if err != nil {
				fmt.Println("proxy:err", err)
				ctx.Abort()
				return
			}
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.Director = func(req *http.Request) {
				req.Header = ctx.Request.Header
				req.Host = remote.Host
				req.URL.Scheme = remote.Scheme
				req.URL.Host = remote.Host
				req.URL.Path = remote.Path
			}
			proxy.ServeHTTP(ctx.Writer, ctx.Request)
		}
	})

	log.Printf("listen: http://127.0.0.1:8300")
	if err := app.Run(":8300"); err != nil {
		fmt.Println("rightos app run err: ", err)
	}
}
