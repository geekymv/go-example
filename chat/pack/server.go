package pack

import (
	"net"
	"fmt"
	"os"
)

type Server struct {
	Ip string
	Port int
}

func NewServer(ip string, port int) *Server {
	s := &Server{
		Ip: ip,
		Port: port,
	}
	return s
}

func (server *Server) Handler(conn net.Conn) {
	fmt.Println("handler...", conn.RemoteAddr())

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		conn.Write([]byte("Message received."))
	}

}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("Error listening :", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Listening on %s : %d\n", server.Ip, server.Port)
	defer listener.Close()

	for {
		fmt.Println("accept connect...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go server.Handler(conn)
	}
}

