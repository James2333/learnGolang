package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	//"io"
	"log"
	"net"
)

type Protocol struct {
	Length uint32
	Content []byte
}

const HEADER_LEN = 4

func Packet(content string) []byte {
	buffer := make([]byte, HEADER_LEN + len(content))
	// 将buffer前面四个字节设置为包长度，大端序
	binary.BigEndian.PutUint32(buffer[0:4], uint32(len(content)))
	copy(buffer[4:], content)
	return buffer
}

//解包，先读取4个字节转换成整形，再读包长度字节
func UnPacket(c net.Conn) (*Protocol, error) {
	var (
		p = &Protocol{}
		header = make([]byte, HEADER_LEN)
		//header头长度
	)
	_, err := io.ReadFull(c, header)//读取头部的四个字节
	if err != nil {
		return p, err
	}
	p.Length = binary.BigEndian.Uint32(header) //转换成10进制的数字
	log.Println(p.Length)
	contentByte :=make([]byte, p.Length)
	_, e := io.ReadFull(c, contentByte) //继续读取后续内容
	if e != nil {
		return p, e
	}
	p.Content = contentByte
	return p, nil
}

func (p *Protocol) parseContent() (map[string]interface{}, error) {
	var object map[string]interface{}
	unmarshal := json.Unmarshal(p.Content, &object)
	if unmarshal != nil {
		return object, unmarshal
	}
	return object, nil
}


func main() {
	l, err := net.Listen("tcp", ":9527")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			log.Println(c.RemoteAddr().String())
			protocol, _ := UnPacket(c)
			parseContent, err := protocol.parseContent()
			if (err != nil) {
			}
			s := parseContent["serviceId"].(string)
			cstr := parseContent["data"].(string)
			if s == "Hello.world" {
				fmt.Printf("serviceId: %s, content: %s", s, cstr)
				writeByte := []byte(cstr)
				c.Write(writeByte);
			}
			c.Close()
		}(conn)
	}
}