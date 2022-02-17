package main

import (
	"sort"
	"fmt"
)

type StringSlice []string

// 实现了 sort.Interface 接口
func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool{
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	names := []string{"tom", "jack", "lily"}
	sort.Sort(StringSlice(names))

	fmt.Println(names)
}



