package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// 函数当做参数传递
func opFunc(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {

	// 函数变量
	addFunc := add

	fmt.Printf("%T\n", addFunc)

	result := addFunc(1, 2)
	fmt.Println(result)

	// 函数变量，可以将函数当做参数进行传递
	opResult := opFunc(2, 3, add)
	fmt.Println(opResult)

	opResult = opFunc(2, 3, sub)
	fmt.Println(opResult)

}
