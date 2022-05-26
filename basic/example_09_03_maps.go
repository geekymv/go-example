package main

import "fmt"

func main() {

	m1 := make(map[string]map[string]struct{})
	v1, ok := m1["1"]
	if !ok {
		// map 延迟初始化
		v1 = make(map[string]struct{})
		m1["1"] = v1
	}
	v1["a"] = struct{}{}

	//v1, ok = m1["1"]
	v1["b"] = struct{}{}

	fmt.Println("m1", m1)
}
