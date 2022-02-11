package main

import (
	"fmt"
)

// 闭包（匿名函数）

func intSeq() func() int {
	i := 0
	// 返回一个函数，内部函数使用外部函数变量
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())

}
