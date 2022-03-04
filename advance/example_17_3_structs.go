package main

import "fmt"

/*
匿名结构体
在 Go 中，组合是面向对象编程方式的核心
*/

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point  // 匿名结构体
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {

	var w Wheel
	w.X = 100
	w.Y = 100
	w.Radius = 10
	w.Spokes = 20

	w2 := Wheel{
		Circle: Circle{
			Point:  Point{100, 100},
			Radius: 100,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w2)

}
