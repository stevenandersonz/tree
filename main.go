package tree

import (
	"strings"
	"syscall/js"
)

type TreeElement js.Value

func checkIfElementExist(node js.Value)bool {
    if node.Truthy() {
        return true
    }
    return false
}

func convertToTreeEl(node js.Value) TreeElement {
	return TreeElement(node)
}

func (node TreeElement) GetElementById (id string) TreeElement { 
    el := js.Value(node).Call("getElementById", id)
    return TreeElement(el)
}
func (node TreeElement) AppendChild (child TreeElement) { 
    js.Value(node).Call("appendChild", js.Value(child))
}

func (node TreeElement) CreateElement (tag string, props []string) TreeElement { 
    el := js.Value(node).Call("createElement", tag)
    for _, prop := range props {
        propArr := strings.Split(prop, "=")
        name := propArr[0]
        value := propArr[1]
        el.Call("setAttribute", name, value)
    }
    return TreeElement(el)
}





