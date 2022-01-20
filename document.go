package tree

import (
	"syscall/js"
)
const DOCUMENT_NODE = 9
//UI Component - Tries to imitate HTML Element API
//jsValue -> js.Value 
//elem -> a wrapper of js.Value targeted to HTML Elements
//truthy -> determines wheter the element exist or not
type UIDocument struct {
	JSValue js.Value
	Truthy bool 
}
func checkDocumentNode (el js.Value) {
	if GetNodeType(el) != DOCUMENT_NODE {
		panic("Not a DOCUMENT_NODE")  
	}
}
func GetDocument () *UIDocument {
	doc := js.Global().Get("document")
	checkDocumentNode(doc)
	uiDoc := new (UIDocument)
	uiDoc.Truthy = doc.Truthy()
	uiDoc.JSValue = doc
	return uiDoc
}

// gets a element by its ID
func (doc *UIDocument) GetElementById (id string) *UIElement {
    checkDocumentNode(doc.JSValue)
	htmlElement := doc.JSValue.Call("getElementById", id)
	return CreateUIElement(htmlElement)
}
func (doc *UIDocument) GetElementBySelector  (selector string) *UIElement {
	checkDocumentNode(doc.JSValue)
	htmlElement := doc.JSValue.Call("querySelector", selector)
	return CreateUIElement(htmlElement)
}
func (doc *UIDocument) CreateHTMLElement (tag string) *UIElement {
	htmlElement := doc.JSValue.Call("createElement", tag)
	return CreateUIElement(htmlElement)
}
func (doc *UIDocument) CreateHTMLTextNode (text string) *UIElement {
	htmlElement := doc.JSValue.Call("createTextNode", text)
	return CreateUIElement(htmlElement)
}

