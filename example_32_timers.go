package main

import (
	"time"
	"fmt"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	// channel block
	<-timer1.C
	fmt.Println("Timer 1 fired")


	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// 取消 timer
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
