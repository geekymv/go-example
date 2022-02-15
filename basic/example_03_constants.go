package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000
	fmt.Println(math.Sin(n))

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// 常量，圆周率
	fmt.Println(math.Pi)

	fmt.Println(time.Sunday)
}
