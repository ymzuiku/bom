package bom

import "github.com/gopherjs/gopherjs/js"

func init() {
	Window.AddEventListener("load", func(_ *js.Object) {
		Document.Body = &Element{Document.Get("body")}
	})
}
