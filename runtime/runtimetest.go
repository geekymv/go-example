package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	// 获取运行时 go 版本号
	fmt.Println("go version is", runtime.Version())

	// go1.16.14
	v := runtime.Version()

	first := strings.IndexByte(v, '.')
	last := strings.LastIndexByte(v, '.')
	fmt.Println(first, last)

	// 16
	v = v[first+1 : last]
	fmt.Println(v)
}
