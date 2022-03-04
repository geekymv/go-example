package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"
)

/*
查看文章 https://go.dev/blog/error-handling-and-go
*/

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// errorString is a trivial implementation of error.
		//return 0, errors.New("math: square root of negative number")
		// 使用 fmt.Errorf 格式化返回 error
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}

func main() {

	file, err := os.Open("1.txt")
	if err != nil {
		// 打印错误信息并退出
		log.Fatal("读取文件发生错误：", err.Error())
	}

	// 读取文件内容
	b := make([]byte, 128)
	file.Read(b)
	fmt.Println("file content", string(b))

	fmt.Printf("%T\n", file)

	// error 接口的实现
	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%T\n", err)
	fmt.Println("f = ", f)

}
