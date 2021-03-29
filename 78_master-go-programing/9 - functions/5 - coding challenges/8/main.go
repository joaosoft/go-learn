/*
Consider the following Go program that prints out:

----
The Go gopher is the iconic mascot of the Go project.

Hello, Go playground!
----

And Modify only the line in the main() body function where the print() function
is invoked so that the program will print out "Hello, Go playground!" and then
"The Go gopher is the iconic mascot of the Go project."



package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}

*/

package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}
