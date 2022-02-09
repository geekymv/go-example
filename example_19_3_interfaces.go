package main

import "fmt"

type Country2 struct {
	Name string
}

type City2 struct {
	Name string
}

type Stringable interface {
	ToString() string
}

func (c Country2) ToString() string {
	return "Country = " + c.Name
}

func (c City2) ToString() string {
	return "City = " + c.Name
}

// 参数类型为接口（面向接口编程）
func PrintStr(p Stringable)  {
	fmt.Println(p.ToString())
}

func main()  {
	c1 := Country2{"China"}
	c2 := City2{"Hefei"}

	PrintStr(c1)
	PrintStr(c2)
}

