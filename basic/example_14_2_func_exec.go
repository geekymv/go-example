package main

import (
	"fmt"
)

func main() {

	// 立即执行函数
	res := func(i, j int) int {
		return i + j
	}(1, 2)

	fmt.Println("res = ", res)
}
