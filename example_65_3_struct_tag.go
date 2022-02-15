package main

import (
	"reflect"
	"fmt"
)

type StudentInfo struct {
	Name string `info:"name" doc:"学生名称"`
	Age int `info:"age" doc:"学生年龄"`
}

func printTag(stu interface{}) {
	t := reflect.TypeOf(stu).Elem()

	for i := 0; i < t.NumField(); i++ {
		infoTag := t.Field(i).Tag.Get("info")
		docTag := t.Field(i).Tag.Get("doc")
		fmt.Println("infoTag = ", infoTag, "docTag = ", docTag)
	}
}

func main() {
	var stu StudentInfo
	printTag(&stu)
}


