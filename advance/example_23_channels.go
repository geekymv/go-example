package main

import (
	"fmt"
	"time"
)

func main() {

	// 默认channel是没有缓冲的，当接收数据执行的时候才能发送数据
	message := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			// 发送ping到channel
			message <- "ping"
			fmt.Println("send:", i)
		}
	}()

	time.Sleep(time.Second * 10)

	// 会阻塞，从channel中接收数据
	msg := <-message
	fmt.Println(msg)

}
