package main

import "fmt"

type HelloPoint struct {
	X, Y int
}

// h 是形参接收者
func (h *HelloPoint) Hello() {
	fmt.Println("hello method")
}

func main() {
	// he 实参接收者
	he := HelloPoint{1, 2}

	// 编译器会对变量he进行 &he 隐式类型转换
	he.Hello()
}
