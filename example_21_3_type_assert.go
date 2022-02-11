package main

import (
	"fmt"
	"encoding/json"
)

/*
https://go.dev/ref/spec#Type_assertions
https://blog.csdn.net/wuhuimin521/article/details/83066417
*/
func main() {

	var x interface{} = 1
	i := x.(int)
	fmt.Println("i = ", i)

	// type assertion
	str := "hello"
	serr, ok := interface{}(str).(string)
	fmt.Println("ok = ", ok)
	fmt.Println("serr = ", serr)

	err := &json.SyntaxError{}
	msg, ok := interface{}(err).(*json.SyntaxError)
	fmt.Println("msg = ", msg)
	fmt.Println("ok = ", ok)
}
