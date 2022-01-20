package tree

import (
	"syscall/js"
)
//An Element node like <p> or <div>
const ELEMENT_NODE = 1
// An Attribute of a Element
const ATTRIBUTE_NODE = 2
const TEXT_NODE = 3
//UI Component - Tries to imitate HTML Element API
//jsValue -> js.Value 
//elem -> a wrapper of js.Value targeted to HTML Element
//truthy -> determines wheter the element exist or not
type UIElement struct {
	JSValue js.Value
	Truthy bool
}



func checkElementNode (el js.Value) {
	if GetNodeType(el) != ELEMENT_NODE && GetNodeType(el) != TEXT_NODE {
		panic("Not a ELEMENT_NODE")
	}
}


// creates a UIComponent
func CreateUIElement (e js.Value) *UIElement {
	checkElementNode(e)
	element := new(UIElement)
	element.Truthy = true
	element.JSValue = e
	return element
}

func (c *UIElement) AppendChild (child *UIElement) {
	checkElementNode(c.JSValue)
	c.JSValue.Call("appendChild", child.JSValue)
}
func (c *UIElement) InsertBefore (newEl UIElement, e UIElement) {
	checkElementNode(c.JSValue)
	c.JSValue.Call("InsertBefore", newEl.JSValue, e.JSValue)
}

