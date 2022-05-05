package main

import "fmt"

type Redis struct {
	pool map[string]int
}

func main() {
	r := &Redis{
		pool: nil,
	}

	p := r.pool
	fmt.Println("p=", p)

	//p["a"] = "a"

	// map 为 nil 进行取值时，返回的是对应类型的零值，不存在也是返回零值，但增加值会报 panic。
	if a, ok := r.pool["a"]; ok {
		fmt.Println("a=", a)
	} else {
		fmt.Println("default a=", a)
	}
}
