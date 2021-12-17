package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	// range 返回两个值，第一个值表示slice下标，第二个表示slice下标对应的值
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	
	for k := range kvs {
		fmt.Println("key:", k)
	}
	
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}