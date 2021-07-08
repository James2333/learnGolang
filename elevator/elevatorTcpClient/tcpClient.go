package elevatorTcpClient

import (
	"encoding/binary"
	"fmt"
	"io"
	"learn101/elevator/packet"
	"learn101/elevator/reply"
	"learn101/elevator/simlate"
	"log"
	"net"
)

/**
  client 发送端 程序
  问题：如何区分  c net.Conn 的 Write 与 Read 的数据流向?
      1. c.Write([]byte("hello"))
         c <- "hello"
      2. c.Read(buf []byte)
         c -> buf (空buf)
  客户端 和 服务器端都有 Close conn 的功能
*/
func cConnHandler(c net.Conn) {
	if c == nil {
		log.Println("conn无效")
		return
	}
	defer func() {
		log.Println("disconnect",c.RemoteAddr().String())
		c.Close()
	}()
	//这个client模拟的请求电梯操作
	log.Println("客户端建立连接成功...")
	//模拟发起一次任务
	task1:=reply.Task{
		TaskID:     "111",
		Start:      3,
		End:        5,
	}
	NewTestTask(c,task1)

	for  {
		//此处应该先 解包识别byte[0:2]的code 然后去传入 不同的方法。
		head := make([]byte, packet.HEADER_LEN)
		_, err := io.ReadFull(c, head) //读取头部的2个字节
		if err != nil {
			log.Println(err)
		}
		code := binary.BigEndian.Uint16(head)
		simlate.ParseCodeScheduling(code,c)
	}
}

func NewClientSocket() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("客户端建立连接失败")
		return
	}
	cConnHandler(conn)
}

func NewTestTask(c net.Conn,task reply.Task)  {
	b:=packet.Packet(task,reply.CHOOSE_ELE)
	c.Write(b)
}

