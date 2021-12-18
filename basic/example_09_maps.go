package main

import (
	"fmt"
)


func printUsers(users map[string]int) {
	// map是引用传递
	for k, v := range users {
		fmt.Println("k = ", k, ", v = ", v)
	}
}

func main() {

	var mp map[string]int
	if mp == nil {
		fmt.Println("mp is nil")
	}

	/*
	在使用map前需要先使用make分配数据空间
	The make built-in function allocates and initializes an object of type slice, map, or chan (only).
	*/
	mp = make(map[string]int)
	mp["k1"] = 1

	// make(map[key-type]val-type)
	m := make(map[string]int)

	// name[key] = val
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	//delete(m, "k2")
	//fmt.Println("map:", m)

	/*
		判断key是否存在
		第一个返回值是key对应的value，使用下划线_表示忽略返回值
		第二个返回值表示key是否存在
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 声明并初始化map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	users := map[string]int{
		"baby": 100,
		"jack": 90,
		"bob": 98, // 最后一行要有逗号
	}
	fmt.Println("users:", users)

	// 遍历
	printUsers(users)


	// 删除
	delete(users, "bob")
	fmt.Println("users:", users)


}
