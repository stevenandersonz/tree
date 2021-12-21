package main

import (
	"syscall/js"

	"github.com/stevenandersonz/tree"
)


func printHello (this js.Value, args []js.Value) interface {} { return true }


func main() {
	api := tree.NewAPI("myAPI")
	api.Add("printHello",printHello)
	api.Publish()
	<-make(chan bool)
}