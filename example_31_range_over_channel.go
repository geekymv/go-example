package main

import (
	"fmt"
)

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 关闭 channel
	close(queue)

	// 关闭 channel 之后，channel里面的元素继续保留
	for elem := range queue {
		fmt.Println(elem)
	}

}
