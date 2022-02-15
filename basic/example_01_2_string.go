package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	// 字符串，字符串是不可变的字节序列
	s := "hello, world"
	fmt.Println(len(s))

	// 字符串的第i个字节不一定就是第i个字符，因为非ASCII字符的UTF-8码点需要两个字节或多个字节
	fmt.Println(s[1]) // 101 -> e

	// len 返回的是字节长度，UTF-8下一个中文占3个字节
	name := "hello 中国"
	fmt.Println(len(name))

	// 获取字符串长度
	fmt.Println(name, "的长度是：", utf8.RuneCountInString(name))
	fmt.Println(name, "的长度是：", len([]rune(name)))

	// 产生一个新字符串
	hello := s[0:5]
	fmt.Println(hello)

	world := s[7:]
	fmt.Println(world)

	// 字符串内部的数据不可修改
	// s[0] = 'H'

	// 不可变意味着两个字符串能安全地共用同一段底层内存，使得复制任何长度字符串的开销低廉


	// 字符串字面量
	v := `Go is a tool
				for managing Go source code.`
	fmt.Println(v)

	// 字符串操作的标准包
	// bytes strings strconv unicode

	// int 转字符串
	is := strconv.Itoa(100)
	fmt.Printf("%T\n", is)

	// 字符串转 int
	if i, err := strconv.Atoi("abc"); err != nil {
		fmt.Printf("转换失败：%s\n", err.Error())
	}else {
		fmt.Printf("%v\n", i)
	}

	// 判断一个字符是否为数字
	fmt.Println(unicode.IsDigit('a'))

}