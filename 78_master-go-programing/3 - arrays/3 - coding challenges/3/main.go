/*
There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = a

	myArray[3] = 10.10

	fmt.Println(myArray)

}
*/
package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	// ERROR -> invalid type int a set to a float64 type
	// myArray[0] = a
	myArray[0] = float64(a)

	// ERROR -> invalid array index 3 (out of bounds for 3-element array)
	// myArray[3] = 10.10

	myArray[2] = 10.10

	fmt.Println(myArray)

}
