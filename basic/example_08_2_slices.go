package main

import (
	"fmt"
	"strings"
)

/*
Go 中数组参数是值传递，不会修改原始数组
*/
func updateArr(names [4]string) {
	for i, v := range names {
		names[i] = strings.ToUpper(v)
	}
}

/*
Slice 包含了指向底层数组元素的指针，会修改原始数组
*/
func updateSlice(names []string) {
	for i, v := range names {
		names[i] = strings.ToUpper(v)
	}
}


func main() {

	names := []string{"tom", "bob", "jack"}
	fmt.Println(names)
	fmt.Printf("len is %d and cap is %d\n", len(names), cap(names))
	names = append(names, "kitty")
	fmt.Printf("len is %d and cap is %d\n", len(names), cap(names))

	updateSlice(names)
	fmt.Println(names)

	namesArr := [4]string{"tom", "bob", "jack", "kitty"}
	updateArr(namesArr)
	fmt.Println(namesArr)


	s := []string{0:"a", 1:"2"}
	fmt.Printf("%T, %v\n", s, s)

}
