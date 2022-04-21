package main

import "fmt"

func main() {

	s := []int{1, 2, 3}

	l := len(s)
	fmt.Println(l)

	s2 := s[2:3]
	fmt.Println(s2)
}
