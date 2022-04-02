package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	duration := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	// Not doing this will cause memory leaks.
	defer cancel()

	// Create a channel to receive a signal that work is done.
	ch := make(chan bool, 1)

	// Ask the goroutine to do some work for us.
	go func() {
		time.Sleep(2 * time.Second)
		//time.Sleep(10 * time.Second)
		ch <- true
	}()

	// wait for the work to finish.
	// If it takes too long, move on.
	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

}
