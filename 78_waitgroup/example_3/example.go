package main

import (
	"fmt"
	"sync"
	"time"
)

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		duration := millisecs * time.Millisecond
		time.Sleep(duration)
		fmt.Println("Function in background, duration:", duration)
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	dosomething(200, &wg)
	dosomething(400, &wg)
	dosomething(150, &wg)
	dosomething(600, &wg)
	wg.Wait()
	fmt.Println("Done")
}
