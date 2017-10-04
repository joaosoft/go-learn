package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_methods()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
// inicio -
// METODOS 1
type Vertex struct {
	X, Y float64
}

func (v Vertex) Soma() float64 {
	return v.X + v.Y
}

// METODOS 2
type MyFloat float64

func (f MyFloat) MyAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// METODO 3
func (v Vertex) Abs() float64 {
	if v.Soma() < 0 {
		return float64(-v.Soma())
	}
	return float64(v.Soma())
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

// METODO 4
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func _methods() {
	fmt.Println("_methods()")
	// METODO 1
	v := Vertex{3, 4}
	fmt.Println("Função Soma: ", v.Soma())

	// METODO 2
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.MyAbs())

	//METODO 3
	v3 := Vertex{1, 2}
	v3.Scale(10)
	fmt.Println(v3.Abs())

	// METODO 4
	v4 := Vertex{3, 4}
	v4.Scale(2)
	ScaleFunc(&v4, 10)

	p4 := &Vertex{4, 3}
	p4.Scale(3)
	ScaleFunc(p4, 8)

	fmt.Println(v4, p4)
	fmt.Println(v4, *p4)
}
