package main

import (
	"learn101/elevator/elevatorTcpClient"
)

type ReqEl struct {
	//鉴别更新还是获取的操作由operation识别
	Operation string //传值0为更新 1为请求电梯
	ReqEle
}

type ReqEle struct {
	//传前者是请求电梯，后者是更新电梯
	TaskId_state string //taskId or state
	Start_floor  int64  //strat or floor
	EleId        string //eleId or eleId
}

type Res struct {
	Result bool
	EleId  string
}

//func cConnHandler(c net.Conn) {
//	//返回一个拥有 默认size 的reader，接收客户端输入
//	//reader := bufio.NewReader(os.Stdin)
//	//缓存 conn 中的数据
//	//buf := make([]byte, 1024)
//	//fmt.Println("正在解析客户端请求数据...")
//	//input:="elevator 1 3"
//	//input:="update 5 3 0"
//	input := &ReqEl{
//		Operation: "0",
//		ReqEle: ReqEle{
//			TaskId_state: "0",
//			Start_floor:  1,
//			EleId:        "5",
//		},
//	}
//	//inputJson, err := json.Marshal(input)
//	d := json.NewDecoder(c)
//	e := json.NewEncoder(c)
//	for {
//		//客户端输入
//		//input, _ := reader.ReadString('\n')
//		//去除输入两端空格
//		//buf := make([]byte, 1024)
//		//input = strings.TrimSpace(input)
//		//客户端请求数据写入 conn，并传输
//		//c.Write(inputJson)
//		e.Encode(input)
//
//		var msg Res
//		err := d.Decode(&msg)
//		if err != nil {
//			log.Printf("decode service msg faild %v", err)
//		}
//		fmt.Printf("收到服务器返回信息%v,%s\n", msg.Result,msg.EleId)
//		//服务器端返回的数据写入空buf
//		//cnt, err := c.Read(buf)
//		//
//		//if err != nil {
//		//	fmt.Printf("客户端读取数据失败 %s\n", err)
//		//	continue
//		//}
//		////回显服务器端回传的信息
//		//fmt.Print("\n服务器端回复" + string(buf[0:cnt]))
//		time.Sleep(time.Second * 10)
//	}
//}
//
//func NewClientSocket() {
//	conn, err := net.Dial("tcp", "127.0.0.1:20000")
//	if err != nil {
//		fmt.Println("客户端建立连接失败")
//		return
//	}
//	fmt.Println("与服务端建立连接成功...")
//
//	cConnHandler(conn)
//}
func main() {
	elevatorTcpClient.NewClientSocket2()
}
