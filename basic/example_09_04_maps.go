package main

import "fmt"

func main() {

	m1 := make(map[string][]string)

	v1, ok := m1["a"]
	if !ok {
		//v1 = []string{}
	}
	v1 = append(v1, "1")
	m1["a"] = v1

	fmt.Println(m1)
}
