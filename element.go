package tree

import (
	"strings"
	"syscall/js"

)
//An Element node like <p> or <div>
const ELEMENT_NODE = 1
// An Attribute of a Element
const ATTRIBUTE_NODE = 2
//UI Component - Tries to imitate HTML Element API
//jsValue -> js.Value 
//elem -> a wrapper of js.Value targeted to HTML Elements
//truthy -> determines wheter the element exist or not
type UIElement struct {
	jsValue js.Value
	truthy bool
}



func checkElementNode (el js.Value) {
	if GetNodeType(el) == ELEMENT_NODE {
		panic("Not a ELEMENT_NODE")
	}
}


// creates a UIComponent
func CreateUIElement (e js.Value) *UIElement {
	checkElementNode(e)
	element := new(UIElement)
	element.truthy = true
	element.jsValue = e
	return element
}

func (c *UIElement) AppendChild (child UIElement) {
	checkElementNode(c.jsValue)
	c.jsValue.Call("appendChild", child.jsValue)
}
