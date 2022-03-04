package main

import (
	"fmt"
)

/*
指定方向的channel
当使用 channel 作为函数参数的时候，可以指定 channel 只能发送或接收数据，这种方式增加了类型安全。
*/
func ping(pings chan<- string, msg string) {
	// pings 只能接收数据
	pings <- msg
	// fmt.Println(<-pings) // 如果试着发送数据，则在编译时就会报错
}

func pong(pings <-chan string, pongs chan<- string) {
	// pings <-chan string 只能发送数据
	msg := <-pings
	// pongs chan<- string 只能接收数据
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello")
	pong(pings, pongs)

	fmt.Println(<-pongs)

}
