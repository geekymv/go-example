package main

import (
	"fmt"
)

// 可变参数
func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2, 3)

	// slice
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
