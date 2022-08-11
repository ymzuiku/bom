package bom

import "github.com/gopherjs/gopherjs/js"

func (w *WindowObject) Alert(s string) {
	w.Call("alert", s)
}

func IsNaN(v *js.Object) bool {
	return Window.Call("isNaN", v).Bool()
}
