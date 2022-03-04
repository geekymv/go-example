package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println("received", msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")

	}
}
