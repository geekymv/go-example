package main

import (
	"fmt"
)

/*
使用select default 语句实现非阻塞发送、接收
*/
func main() {

	// 没有缓冲的channel
	messages := make(chan string)

	// 非阻塞接收
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞发送
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")

	}

}
