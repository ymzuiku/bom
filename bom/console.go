package bom

import "github.com/gopherjs/gopherjs/js"

type ConsoleObject struct {
	*js.Object
}

var Console = ConsoleObject{Window.Get("console")}

func (c *ConsoleObject) Log(args ...any) {
	Console.Call("log", args...)
}

func (c *ConsoleObject) Info(args ...any) {
	Console.Call("info", args...)
}

func (c *ConsoleObject) Assert(args ...any) {
	Console.Call("assert", args...)
}

func (c *ConsoleObject) Clear(args ...any) {
	Console.Call("clear", args...)
}

func (c *ConsoleObject) Warn(args ...any) {
	Console.Call("warn", args...)
}

func (c *ConsoleObject) Time(args ...any) {
	Console.Call("time", args...)
}

func (c *ConsoleObject) TimeEnd(args ...any) {
	Console.Call("timeEnd", args...)
}
