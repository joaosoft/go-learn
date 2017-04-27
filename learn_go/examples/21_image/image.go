package main

import (
	"fmt"
	"image"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_image()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _image() {
	fmt.Println("_image()")
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
