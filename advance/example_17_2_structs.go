package main

import "fmt"

type MyPoint struct {
	X int
	Y int
}

func newPoint(x, y int) *MyPoint {
	return &MyPoint{X: x, Y: y}
}

func main() {

	point := newPoint(100, 100)

	fmt.Printf("%v", point)

}
