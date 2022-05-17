package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name string
	Age  int
}

func handleUsers1(users []User) {
	var wg sync.WaitGroup
	for _, u := range users {
		fmt.Printf("u:=%p\n", &u)
		wg.Add(1)
		go func() {
			//time.Sleep(1)
			fmt.Println("u=", u)
			wg.Done()
		}()
	}
	wg.Wait()
}

func handleUsers2(users []User) {
	var wg sync.WaitGroup
	for _, u := range users {
		wg.Add(1)
		go func(u User) {
			time.Sleep(1)
			fmt.Println("u=", u)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

func handleUsers3(users []User) {
	var wg sync.WaitGroup
	for _, v := range users {
		wg.Add(1)
		v := v
		go func() {
			//time.Sleep(1)
			fmt.Println("v=", v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	users := make([]User, 0, 10)
	users = append(users, User{
		Name: "tom",
		Age:  20,
	})
	users = append(users, User{
		Name: "tom",
		Age:  30,
	})

	// for 中的 users 是副本，u 也是副本
	for _, u := range users {
		u.Age += 1
	}
	fmt.Printf("users:%#v\n", users)
	// users:[]main.User{main.User{Name:"tom", Age:20}, main.User{Name:"tom", Age:30}}

	for i, _ := range users {
		users[i].Age += 1
	}
	fmt.Printf("users:%#v\n", users)
	// users:[]main.User{main.User{Name:"tom", Age:21}, main.User{Name:"tom", Age:31}}
}
