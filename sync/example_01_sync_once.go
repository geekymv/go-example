package main

import (
	"fmt"
	"sync"
)

type program struct {
	once sync.Once
}

func (p *program) Close() error {
	p.once.Do(func() {
		// 保证只执行一次
		fmt.Println("close method")
	})
	return nil
}

func main() {
	p := &program{}
	p.Close()
	p.Close()
	p.Close()
}
