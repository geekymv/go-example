package main

import "fmt"

/*
参见 context.Context 接口的实现
*/

type A interface {
	a(int) int
}

type AImpl struct {
}

func (a *AImpl) a(i int) int {
	return i * 2
}

type B struct {
	A
}

func main() {
	/*
		b := B{
			A: new(AImpl),
		}
	*/
	b := B{
		A: &AImpl{},
	}

	r := b.a(1)
	fmt.Println("r =", r)
}
