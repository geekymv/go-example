package main

import "os"

func main() {
	//panic("a problem")

	file, err := os.Create("/tmp/1.txt")

	if err != nil {
		panic(err)
	}

	file.Close()
}