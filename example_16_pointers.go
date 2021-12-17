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

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i 获取变量i的内存地址
	zeroptr(&i)

	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

}
