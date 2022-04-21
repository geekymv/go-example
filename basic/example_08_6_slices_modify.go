package main

import "fmt"

type UserInfo struct {
	Name string
	Age  int
}

func main() {
	var userList []UserInfo
	userList = append(userList, UserInfo{
		Name: "zhangsan",
		Age:  22,
	})

	userList = append(userList, UserInfo{
		Name: "lisi",
		Age:  20,
	})

	// 用户信息没有改掉
	for _, user := range userList {
		// user 是值拷贝
		user.Age = user.Age + 1
		// %p取地址
		fmt.Printf("range %p\n", &user)
	}
	fmt.Printf("%v\n", userList)

	// 通过索引找到 user
	for i, user := range userList {
		userList[i].Age = user.Age + 1
		fmt.Printf("index %p\n", &userList[i])
	}
	fmt.Printf("%v\n", userList)
}
