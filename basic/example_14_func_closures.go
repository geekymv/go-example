package main

import (
	"fmt"
)

// 闭包（匿名函数）

func intSeq() func() int {
	i := 0
	// 返回一个函数
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
