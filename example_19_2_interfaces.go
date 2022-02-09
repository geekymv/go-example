package main

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

// Country 和 City 都实现了 Printable 接口的 PrintStr 方法，Go 中可以只实现接口中的部分方法
func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

func (c City) PrintStr()  {
	fmt.Println(c.Name)
}


func main() {

	c1 := Country{"China"}
	c2 := City{"Hefei"}

	c1.PrintStr()
	c2.PrintStr()
}

