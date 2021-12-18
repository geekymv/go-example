package main

/*
https://gobyexample.com/slices
slice内部实现 https://go.dev/blog/slices-intro

切片 value := []类型 {}
数组 value := [数组长度]类型 {}
数组与切片的区别：
1、声明数组需要指定数组长度，切片则不需要；
2、作为函数参数时，数组传递的是数组的副本，切片传递的是指针。
*/
import (
	"fmt"
)

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Printf("the s len is %d, capacity is %d\n", len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s)

	fmt.Println("len:", len(s))

	// append
	s = append(s, "d")
	// capacity 扩大到原来的2倍
	fmt.Printf("the s len is %d, capacity is %d\n", len(s), cap(s))

	s = append(s, "e", "f")
	fmt.Println("apd:", s) // [a b c d e f]

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice[low:high] 包括low，不包括high
	l := c[2:4]
	fmt.Println("sl1:", l)

	l = c[2:]
	fmt.Println("sl2:", l)

	l = c[:4]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

}
