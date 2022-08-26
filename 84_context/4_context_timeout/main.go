package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	timeout := 1500 * time.Millisecond
	ctx, cancelCtx := context.WithTimeout(ctx, timeout)
	defer func(f context.CancelFunc) {
		fmt.Println("running cancel context")
		f()
	}(cancelCtx)

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}
