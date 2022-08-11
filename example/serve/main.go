package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ymzuiku/bom/helper"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	app := gin.New()

	helper.Port = "9100"
	helper.Proxy(app, "example/client/main.go")

	log.Printf("listen: http://127.0.0.1:8300")
	if err := app.Run(":8300"); err != nil {
		fmt.Println("rightos app run err: ", err)
	}
}
