package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/ymzuiku/w/bom"
)

func main() {
	bom.Window.AddEventListener("load", func(_ *js.Object) {
		bom.Console.Log("hello", "world")
		ele := bom.Document.CreateElement("div")
		ele.TextContent("hello")
		bom.Console.Log(bom.Document.Body)
		bom.Document.Body.Append(ele)
	})
}
