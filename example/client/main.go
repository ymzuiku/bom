package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/ymzuiku/bom/bom"
)

func main() {
	bom.Window.AddEventListener("load", func(_ *js.Object) {
		bom.Console.Log("hello", "world")
		ele := bom.Document.CreateElement("div")
		ele.TextContent("hello")
		ele.ClassName("bg-red-200")
		bom.Console.Log(bom.Document.Body)
		bom.Document.Body.Append(ele)
	})
}
