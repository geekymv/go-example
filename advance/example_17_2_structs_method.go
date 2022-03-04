package main

import "fmt"

type Foo struct {
	val int
}

func (f *Foo) SetP(v int, callBy string) {
	f.val = v
	fmt.Printf("In SetP(): %p\n", f)
	fmt.Printf("In SetP(): call by:%s\tval:%d\n", callBy, f.val)
}

func (f Foo) SetV(v int, callBy string) {
	f.val = v
	fmt.Printf("In SetV(): %p\n", &f)
	fmt.Printf("In SetV(): call by:%s\tval:%d\n", callBy, f.val)
}

func main() {
	f := Foo{0}

	fmt.Printf("In main(): val:%d\n", f.val)
	f.SetP(1, "main")
	fmt.Printf("In main(): val:%d\n", f.val)

	fmt.Println("-------------------------------")
	f.SetV(2, "main")
	fmt.Printf("In main(): val:%d\n", f.val)

	fmt.Println("-------------------------------")
	f2 := Foo{2}
	f2.SetP(11, "main")
	f2.SetP(22, "main")

	items := []string{"a", "b"}
	for _, v := range items {
		fmt.Println(v)
	}

}
