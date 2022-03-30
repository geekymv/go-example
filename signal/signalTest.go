package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
cd signal
go build .
./signal 运行
*/
func main() {
	signalChan := make(chan os.Signal, 1)
	// Ctrl + C or kill -15 pid
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// 获取当前进程ID
	pid := os.Getpid()
	fmt.Printf("process id %d\n", pid)

	for {
		select {
		case s := <-signalChan:
			fmt.Printf("signal %v\n", s)
			goto stop
		}
	}

stop:
	fmt.Println("exist...")
}
