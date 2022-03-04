package server

import (
	"fmt"
	c "go-example/net/client"
	"io"
	"log"
	"net"
	"sync"
)

type Server struct {
	tcpListener net.Listener
	conns       sync.Map
}

func (s *Server) New(address string) error {
	var err error
	s.tcpListener, err = net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("listen (%s) failed - %s", address, err)
	}
	return nil
}

func (s *Server) Start() {
	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			log.Fatalln(err)
			break
		}

		go s.Handle(conn)
	}
	fmt.Println("end for")

}

func (s *Server) Handle(conn net.Conn) {
	log.Println("New client ", conn.RemoteAddr())

	// read magic
	buff := make([]byte, 4)
	_, err := io.ReadFull(conn, buff)
	if err != nil {
		conn.Close()
		return
	}
	fmt.Println("magic is", string(buff))

	// TODO 判断 magic 是否合法

	client := c.NewClient(conn)
	// 保存连接
	s.conns.Store(conn.RemoteAddr(), client)

	client.IOLoop()
}
