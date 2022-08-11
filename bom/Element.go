package bom

import (
	"github.com/gopherjs/gopherjs/js"
)

type Element struct {
	*js.Object
}

func (ev *Element) AddEventListener(key string, fn func(e *js.Object)) {
	ev.Call("addEventListener", key, fn)
}

func (ev *Element) Append(args ...any) {
	ev.Call("append", args...)
}

func (ev *Element) TextContent(text string) {
	ev.Set("textContent", text)
}

func (ev *Element) ClassName(text string) {
	ev.Set("className", text)
}

func (ev *Element) GetClassName() string {
	return ev.Get("className").String()
}

func (ev *Element) GetTextContent(text string) string {
	return ev.Get("textContent").String()
}
