package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Printf("name = %s", u.Name)
}

func DoFieldAndMethod(input interface{}) {
	// 获取类型
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is :", inputType.Name())

	// 获取 value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is :", inputValue)

	fmt.Println("------------------------------")

	// 通过 type 获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过 type 遍历里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s : %v\n", method.Name, method.Type)
	}
}

func main() {
	u := User{1, "tom", 18}
	DoFieldAndMethod(u)
}
