package tree

import (
	"reflect"
	"syscall/js"
	"testing"
)


func TestGetDocument(t *testing.T) {
	want := true
	o := GetDocument()
	if got := o.exist; got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestCreateTreeElement(t *testing.T) {
	want := true
	o := js.Global().Get("document")
	treeElement := CreateTreeElement(o)
	
	if got := treeElement.exist; got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}
	treeElementType := "*tree.TreeElement"
	if got := reflect.TypeOf(treeElement).String(); got != treeElementType {
		t.Errorf("got %#v, want %#v", got, want)
	}
	jsValueType := "js.Value"
	if got := reflect.TypeOf(treeElement.jsValue).String(); got != jsValueType {
		t.Errorf("got %#v, want %#v", got, want)
	}
	
	leafType := "tree.TreeLeaf"
	if got := leafType; got != reflect.TypeOf(treeElement.leaf).String() {
		t.Errorf("got %#v, want %#v", got, want)
	}
}