package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {

	// var x interface{} = []int{1, 2, 3}
	// fmt.Println(x == x) // panic

	flagSet := flag.NewFlagSet("hello", flag.ExitOnError)
	flagSet.String("host", "127.0.0.1", "ip address")

	ip := flagSet.Lookup("host").Value.String()
	fmt.Println("ip", ip)

	// 获取运行时 go 版本号
	fmt.Println("go version is", runtime.Version())

	//errors.Is()

}
