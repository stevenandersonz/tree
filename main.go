package tree

import (
	"strings"
	"syscall/js"
)

type TreeLeaf js.Value
type TreeElement struct {
    jsValue js.Value
    leaf TreeLeaf
    exist bool
}
func CreateTreeElement (value js.Value) *TreeElement {
    t := new(TreeElement)
    if !value.Truthy() {
        t.exist = false
        return t
    }
    t.exist = true
    t.jsValue = t.jsValue
    t.leaf = TreeLeaf(value)
    return t
}

func GetDocument () *TreeElement {
    t := CreateTreeElement(js.Global().Get("document"))
    return t
}
func ConsoleLog(args ...interface{})  {
    console := js.Global().Get("console")
    if !console.Truthy() {
        panic("window.console couldn't be fetch, make sure make sure there is a running instance of window object.")
    }
    if 1 > len(args){
       console.Call("error", "ConsoleLog was expecting at least 1 argument instead got 0")
    }
    for _, arg :=range(args){
        val := js.ValueOf(arg) 
        console.Call("log",val) 
    }
}

func (t TreeElement) GetElementById (id string) *TreeElement { 
    jsDoc := t.jsValue.Call("getElementById", id)
    return CreateTreeElement(jsDoc)
}
func (t TreeElement) AppendChild (child TreeElement) { 
    t.jsValue.Call("appendChild", child.jsValue)
}
func (t TreeElement) SetClassName (class string) {
    t.jsValue.Set("className", class)
}
func (t TreeElement) CreateElement (tag string, props []string) *TreeElement { 
    jsDoc := t.jsValue.Call("createElement", tag)
    for _, prop := range props {
        propArr := strings.Split(prop, "=")
        name := propArr[0]
        value := propArr[1]
        jsDoc.Call("setAttribute", name, value)
    }
    newT := CreateTreeElement(jsDoc)
    return newT
}






