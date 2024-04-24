/*
Change the code from Exercise #3 and launch 50 goroutines
that calculate concurrently the square root of all the numbers between 100 and 149 (both included).



package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f is %.4f\n", f, x)
		wg.Done()
	}(16.1, &wg)

	wg.Wait()
}

*/

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(50)

	// IMPORTANT: i starts from 100. not 100
	// i is float64, not int. math.Sqrt() takes in a float64.
	for i := 100.; i < 150.; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			x := math.Sqrt(f)
			fmt.Printf("Square root of %.2f is %.6f\n", f, x)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
