package main

import "fmt"

/**
init 函数的执行顺序在其所在包的包级变量之后。
*/
func init() {
	fmt.Printf("init i =%d\n", i)
	i = 1
	fmt.Printf("init i =%d\n", i)
}

var (
	i int = setI()
)

func setI() int {
	fmt.Println("set i")
	return 2
}

func main() {
	fmt.Printf("main i =%d\n", i)
}
