package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func fun(x int) {
	fmt.Printf("f(%d)\n", x/x)
	defer fmt.Printf("defer %d\n", x)
	fun(x - 1)
}

func printStack() {
	// 输出函数栈信息
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println("-----------------------")
	fmt.Println("stack", string(buf[:n]))
	fmt.Println("-----------------------")
}

func main() {
	//defer printStack()
	defer debug.PrintStack()
	fun(3)
}
