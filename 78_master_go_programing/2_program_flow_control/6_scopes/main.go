/////////////////////////////////
// Scopes in Go
// Go Playground: https://play.golang.org/p/jyKwZ_glrrY
/////////////////////////////////

// There are 3 Scopes:
//  - File Scope
//  - Package Scope
//  - Block (local)  Scope

package main

// import statements are file scoped
import (
	"fmt"

	// import "fmt" -> error, within the same scope, unique names

	// importing as another name (alias) is permitted
	f "fmt"
)

// variables or constant declared outside any function are package scoped
const done = false

func main() { // package scoped

	// block scoped: visible until the end of the block "}"
	var task = "Running:"
	fmt.Println(task, done) // => Running: false (this is done from package scope)
	f.Println("Bye bye!")

	// names must be unique only within the same scope
	const done = true                        // local scoped
	fmt.Printf("done in main(): %v\n", done) // => done in main(): true
	f1()
}

func f1() {
	fmt.Printf("done in f(): %v\n", done) //this is done from package scope
}