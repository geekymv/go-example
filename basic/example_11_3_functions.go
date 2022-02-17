package main

import (
	"fmt"
)

/*
闭包
函数返回了另一个函数
*/
func squares() func() int {
	var x int

	// 匿名函数
	return func() int {
		x++
		return x * x
	}
}

func main() {

	f := squares()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
