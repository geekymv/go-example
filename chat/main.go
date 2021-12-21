package main

/*
cd chat
编译
go build -o server.exe main.go server.go
运行
./server.exe

Linux 查看进程占用线程数
cat /proc/{pid}/status

*/
import(
	"github.com/geekymv/go-chat/pack"
)

func main() {
	server := pack.NewServer("0.0.0.0", 8088)
	server.Start()
}
