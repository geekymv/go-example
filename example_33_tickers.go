package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Tick is done")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()

	done <- true
	fmt.Println("Ticker stopped")


}
