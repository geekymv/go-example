package main

import (
	"fmt"
	"time"
)

/*
带缓冲的 channel
当 channel 已满，再向里面写数据，就会阻塞
当 channel 已空，从里面取数据，也会阻塞
*/
func main() {

	// 带缓冲的channel，缓冲2个值
	messages := make(chan string, 2)

	go func() {
		defer fmt.Println("child goroutine exit")

		for i := 0; i < 3; i++ {
			messages <- fmt.Sprintf("%s%d", "hello-", i)
			fmt.Println("message len is ", len(messages), ", cap is ", cap(messages))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println(<- messages)
	}

	fmt.Println("main exit")

}
