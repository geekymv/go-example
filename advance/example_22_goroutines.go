package main

import (
	"fmt"
	"time"
)

func ff(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	ff("direct")
	go ff("goroutine")

	// 匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 5)

	fmt.Println("done")
}
