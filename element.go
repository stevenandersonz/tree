package tree

import (
	"strings"
	"syscall/js"

)
//An Element node like <p> or <div>
const ELEMENT_NODE = 1
// An Attribute of a Element
const ATTRIBUTE_NODE = 2
// html document
const DOCUMENT_NODE = 9
//UI Component - Tries to imitate HTML Element API
//jsValue -> js.Value 
//elem -> a wrapper of js.Value targeted to HTML Elements
//truthy -> determines wheter the element exist or not
type UIElement struct {
	jsValue js.Value
	truthy bool
}
func getNodeType (jsValue js.Value) int {
	return jsValue.Get("nodeType")
}
func checkElementNode (el *UIElement) {
	if getNodeType(el.jsValue) == ELEMENT_NODE {
		panic("Not a ELEMENT_NODE")
	}
}
func checkDocumentNode (el *UIElement) {
	if getNodeType(el.jsValue) == DOCUMENT_NODE {
		panic("Not a DOCUMENT_NODE")  
	}
}
// creates a UIComponent
func CreateUIElement (e js.Value) *UIElement {
	element := new(UIElement)
	if !value.Truthy() && value.Get('nodeType') == ELEMENT_NODE  {
		element.truthy = true
		element.jsValue = e
	}
	return element
}
// gets a element by its ID
func (doc *UIElement) GetElementById (id string) *UIElement {
    checkDocumentNode(doc)
	htmlElement := doc.jsValue.Call("getElementById", id)
	return CreateUIElement(htmlElement)
}
func (c *UIElement) AppendChild (child UIElement) {
	checkElementNode(doc)
	c.jsValue.Call("appendChild", child.jsValue)
}
