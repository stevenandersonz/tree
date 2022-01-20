package tree

import (
	"syscall/js"
)

type UIConsole struct {
	jsValue js.Value
	truthy bool 
}
func GetConsole () *UIConsole {
	console := js.Global().Get("console")
	uiConsole := new (UIConsole)
	uiConsole.truthy = console.Truthy()
	uiConsole.jsValue = console
	return uiConsole
}
func (console *UIConsole) Log (msg interface{}) {
	console.jsValue.Call("log", msg)
}
func (console *UIConsole) Error (msg string) {
	console.jsValue.Call("error", msg)
}
