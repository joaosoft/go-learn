/*
You work at a Banking Application and have created 2 functions:
	one that deposits a value into an account and another that withdraws a value from the account.

You want to simulate many deposits and withdraws that take place simultaneously and start some goroutines.

During testing you notice that a date race occurred.

Your task is to change the code in order to protect the account's balance from simultaneously writing using a mutex.

This is the initial program that has errors:

package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup) {
	*b += n
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup) {
	*b -= n
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg)
		go withdraw(&balance, i, &wg)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}

*/

package main

import (
	"fmt"
	"sync"
)

// pass pointer to sync.Mutex (mutex is a value type)
func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b -= n
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	// the final balance value be always 100
	fmt.Println("Final balance value:", balance)
}
