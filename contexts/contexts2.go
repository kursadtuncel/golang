package main

import (
	"context"
	"fmt"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "abc123")

	F1(ctx)
}

func F1(ctx context.Context) {
	fmt.Println("f1", ctx.Value("correlation-id"))
	F2(ctx)
}
func F2(ctx context.Context) {
	fmt.Println("f2", ctx.Value("correlation-id"))
	F3(ctx)
}
func F3(ctx context.Context) {
	fmt.Println("f3", ctx.Value("correlation-id"))
}
