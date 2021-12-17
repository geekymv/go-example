package main

import (
	"fmt"
)

func main() {

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
	
	
}