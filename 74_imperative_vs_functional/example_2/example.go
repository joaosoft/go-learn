package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var start time.Time
	var elapsed time.Duration
	start = time.Now()

	sum := Sum(1, 2, 3, 4)
	fmt.Println(sum)

	elapsed = time.Since(start)
	log.Printf("Imperative took %s", elapsed)
}

type any interface{}

func Map(operation func(p any) any, is ...any) []any {
	r := make([]any, len(is))
	for index, value := range is {
		r[index] = operation(value)
	}
	return r
}

func Reduce(operation func(a any, b any) any, is ...any) any {
	var r any
	for index, value := range is {
		if index == 0 {
			r = value
		} else {
			r = operation(r, value)
		}
	}
	return r
}

func Add(a, b any) any {
	return (a).(int) + (b).(int)
}

func Sum(i ...any) any {
	return Reduce(Add, i...)
}
