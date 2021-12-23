package tree
import "syscall/js"

func GetNodeType (jsValue js.Value) int {
	return jsValue.Get("nodeType")
}
