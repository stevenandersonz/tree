package main

import (
	"syscall/js"
	"github.com/stevenandersonz/tree"
	"time"
	"fmt"
)


func printHello (this js.Value, args []js.Value) interface {} { return true }

func fib(n int) int {
	if n<1 {
		return 0
	}
	if n<=2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	doc := tree.GetDocument()
	console := tree.GetConsole()
	console.Log(tree.GetNodeType(doc.JSValue))
//	mainSection := doc.GetElementById("tree-main")
	start := time.Now()
	f := fib(30)
//	for i:=0; i<1000; i++ {
//		p := doc.CreateHTMLElement("p")
//		text := doc.CreateHTMLTextNode("Hello from tree")
//		p.AppendChild(text)
//		mainSection.AppendChild(p)
//	}
	elapsed:=time.Since(start)
	console.Log(fmt.Sprintf("took %v to compute fib(%v)", elapsed, f))
	api := tree.NewAPI("myAPI")
	api.Add("printHello",printHello)
	api.Publish()
	<-make(chan bool)
}
