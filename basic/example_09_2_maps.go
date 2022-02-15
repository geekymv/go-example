package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := make(map[string]int)
	ages["tom"] = 18
	ages["bob"] = 22
	ages["alice"] = 20

	// range on map iterates over key/value pairs.
	for k, v := range ages {
		fmt.Printf("%s = %d\n", k, v)
	}

	/*
	delete(ages, "alice")

	for k, v := range ages {
		fmt.Printf("%s = %d\n", k, v)
	}
	*/

	fmt.Println("--------------------------")

	// 判断 key 是否存在
	name := "alice"
	if age, ok := ages[name]; ok {
		fmt.Printf("%s age = %d\n", name, age)
	} else {
		fmt.Printf("%s 不存在\n", name)
	}


	fmt.Println("--------------------------")

	// var names []string
	// 指定 slice 容量
	names := make([]string, 0, len(ages))

	// range can also iterate over just the keys of a map.
	for name := range ages {
		names = append(names, name)
	}
	//fmt.Println(names)

	// 根据 key 排序
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}


	// 大多数 map 操作都可以安全的在 map 的零值 nil 上执行
	var scores map[string]int
	fmt.Println(len(scores))
	fmt.Println(scores == nil)


	// 未初始化的 map 不能添加元素
	// scores["tom"] = 100


}
