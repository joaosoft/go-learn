// testes

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_mutex()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
// MUTEX
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	fmt.Println("Inc", key)
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	fmt.Println("Value", key)
	defer c.mux.Unlock()
	return c.v[key]
}
func _mutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(strconv.Itoa(i))
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
