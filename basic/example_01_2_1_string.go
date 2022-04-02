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
}
