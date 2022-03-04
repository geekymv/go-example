package main

import (
	"fmt"
	"strings"
)

// 查找字符串 t 在切片 vs 中的索引
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0, 10)

	//fmt.Println(len(vsf))
	//fmt.Println(cap(vsf))

	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"an", "any", "cool", "money"}
	fmt.Println(Index(strs, "cool"))

	fmt.Println(Include(strs, "any"))

	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "an")
	}))

	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "an")
	}))

	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "an")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
