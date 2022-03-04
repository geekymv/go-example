package main

import (
	"fmt"
)

// int 的别名
type myInt int

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 21
	return &p
}

func printPerson(p Person) {
	// 传递一个Person的副本
	fmt.Println("p name is ", p.name)
	fmt.Println("p age is ", p.age)
}

func changeName(p *Person) {
	// 传递一个指针
	p.name = "lisi"
}

func main() {

	var i myInt = 10
	fmt.Printf("type of i is %T\n", i)

	var p1 Person
	p1.name = "Tom"
	p1.age = 22
	printPerson(p1)

	fmt.Println("================")

	changeName(&p1)
	printPerson(p1)

	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})

	fmt.Println(newPerson("Tony"))

	// 初始化结构体
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 22
	fmt.Println(sp.age)
	fmt.Println(s.age)

}
