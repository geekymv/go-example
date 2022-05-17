package main

import "fmt"

func main() {
	var s []int
	fmt.Println("s=", s == nil)    // true
	fmt.Println("len(s)=", len(s)) // 0

	s = append(s, 1)
	fmt.Println(s) // [1]

	var m map[int]string
	fmt.Println("m=", m == nil)           // true
	fmt.Println("len(m)=", len(m))        // 0
	fmt.Printf("m[10001]=%q\n", m[10001]) // ""
	m[10001] = "tom"                      // panic
}
