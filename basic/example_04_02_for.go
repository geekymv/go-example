package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// 变量在 for range 的每次循环中都会被重用，而不是重新声明
	/*
		for i := 0; i < 10; i++ {
			// i 是同一块内存地址
			fmt.Printf("%v, %v\n", &i, i)
		}
	*/

	wg := sync.WaitGroup{}
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		wg.Add(1)

		/*
			// i, v := i, v // 解决方式一
			go func() {
				time.Sleep(time.Second * 3)
				// i, v 在整个循环中是重用的
				fmt.Println(i, v)
				wg.Done()
			}()
		*/

		// 解决方式二
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
			wg.Done()
		}(i, v)
	}
	wg.Wait()

}
