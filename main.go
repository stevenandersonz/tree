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


func (t TreeElement) GetElementById (id string) *TreeElement { 
    jsDoc := t.jsValue.Call("getElementById", id)
    return CreateTreeElement(jsDoc)
}
func (t TreeElement) AppendChild (child TreeElement) { 
    t.jsValue.Call("appendChild", child.jsValue)
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





