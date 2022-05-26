package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	is := []int{10, 1, 100, 22}
	//sort.Ints(i)
	sort.Slice(is, func(i, j int) bool {
		return is[i] > is[j]
	})
	fmt.Println(is)

	s := []string{"10", "1", "12", "100"}
	//sort.Strings(s)
	sort.Slice(s, func(i, j int) bool {
		s1, _ := strconv.Atoi(s[i])
		s2, _ := strconv.Atoi(s[j])
		return s1 > s2
	})
	fmt.Println(s)
}
