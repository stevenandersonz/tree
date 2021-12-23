package tree

import (
	"strings"
	"syscall/js"

)
const DOCUMENT_NODE = 9
//UI Component - Tries to imitate HTML Element API
//jsValue -> js.Value 
//elem -> a wrapper of js.Value targeted to HTML Elements
//truthy -> determines wheter the element exist or not
type UIDocument struct {
	jsValue js.Value
	truthy bool 
}
func checkDocumentNode (el js.Value) {
	if GetNodeType(el) == DOCUMENT_NODE {
		panic("Not a DOCUMENT_NODE")  
	}
}
func GetDocument () *UIDocument {
	doc := js.Global().Get("document")
	checkDocumentNode(doc)
	uiDoc := new (UIDocument)
	uiDoc.truthy = doc.Truthy()
	uiDoc.jsValue = doc
	return uiDoc
}

// gets a element by its ID
func (doc *UIDocument) GetElementById (id string) *UIElement {
    checkDocumentNode(doc.jsValue)
	htmlElement := doc.jsValue.Call("getElementById", id)
	return CreateUIElement(htmlElement)
}
