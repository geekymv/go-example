package main

import "fmt"

type Foo struct {
	val int
}

func (f *Foo) SetP(v int, callBy string) {
	f.val = v
	fmt.Printf("In SetP(): call by:%s\tval:%d\n", callBy, f.val)
}

func main() {
	f := Foo{0}
	fmt.Printf("In main(): val:%d\n", f.val)
	f.SetP(1, "main")
	fmt.Printf("In main(): val:%d\n", f.val)
}
