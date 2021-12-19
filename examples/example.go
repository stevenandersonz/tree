package main

import "github.com/stevenandersonz/tree"


func printHello (this tree.This, args tree.Args) interface {} {
	return "Hello"
}

func main() {
	api := tree.NewAPI("myAPI")
	api.Add("printHello", printHello)
	api.Publish()
}