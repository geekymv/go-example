package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
func a() (int, error) {
	fmt.Println("a1...")
	//return errors.New("error a")
	time.Sleep(1 * time.Second)
	fmt.Println("a2...")
	return 1, nil
}

func b() (int, error) {
	fmt.Println("b1...")
	time.Sleep(10 * time.Second)
	//return errors.New("error b")
	fmt.Println("b2...")
	return 2, nil
}

func main() {
	exitCh := make(chan error)
	var once sync.Once

	exitFunc := func(i int, err error) {
		fmt.Println("exit func i =" + strconv.Itoa(i))
		if i == 1 {
			time.Sleep(5 * time.Second)
		}
		once.Do(func() {
			if err != nil {
				fmt.Println("err...", err)
			}
			exitCh <- err
		})
	}

	wg := WaitGroupWrapper{}

	wg.Wrap(func() {
		exitFunc(a())
	})

	wg.Wrap(func() {
		exitFunc(b())
	})

	err := <-exitCh
	fmt.Println("exit err", err)
}
