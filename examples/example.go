package main

import (
	"syscall/js"
	"github.com/stevenandersonz/tree"
)


func printHello (this js.Value, args []js.Value) interface {} { return true }


func main() {
	doc := tree.GetDocument()
	console := tree.GetConsole()
	console.Log(tree.GetNodeType(doc.JSValue))
	mainSection := doc.GetElementById("tree-main")
	p := doc.CreateHTMLElement("p")
	text := doc.CreateHTMLTextNode("Hello from tree")
	p.AppendChild(text)
	mainSection.AppendChild(p)
	api := tree.NewAPI("myAPI")
	api.Add("printHello",printHello)
	api.Publish()
	<-make(chan bool)
}
