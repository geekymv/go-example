package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// https://cloud.tencent.com/developer/article/1847917
func main() {

	var count int64

	// 无法指定pool里面可以存储的元素个数
	pool := sync.Pool{
		New: func() interface{} {
			atomic.AddInt64(&count, 1)
			fmt.Println("return")
			return make([]int, 5)
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			p := pool.Get()
			fmt.Printf("%v\n", &p)
			// 用完了 放回去
			pool.Put(p)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(count)

}
