package main

import "fmt"
import (
	"github.com/keltia/leftpad"
	"strings"
)

func main() {
	// WITH LEFTPAD

	str1, _ := leftpad.Pad("foo", 5) // "  foo"
	fmt.Println(str1)

	str2, _ := leftpad.Pad("foobar", 8) // "  foobar"
	fmt.Println(str2)

	str3, _ := leftpad.Pad("foobar", 6) // "foobar"
	fmt.Println(str3)

	str4, _ := leftpad.PadChar("foo", 5, 'X') // "XXfoo"
	fmt.Println(str4)

	teste := strings.Join([]string{str1, str2, str3, str4}, ".")
	fmt.Println(teste)

	// WITH SPRINTF
	a := 1
	b := 2
	c := 3
	fmt.Println(fmt.Sprintf("%.4d.%.2d.%.2d", a, b, c))
}
