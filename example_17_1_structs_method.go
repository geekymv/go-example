package main

import (
	"fmt"
)

// 类名、属性名、方法大小写，大写表示公有的，其他包可以访问，小写表示私有的，本包可以访问
type Student struct {
	name string
	age  int
}

func (s Student) getName() string {
	return s.name
}

func (s Student) setName(name string) {
	// s 是调用该方法对象的一个拷贝
	s.name = name
}

// 成员函数
func (s *Student) setName2(name string) {
	s.name = name
}

// 方法名大写，表示方法是公有的
func (s Student) Show() {
	fmt.Println("s name is ", s.name)
	fmt.Println("s age is ", s.age)
}

func main() {

	s := Student{name: "tony", age: 22}
	fmt.Println(s.getName())

	s.setName("lisi")
	fmt.Println(s.getName())

	s.setName2("tim")
	fmt.Println(s.getName())

}
