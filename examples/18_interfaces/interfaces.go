package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_interfaces()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
// inicio -
// INTERFACES 1
type Abser interface {
	Abs() float64
}
type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex1 struct {
	X, Y float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// INTERFACES 2

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// INTERFACES 4
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func _interfaces() {
	fmt.Println("_interfaces()")
	// Um tipo de interface é definida por um conjunto de métodos
	// Um valor de tipo de interface pode conter qualquer valor que implementa esses métodos.

	// INTERFACES 1
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := Vertex1{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v // ERRO!
	fmt.Println(a.Abs())

	// INTERFACES 2
	var i I = T{"hello"}
	i.M()

	// INTERFACES 3
	var ii interface{} = "hello"

	ss := ii.(string)
	fmt.Println(ss)

	ss, okk := ii.(string)
	fmt.Println(ss, okk)

	ff, okk := ii.(float64)
	fmt.Println(ff, okk)

	//ff = ii.(float64) // panic
	//fmt.Println(ff)

	// INTERFACES 4
	do(21)
	do("hello")
	do(true)
}
