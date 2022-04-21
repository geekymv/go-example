package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	// for range channel 会阻塞，直到 channel 关闭
	for v := range c {
		fmt.Println(v)
	}

}
