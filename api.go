package tree

import "syscall/js"

type APIFunction func (This, Args) interface {}
type This js.Value
type Args []js.Value

type API struct {
	name string
	entryPoint js.Func
    functions map[string]APIFunction
}

func wrapFunction (fn func (js.Value, []js.Value) interface{}) js.Func {
	return js.FuncOf(fn)
}

func (api *API) Add (fnName string, fn APIFunction) {
	api.functions[fnName] = fn
}

func (api *API) Publish () {
	api.entryPoint = wrapFunction(func (this js.Value, args []js.Value) interface {} {
		var publicFunctions = make(map[string]APIFunction)
		for jsName, fn := range(api.functions){
			publicFunctions[jsName] = fn
		}
		return publicFunctions
	})
	js.Global().Set(api.name, api.entryPoint)
}

func NewAPI (name string) *API {
	api := new(API)
	api.name = name
	return api
}