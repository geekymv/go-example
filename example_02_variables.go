package main

import "fmt"

func main() {
	var a ="initial"
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
	// := 这种比较常用
	f := "apple"
	fmt.Println(f)
	
}