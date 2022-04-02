package main

type File struct {
	fd      int
	name    string
	dirinfo string
	nepipe  int
}

/*
new 的作用
It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory,
it only zeros it. That is, new(T) allocates zeroed storage for a new item of type T and returns its address,
a value of type *T. In Go terminology, it returns a pointer to a newly allocated zero value of type T.

make 作用于 slice map channel

*/

/*
new(File) 等价于 &File{}
*/
func NewFile0(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	return f
}

func NewFile1(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	// 这种方式所有字段都需要指定
	f := File{fd, name, "", 0}
	return &f
}

func NewFile2(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	/*
		by labeling the elements explicitly as field:value pairs,
		the initializers can appear in any order,
		with the missing ones left as their respective zero values
	*/
	return &File{fd: fd, name: name}
}

func main() {

}
