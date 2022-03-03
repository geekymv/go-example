package main

import (
	s "go-example/net/server"
	"log"
)

func main() {
	server := &s.Server{}
	err := server.New("127.0.0.1:8088")
	if err != nil {
		log.Fatalln(err)
		return
	}
	server.Start()

}
