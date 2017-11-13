// Package main serves as an example application that makes use of the observer pattern.
// Playground: https://play.golang.org/p/cr8jEmDmw0
package main

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func main() {
	add := Operation{Addition{}}
	add.Operate(3, 5) // 8

	mult := Operation{Multiplication{}}
	mult.Operate(3, 5) // 15
}
