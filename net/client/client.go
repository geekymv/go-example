package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const defaultBufferSize = 16 * 1024

type Client struct {
	ClientId string
	// 组合
	net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func NewClient(conn net.Conn) *Client {
	c := &Client{
		ClientId: "1",
		Conn:     conn,
		Reader:   bufio.NewReaderSize(conn, defaultBufferSize),
		Writer:   bufio.NewWriterSize(conn, defaultBufferSize),
	}
	return c
}

func (c *Client) IOLoop() {
	// 读取内容
	for {
		line, err := c.Reader.ReadSlice('\n')
		if err != nil {
			log.Println("close conn", c.RemoteAddr(), err)
			c.Close()
			break
		}
		// trim \n
		n := line[len(line)-1]
		fmt.Println(n == '\n')

		line = line[:len(line)-1]
		fmt.Println("handle", string(line))
	}
}
