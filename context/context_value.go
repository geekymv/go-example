package main

import (
	"context"
	"fmt"
	"sync"
)

func GET(ctx context.Context, k string) {
	if v, ok := ctx.Value(k).(string); ok {
		fmt.Printf("%s v is %s\n", k, v)
	}
}

// context with value
func main() {
	ctx := context.WithValue(context.Background(), "id", "123456")
	GET(ctx, "id")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		ctx2 := context.WithValue(ctx, "name", "tom")
		GET(ctx2, "id")
		go func() {
			GET(ctx2, "name")
			wg.Done()
		}()

		wg.Done()
	}()

	wg.Wait()
}
