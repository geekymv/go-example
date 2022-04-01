package main

import "fmt"

type A struct {
	Name string
}

func main() {

	var a A
	//a.Name = "1"

	// 判断结构体是否为空
	if a == (A{}) {
		fmt.Println("empty")
	}

}
