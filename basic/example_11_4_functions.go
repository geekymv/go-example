package main

import (
	"fmt"
)

// 可变参数
func mySum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}


func main() {

	result := mySum()
	fmt.Println(result)

	result = mySum(1, 2, 3)
	fmt.Println(result)

	val := []int{1, 2, 3, 4}
	result = mySum(val...)
	fmt.Println(result)

}
