package main

import "fmt"

/*
切片打开了操作数组的窗口
*/
func main() {
	// 数组的切片化
	arr := [5]int{1, 2, 3, 4, 5}
	s2 := arr[1:3]
	fmt.Println(len(s2), cap(s2)) // len = 2 cap = 4

	// cap 是取决于底层数组的长度，起始位置到数组的尾部
}
