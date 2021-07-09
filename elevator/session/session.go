package session

import "net"

type Session struct {
	Ch chan []byte
	C net.Conn
}
//实现类似注册的功能，每条连接都会新生成一个session。
//想向其他的连接里面写值，就获取他的session，然后向channel里面写值
func NewSession(c net.Conn,ch chan []byte) *Session {
	s := &Session{}
	s.Ch = ch
	s.C=c
	return s
}
