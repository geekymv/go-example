package main

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

// 零值可用
func main() {
	// 切片零值可用
	var zeroSlice []int
	fmt.Println(zeroSlice) // []

	zeroSlice = append(zeroSlice, 1)
	zeroSlice = append(zeroSlice, 2)

	fmt.Println(zeroSlice) // [1 2]

	// String() 方法对 nil 做了判断
	var p *net.TCPAddr
	fmt.Println(p) // 等价 p.String() 输出 <nil>

	// 零值可用无锁
	var mu sync.Mutex
	fmt.Println(mu)
	mu.Lock()
	mu.Unlock()

	// 零值可用，Buffer 底层使用了零值可用的切片 []byte
	var buf bytes.Buffer
	buf.Write([]byte("hello"))
	fmt.Println(buf.String()) // hello

}
