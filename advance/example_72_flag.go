package main

import (
	"flag"
	"fmt"
	"os"
)

/*
https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
*/
func main() {

	// var x interface{} = []int{1, 2, 3}
	// fmt.Println(x == x) // panic

	//wordPtr := flag.String("word", "foo", "a string foo")
	//flag.Parse()
	//fmt.Println("word:", *wordPtr)

	flagSet := flag.NewFlagSet("example", flag.ExitOnError)
	flagSet.String("host", "127.0.0.1", "ip address")
	flagSet.Bool("version", false, "print version string")

	// 解析命令行参数
	flagSet.Parse(os.Args[1:])
	ip := flagSet.Lookup("host").Value.String()
	fmt.Println("ip", ip)

	// 通过命令行参数只指定参数名 如 -version 没有指定值的情况下 bool类型会默认设置成true
	// 运行 ./example_72_flag -version
	version := flagSet.Lookup("version").Value.(flag.Getter).Get().(bool)
	fmt.Println("version", version)
	//errors.Is()

}
