package main

import (
	"fmt"
	"strconv"
)

const (
	name   = "TESTE 0"
	number = 0
)

// Controller ...
type Controller struct {
	name   string
	number int
}

// ControllerOption ...
type ControllerOption func(controller *Controller)

// WithControllerNumber ...
func WithControllerNumber(number int) ControllerOption {
	return func(controller *Controller) {
		if number >= 1 {
			controller.number = number
		}
	}
}

// WithControllerName ...
func WithControllerName(name string) ControllerOption {
	return func(controller *Controller) {
		controller.name = name
	}
}

// NewController ...
func NewController(opts ...ControllerOption) *Controller {
	controller := &Controller{
		name:   name,
		number: number,
	}

	controller.Reconfigure(opts...)

	return controller
}

// Reconfigure ...
func (controller *Controller) Reconfigure(opts ...ControllerOption) {
	for _, opt := range opts {
		opt(controller)
	}
}

// HandleRequest ...
func (controller *Controller) HandleRequest(msg string, opts ...ControllerOption) string {
	controller.Reconfigure(opts...)

	return msg
}

func main() {
	controller := NewController(WithControllerName("TESTE 1"), WithControllerNumber(1))
	fmt.Println("NAME:" + controller.name + "RETRIES:" + strconv.Itoa(controller.number))

	controller.HandleRequest("TESTE 2", WithControllerName("TESTE 2"), WithControllerNumber(2))
	fmt.Println("NAME:" + controller.name + "RETRIES:" + strconv.Itoa(controller.number))

	controller.HandleRequest("TESTE 3")
	fmt.Println("NAME:" + controller.name + "RETRIES:" + strconv.Itoa(controller.number))

}
