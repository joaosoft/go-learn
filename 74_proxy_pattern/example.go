package main

import "fmt"

// To use proxy and to object they must implement same methods
type IObject interface {
	ObjDo(action string)
}

// Object represents real objects which proxy will delegate data
type Object struct {
	action string
}

// ObjDo implements IObject interface and handel's all logic
func (obj *Object) ObjDo(action string) {
	// Action behavior
	fmt.Printf("I can, %s", action)
}

// ProxyObject represents proxy object with intercepts actions
type ProxyObject struct {
	object *Object
}

// ObjDo are implemented IObject and intercept action before send in real Object
func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "Run" {
		p.object.ObjDo(action) // Prints: I can, Run
	}
}

func main() {
	proxy := ProxyObject{}
	proxy.ObjDo("Run")
}
