package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	// 必须使用 2006-01-02 15:04:05 来格式化
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
