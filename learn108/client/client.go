package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

const HEADER_LEN  = 4

type Content struct {
	ServiceId string `json:"serviceId"`
	Data interface{} `json:"data"`
}


func Packet(serviceId string, content string) []byte {
	bytes, _ := json.Marshal(Content{ServiceId:serviceId, Data:content})
	buffer := make([]byte, HEADER_LEN + len(bytes))
	// 将buffer前面四个字节设置为包长度，大端序
	binary.BigEndian.PutUint32(buffer[0:4], uint32(len(bytes)))
	//设置前四个字节的内容
	copy(buffer[4:], bytes)
	//copy 4:之后的内容变为bytes的内容。
	return buffer
}

func main() {
	conn, e := net.Dial("tcp", ":9527")
	if e != nil {
		log.Fatal(e)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	//buffer := new(bytes.Buffer)
	buffer := Packet("Hello.world", text)
	conn.Write(buffer)


	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

	defer conn.Close()
}