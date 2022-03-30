package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	now := time.Now()
	// 必须使用 2006-01-02 15:04:05 来格式化
	TIME_FORMAT := "2006-01-02 15:04:05"
	fmt.Println(now.Format(TIME_FORMAT))

	t := time.Now()
	fmt.Println(t.Year(), t.Month(), t.Day())

	var i0 int = 0   // 8
	var i1 int32 = 1 // 4
	var i2 int64 = 2 // 8
	fmt.Println(unsafe.Sizeof(i0))
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))

	// 日期格式转换
	t2 := time.Now().Unix()
	tf := time.Unix(t2, 0).Format(TIME_FORMAT)
	fmt.Println(tf)
}
