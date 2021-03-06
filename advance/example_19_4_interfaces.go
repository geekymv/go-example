package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const debugf = false

func main() {

	// 获取接口值的动态类型
	var w io.Writer
	fmt.Printf("%T\n", w) // nil
	//w.Write([]byte("hello")) // panic

	w = os.Stdout
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)

	//var buf *bytes.Buffer
	var buf io.Writer
	if debugf {
		buf = new(bytes.Buffer)
	}

	f(buf)

}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("hello"))
	}
}
