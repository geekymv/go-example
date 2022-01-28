package pack

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// 创建用户
func newUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	u := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	return u
}
