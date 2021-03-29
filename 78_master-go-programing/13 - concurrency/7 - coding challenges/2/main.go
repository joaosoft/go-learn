/*
1. Create a function called sum() that calculates and
then prints out the sum of 2 float numbers it receives as arguments.

Format the result with 2 decimal points.

2. From main launch 3 goroutines that execute the function you have just created (sum)

3. Synchronize the goroutines and the main function using WaitGroups
*/

package main

import (
	"fmt"
	"sync"
)

func sum(a, b float64, wg *sync.WaitGroup) {
	s := a + b
	fmt.Printf("Sum of %.2f and %.2f is %.2f\n", a, b, s)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go sum(10.3, 20.1, &wg)
	go sum(5, 7, &wg)
	go sum(33.5, 33.66, &wg)

	wg.Wait()
}
