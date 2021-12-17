package main

import (
	"fmt"
)

func main() {

	// 带缓冲的channel，缓冲2个值
	messages := make(chan string, 2)
	messages <- "ping"
	messages <- "hello"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
