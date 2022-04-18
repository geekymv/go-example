package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	if s == "" {
		fmt.Println("empty")
	}

	strs := strings.Split("", ",")
	fmt.Println(strs)

	// Count 统计
	//str2 := `["qwww","sgxzzzzz"]`
	str2 := `[6]`
	count := strings.Count(str2, ",")
	fmt.Println("count =", count)
}
