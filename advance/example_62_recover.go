package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	// Catch the panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// This code will not run
	fmt.Println("After mayPanic")
}
