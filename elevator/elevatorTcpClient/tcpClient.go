package elevatorTcpClient

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
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

func NewTestReqEle() *ReqEl {
	return &ReqEl{
		Operation: "1",
		ReqEle: ReqEle{
			TaskIdState: "111",
			StartFloor:  1,
			EleId:       "111",
		},
	}
}
func cConnHandler(c net.Conn) {
	//返回一个拥有 默认size 的reader，接收客户端输入
	//reader := bufio.NewReader(os.Stdin)
	//缓存 conn 中的数据
	//buf := make([]byte, 1024)
	fmt.Println("请输入客户端请求数据...")
	input := NewTestReqEle()
	d := json.NewDecoder(c)
	e := json.NewEncoder(c)
	for {
		e.Encode(input)
		//客户端输入
		//inputJson, err := json.Marshal(input)
		//if err != nil {
		//	log.Printf("Json marshal faild %v", err)
		//}
		////客户端请求数据写入 conn，并传输
		//c.Write(inputJson)
		//服务器端返回的数据写入空buf
		//_, err = c.Read(buf)
		var resultJson Res
		err := d.Decode(&resultJson)
		//err=json.Unmarshal(buf,&resultJson)
		if err != nil {
			fmt.Printf("客户端解析数据失败 %s\n", err)
			continue
		}
		//回显服务器端回传的信息
		fmt.Printf("\n服务器端回复结果%v,%s", resultJson.Result, resultJson.EleId)
		time.Sleep(time.Second * 2)
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

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}