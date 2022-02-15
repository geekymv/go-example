package main

import (
	"fmt"
	"os"
)

func main() {
	var a = "initial"
	fmt.Println(a)

	// 一次声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 类型推断
	var d = true
	fmt.Println(d)

	// zero-valued
	var e int
	fmt.Println(e)

	// var f string  = "apple"
	// := 短变量声明，这种比较常用
	f := "apple"
	fmt.Println(f)


	file, err := os.Open("1.txt")
	fmt.Println(file)
	fmt.Println(err)

	// 短变量声明最少声明一个新变量，否则代码编译将无法通过
	// file, err := os.Open("2.txt")
	file, err = os.Open("2.txt")



}
