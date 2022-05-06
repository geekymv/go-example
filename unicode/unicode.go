package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func main() {
	ascii := strconv.QuoteToASCII("马")
	fmt.Println(ascii)

	fmt.Println("max int16=", math.MaxUint8)

	str := "中国人"
	// 返回字符串的字节数
	fmt.Println(len(str))
	fmt.Println(str[0:3])
	fmt.Println("---------------------------")

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c\n", r)
		i += size
	}

	fmt.Println("---------------------------")
	// %q 带引号字符串或字符
	for _, s := range str {
		fmt.Printf("%c - %q\n", s, s)
	}

	fmt.Println("---------------------------")
	// https://studygolang.com/articles/26235
	// 通过 tune 处理 unicode 字符
	ss := []rune(str)
	fmt.Println("len ss=", len(ss))
	for i := 0; i < len(ss); i++ {
		fmt.Println(string(ss[i]))
	}

	var a1 rune = '🦍'
	a2 := '🦍'
	// %c 字符(unicode 码点)
	fmt.Printf("%c - %U - %s\n", a1, a1, reflect.TypeOf(a1))
	fmt.Printf("%c - %U - %s\n", a2, a2, reflect.TypeOf(a2))

	// Unicode Code Point
	var a3 rune = '世'
	fmt.Println(a3, reflect.TypeOf(a3))
}
