package main

import (
	"log"
	"os"

	"github.com/ymzuiku/bom/helper"
)

func main() {
	log.Printf("__debug__%v", os.Args)
	if len(os.Args) < 3 {
		panic("args need like: bom xxx/main.go dist")
	}
	helper.Build(os.Args[1], os.Args[2])
}
