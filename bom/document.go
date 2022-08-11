package bom

import "github.com/gopherjs/gopherjs/js"

type DocumentObject struct {
	*Element
	Body *Element
}

var Document = DocumentObject{
	Element: &Element{
		Object: js.Global.Get("document"),
	},
	Body: nil,
}

func (s *DocumentObject) CreateElement(tag string) *Element {
	return &Element{Document.Call("createElement", tag)}
}
