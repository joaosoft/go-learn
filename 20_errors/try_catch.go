package main

import (
	"fmt"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (block Block) Do() {
	if block.Finally != nil {
		defer block.Finally()
	}
	if block.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				block.Catch(r)
			}
		}()
	}
	block.Try()
}

func main() {
	fmt.Println("We started")
	Block{
		Try: func() {
			fmt.Println("I tried")
			Throw("Oh,...sh...")
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()
	fmt.Println("We went on")
}
