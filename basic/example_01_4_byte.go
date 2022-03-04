package main

import (
	"bytes"
	"fmt"
)

func main() {
	split := bytes.Split([]byte("hello\nworld\nhaha"), []byte("\n"))
	fmt.Println(split)
}
