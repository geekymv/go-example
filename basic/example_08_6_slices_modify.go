package main

import "fmt"

type UserInfo struct {
	Name string
	Age  int
}

func main() {
	userList := make([]UserInfo, 0, 2)
	fmt.Printf("len:%v, cap:%v\n", len(userList), cap(userList))

	user1 := UserInfo{
		Name: "zhangsan",
		Age:  22,
	}
	user2 := UserInfo{
		Name: "lisi",
		Age:  20,
	}

	// append value copy
	userList = append(userList, user1, user2)

	// 不是一个地址
	fmt.Printf("user1 addr:%p\n", &user1)
	fmt.Printf("user1 addr:%p\n", &userList[0])

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

	fmt.Println("------------------------------------")

	var userList2 []*UserInfo
	fmt.Println(user1.Age)
	userList2 = append(userList2, &user1, &user2)
	for _, user := range userList2 {
		user.Age = user.Age + 2
	}
	for _, user := range userList2 {
		fmt.Println(*user)
	}

}
