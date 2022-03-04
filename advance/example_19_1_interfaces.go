package main

import (
	"fmt"
	"math"
)

/*
怎样使用接口
https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
*/

// Interfaces are named collections of method signatures.
// 接口是方法签名的集合
type geometry interface {
	area() float64
	perim() float64
}

// 矩形
type rectangle struct {
	width, height float64
}

// 圆
type circle struct {
	radius float64
}

/*
To implement an interface in Go, we just need to implement all the methods in the interface.
在Go语言中为了实现一个接口，我们只需要实现接口中所有的方法
*/

// 计算矩形面积
func (r rectangle) area() float64 {
	return r.width * r.height
}

// 计算矩形周长
func (r rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

// 计算圆的面积
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 计算圆的周长
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 测量方法，方法参数为接口类型
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// interface{} 是万能数据类型
func print(args interface{}) {

	// 类型断言，判断参数 args 的数据类型
	value, ok := args.(string)
	if ok {
		fmt.Println("args 的类型是 string", value)
	}

	fmt.Println(args)
}

func main() {

	var g geometry
	g = rectangle{width: 1, height: 2}
	//g = &rectangle{width: 1, height: 2}
	fmt.Printf("%T\n", g)
	measure(g)

	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	print("hello")

}
