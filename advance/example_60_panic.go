package main

import "os"

func main() {
	//panic("a problem")

	// 在 if 块中处理错误
	if file, err := os.Create("/tmp/1.txt"); err != nil {
		panic(err)
	} else {
		file.Close()
	}

}
