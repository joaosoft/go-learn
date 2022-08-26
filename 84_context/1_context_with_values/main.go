package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}
