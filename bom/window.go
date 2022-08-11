package bom

import "github.com/gopherjs/gopherjs/js"

type WindowObject struct {
	*Element
}

var Window = WindowObject{
	&Element{
		Object: js.Global,
	},
}

func (w *WindowObject) Confirm(s string) bool {
	return w.Call("confirm", s).Bool()
}
