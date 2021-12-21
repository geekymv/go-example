package main

import (
	"net"
	"time"
)

func main() {

	for i := 0; i < 1000; i++ {
		conn, err := net.Dial("tcp", "10.0.37.50:8088")
		if err != nil {

		}
		conn.Write([]byte("hello"))

		time.Sleep(time.Second)
	}

	time.Sleep(time.Minute)

}
