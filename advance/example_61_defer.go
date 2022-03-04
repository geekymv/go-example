package main

import (
	"fmt"
	"os"
)

func main() {

	file := createFile("/tmp/defer.txt")

	// defer 会在方法最后执行，类似 Java 中的 finally
	defer closeFile(file)

	writeFile(file)
}

func createFile(name string) *os.File {
	fmt.Println("create file")

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	return file
}

func writeFile(f *os.File) {
	fmt.Println("write file")
	fmt.Fprintln(f, "this is data")
}

func closeFile(f *os.File) {
	fmt.Println("close file")

	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
}
