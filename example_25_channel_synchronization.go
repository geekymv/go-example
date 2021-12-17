package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 5)
	fmt.Println("done")
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	// 等待worker完成
	<-done
}