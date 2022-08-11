package bom

import "github.com/gopherjs/gopherjs/js"

type ArrayObject struct {
	*js.Object
}

var Array = ArrayObject{Window.Get("Array")}

func (s *ArrayObject) IsArray(v *js.Object) bool {
	return s.Call("isArray", v).Bool()
}
