package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	x := 1
	p := &x         // p 是整型指针，指向 x，p 包含 x 的地址
	fmt.Println(*p) // 1
	*p = 2
	fmt.Println(x) // 2

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i 获取变量i的内存地址
	zeroptr(&i)

	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

}
