/////////////////////////////////////////
// Comments and Naming Conventions in Go
// Go Playground: https://play.golang.org/p/pprI80SPMkS
/////////////////////////////////////////

package main

//** COMMENTS **//

// this is a single line comment

/*
 This is a block comment.
 a := 10
 fmt.Println(a)
*/

var name = "John Wick" // inline comment

//** NAMING CONVENTIONS IN GO **//

// Naming Conventions are important for code readability and maintainability.

// use short, concise names especially in shorter scopes
// common names for common types:
var s string   //string
var i int      //index
var num int    //number
var msg string //message
var v string   //value
var err error  //error value
var done bool  //bool, has been done?

// use mixedCase a.k.a camelCase instead of snake_case (variables and  functions)
var maxValue = 100  // recommended (camelCase)
var max_value = 100 // not recommended (snake_case)

// recommended
func writeToFile() {
}

// not recommended
func write_to_file() {
}

// write acronyms in all caps
var writeToDB = true // recommended
var writeToDb = true // not recommended

func main() {

	// use fewer letters, don’t be too verbose especially in smaller scopes
	var packetsReceived int // NOT OK, too verbose
	var n int               // OK
	_, _ = packetsReceived, n

	// an uppercase first letter has special significance to go (it will be exported in other packages)
}
