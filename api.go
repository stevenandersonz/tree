package tree

import "syscall/js"

type APIFunction func (This, Args) interface {}
type This js.Value
type Args []js.Value

type API struct {
	name string
	entryPoint js.Func
    functions map[string] interface {}
}

func wrapFunction (fn func (js.Value, []js.Value) interface{}) js.Func {
	return js.FuncOf(fn)
}

func (api *API) Add (fnName string, fn func (js.Value, []js.Value)interface{}) {
	api.functions[fnName] = js.FuncOf(fn)
}

func (api *API) Publish () {
	var publicFunctions = make(map[string]interface{})
	for jsName, fn:=range(api.functions){
		publicFunctions[jsName] = fn
	}
	api.entryPoint = js.FuncOf(func (this js.Value, args []js.Value) interface {} {	
		return publicFunctions
	})
	js.Global().Set(api.name, api.entryPoint)
}

func NewAPI (name string) *API {
	api := new(API)
	api.name = name
	api.functions = make(map[string]interface{})
	return api
}