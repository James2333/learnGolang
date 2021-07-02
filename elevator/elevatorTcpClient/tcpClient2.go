package elevatorTcpClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

func ccConnHandler(c net.Conn) {
	input := &ReqEl{
		Operation: "0",
		ReqEle: ReqEle{
			State: "0",
			Floor: 1,
			EleId: "5",
		},
	}
	//inputJson, err := json.Marshal(input)
	d := json.NewDecoder(c)
	e := json.NewEncoder(c)
	for {
		e.Encode(input)

		var msg Res
		err := d.Decode(&msg)
		if err != nil {
			log.Printf("decode service msg faild %v", err)
		}
		fmt.Printf("收到服务器返回信息%v,%s,%s\n", msg.Result, msg.EleId,msg.Error)
		//服务器端返回的数据写入空buf
		//cnt, err := c.Read(buf)
		//
		//if err != nil {
		//	fmt.Printf("客户端读取数据失败 %s\n", err)
		//	continue
		//}
		////回显服务器端回传的信息
		//fmt.Print("\n服务器端回复" + string(buf[0:cnt]))
		time.Sleep(time.Second * 10)
	}
}

func NewClientSocket2() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("客户端建立连接失败")
		return
	}
	fmt.Println("与服务端建立连接成功...")

	ccConnHandler(conn)
}
