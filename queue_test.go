package koa

import (
	"context"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {

	var q Queue

	q.Add(func(ctx context.Context, next func()) {
		fmt.Println(1)
		next()
		fmt.Println(3)
	})
	q.Add(func(ctx context.Context, next func()) {
		fmt.Println(2)
	})

	q.Run(context.Background())

	fmt.Println("ok")
}
