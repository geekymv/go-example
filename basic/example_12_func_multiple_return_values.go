package main

import (
	"fmt"
)

// 多个返回值
func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)
}
