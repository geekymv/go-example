package main

import (
	"encoding/json"
	"fmt"
)

/*
JSON 中只有以大写字母开头的公有属性才能编码和解码
*/
type response1 struct {
	Page   int
	total  int
	Fruits []string
}

// 结构体属性tag 自定义 json 中 key 的名称
type response2 struct {
	Page   int      `json:"page"`
	Total  int      `json:"total"`
	Fruits []string `json:"fruits"`
}

func main() {

	bytesBool, _ := json.Marshal(true)
	fmt.Println(string(bytesBool))

	byteStr, _ := json.Marshal("hello")
	fmt.Println(string(byteStr))

	byteArr, _ := json.Marshal([]string{"hello", "world"})
	fmt.Println(string(byteArr))

	byteMap, _ := json.Marshal(map[string]int{"hello": 1, "world": 2})
	fmt.Println(string(byteMap))

	res1 := &response1{
		Page:   1,
		total:  100,
		Fruits: []string{"Apple", "Pear"},
	}
	byteRes, _ := json.Marshal(res1)
	fmt.Println(string(byteRes))

	res2 := &response2{
		Page:   2,
		Total:  100,
		Fruits: []string{"Apple", "Pear"},
	}
	byteRes2, _ := json.Marshal(res2)
	fmt.Println(string(byteRes2))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

}
