package client

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	ClientId string
	// 组合
	net.Conn
}

func (c *Client) IOLoop() {
	// 读取内容
	for {
		buff := make([]byte, 128)
		_, err := c.Read(buff)
		if err != nil {
			log.Println("close conn", c.RemoteAddr(), err)
			c.Close()
			break
		}
		fmt.Println("handle", string(buff))
	}
}
