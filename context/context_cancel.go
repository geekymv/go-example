package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			// Derive a new Context for this goroutine from the Context owned by the main function.
			ctx := context.WithValue(ctx, id, id)

			// Wait until the Context is cancelled.
			<-ctx.Done()

			fmt.Println("Cancelled:", id)
		}(i)
	}

	time.Sleep(5 * time.Second)

	// Cancel the Context and any derived Context's as well.
	cancel()
	wg.Wait()
}
