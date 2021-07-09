package elevatorTcpClient

import (
	"encoding/binary"
	"fmt"
	"io"
	"learn101/elevator/packet"
	"learn101/elevator/session"
	"learn101/elevator/simlate"
	"log"
	"net"
)

func ccConnHandler(c net.Conn) {
	in := make(chan []byte, 16)
	sess := session.NewSession(c,in)
	if c == nil {
		log.Println("conn无效")
		return
	}
	defer func() {
		log.Println("disconnect",c.RemoteAddr().String())
		c.Close()
	}()
	go func() {
		for {
			select {
			case msg := <-in:
				c.Write(msg)
			}
		}
	}()
	simlate.ReqUpdateElevator(sess)//触发一次更新操作
	for  {
		//此处应该先 解包识别byte[0:2]的code 然后去传入 不同的方法。
		head := make([]byte, packet.HEADER_LEN)
		_, err := io.ReadFull(c, head) //读取头部的2个字节
		if err != nil {
			log.Println(err)
		}
		code := binary.BigEndian.Uint16(head)
		simlate.ParseCodeElevator(code,sess)
	}
}

func NewClientSocket2() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("客户端建立连接失败")
		return
	}
	log.Println("与服务端建立连接成功...")

	ccConnHandler(conn)
}
