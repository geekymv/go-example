package main

import (
	"fmt"
	"math"
)

type Po struct {
	X, Y float64
}

func (p Po) Distance(q Po) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Po{1, 1}
	q := Po{4, 5}

	// 方法表达式
	f := Po.Distance

	fmt.Println(f(p, q))
}
