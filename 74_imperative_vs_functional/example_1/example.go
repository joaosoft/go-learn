package main

import (
	"fmt"
	"log"
	"time"
)

func imperative() {
	n, numElements, s := 1, 0, 0
	for numElements < 10 {
		if n*n%5 == 0 {
			s += n
			numElements++
		}
		n++
	}
	fmt.Println(s) //275
}

func Filter(i interface{}) bool {
	return i.(uint64) > 5
}

func Reduce(memo interface{}, el interface{}) interface{} {
	return len(memo.(string)) + len(el.(string))
}

//func functional() {
//	sum := func(memo interface{}, el interface{}) interface{} {
//		return memo.(float64) + el.(float64)
//	}
//	pred := func(i interface{}) bool {
//		return (i.(uint64)*i.(uint64))%5 == 0
//	}
//	createValues := func() []int {
//		values := make([]int, 100)
//		for num := 1; num <= 100; num++ {
//			values = append(values, num)
//		}
//		return values
//	}
//	Reduce(Filter(pred, createValues), sum, uint64).(uint64)
//}

func main() {

	var start time.Time
	var elapsed time.Duration

	start = time.Now()
	imperative()
	elapsed = time.Since(start)
	log.Printf("Imperative took %s", elapsed)

	start = time.Now()
	imperative()
	elapsed = time.Since(start)
	log.Printf("Functional took %s", elapsed)

	start = time.Now()
	imperative()
	elapsed = time.Since(start)
	log.Printf("Functional took %s", elapsed)

	start = time.Now()
	imperative()
	elapsed = time.Since(start)
	log.Printf("Functional took %s", elapsed)

	start = time.Now()
	imperative()
	elapsed = time.Since(start)
	log.Printf("Functional took %s", elapsed)

	start = time.Now()
	imperative()
	elapsed = time.Since(start)
	log.Printf("Functional took %s", elapsed)
}
