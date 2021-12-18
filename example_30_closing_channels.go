package main

import (
	"fmt"
	"time"
)


func main() {

	jobs := make(chan int, 5)

	// 不带缓冲的channel，会阻塞
	done := make(chan bool)

	go func() {
		for {
			// jobs channel 关闭的时候，more返回false
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)

	for j := 1; j <= 3; j++ {
		fmt.Println("send job", j)
		jobs <- j
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

}
