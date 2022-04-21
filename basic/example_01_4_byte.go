package main

import (
	"bytes"
	"fmt"
)

func main() {
	split := bytes.Split([]byte("hello\nworld\nhaha"), []byte("\n"))
	fmt.Println(split)

	b1 := []byte("a b c d e")
	// [][]byte
	b2 := bytes.Split(b1, []byte(" "))
	fmt.Println(b2[0])
}
