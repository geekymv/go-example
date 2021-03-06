package main

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container 内嵌 base，类似继承
type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("alse num:", co.base.num)
	fmt.Println("describe:", co.describe())

	// 先声明再赋值
	var ct container
	ct.num = 1
	ct.str = "i am ct"
	fmt.Println(ct)

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

}
