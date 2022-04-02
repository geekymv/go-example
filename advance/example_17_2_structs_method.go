package main

import "fmt"

type Foo struct {
	val int
}

/*
方法接收者的值类型和指针类型
https://mp.weixin.qq.com/s/Av3DrzDXa2cjjBbtkj6wuw
https://gosamples.dev/print-address/
*/
func (f *Foo) SetP(v int, callBy string) {
	fmt.Printf("f1 addr:%p\n", f)
	f.val = v
	//fmt.Printf("In SetP(): %p\n", f)
	//fmt.Printf("In SetP(): call by:%s\tval:%d\n", callBy, f.val)

}

func (f Foo) SetV(v int, callBy string) {
	fmt.Printf("f2 addr:%p\n", &f)
	f.val = v
	//fmt.Printf("In SetV(): %p\n", &f)
	//fmt.Printf("In SetV(): call by:%s\tval:%d\n", callBy, f.val)
}

func main() {
	f := Foo{0}
	fmt.Printf("f0 addr:%p\n", &f)

	f.SetP(1, "set p")
	f.SetV(1, "set p")

	/*fmt.Printf("In main(): val:%d\n", f.val)
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
	}*/

}
