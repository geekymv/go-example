package main

import "fmt"

/*
切片扩容
*/
func main() {
	u1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("u1 len = %v cap = %v\n", len(u1), cap(u1)) // 5 5

	u2 := u1[0:4]
	fmt.Printf("u2 len = %v cap = %v\n", len(u2), cap(u2)) // 4 5
	fmt.Println(u1)                                        // 1 2 3 4 5
	fmt.Println(u2)                                        // 1 2 3 4

	// u1 与 u2 底层共用一个数组
	u2 = append(u2, 6)
	fmt.Printf("u2 len = %v cap = %v\n", len(u2), cap(u2)) // 5 5
	fmt.Println(u1)                                        // 1 2 3 4 6
	fmt.Println(u2)                                        // 1 2 3 4 6

	// u2 扩容，底层创建了新的数组
	u2 = append(u2, 7)
	fmt.Printf("u2 len = %v cap = %v\n", len(u2), cap(u2)) // 6 10

	// u1 与 u2 不再共用一个底层数组
	u2[0] = 11
	fmt.Println(u1) // 1 2 3 4 6
	fmt.Println(u2) // 1 2 3 4 6 7
}
