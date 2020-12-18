package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	myInt   bool    // 1 byte
	myFloat float64 // 8 bytes
	myBool  int32   // 4 bytes
}

type myStructOptimized struct {
	myFloat float64 // 8 bytes
	myInt   int32   // 4 bytes
	myBool  bool    // 1 byte
}

func main() {
	a := myStruct{}
	b := myStructOptimized{}

	fmt.Println("Sizes")
	fmt.Printf("a.myInt: %d\n", unsafe.Sizeof(a.myInt))     // unordered 1 bytes
	fmt.Printf("a.myFloat: %d\n", unsafe.Sizeof(a.myFloat)) // unordered 8 bytes
	fmt.Printf("a.myBool: %d\n", unsafe.Sizeof(a.myBool))   // unordered 4 bytes

	fmt.Println("\nUnoptimized struct")
	fmt.Printf("Size: %d\n", unsafe.Sizeof(a)) // unordered 24 bytes

	fmt.Println("\nOptimized struct")
	fmt.Printf("Size: %d\n", unsafe.Sizeof(b)) // ordered 16 bytes
}
