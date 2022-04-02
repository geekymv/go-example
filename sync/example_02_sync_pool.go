package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("return")
			return make([]int, 5)
		},
	}

	v := pool.Get().([]int)
	v[0] = 1
	pool.Put(v)

	// 取出来再放回去
	v1 := pool.Get()
	pool.Put(v1)

	fmt.Printf("%v, v1:%v\n", &v1, v1)
	time.Sleep(1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		v2 := pool.Get()
		fmt.Printf("%v, v2:%v\n", &v2, v2)
		wg.Done()
	}()
	wg.Wait()

}
