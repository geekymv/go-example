package main

import (
	"fmt"
	"time"
	"sync"
)

// if a WaitGroup is explicitly passed into functions, it should be done by pointer.
func myWorker2(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n", id)

}

func main() {

	// 使用 WaitGroup 等待多个goroutine执行完成（类似 Java 中的 CountDownLatch）
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// 捕获迭代变量，为了避免在每个 goroutine 中使用相同的 i
		i := i
		// goroutine 会推迟函数的执行时机
		go func() {
			myWorker2(i, &wg)
		}()
	}

	wg.Wait()
}
