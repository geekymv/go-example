package main

import (
	"fmt"
	"sync"
)

func main() {

	list := []int{}

	ch := make(chan int, 10)

	var wg sync.WaitGroup
	//var lock sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//defer lock.Unlock()
			//lock.Lock()
			//list = append(list, i)
			ch <- i
		}(i)
	}

	wg.Wait()
	close(ch)

	for i := range ch {
		list = append(list, i)
	}

	fmt.Println(len(list))

}
