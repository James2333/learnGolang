package elevatorTcpClient

import (
	"encoding/binary"
	"github.com/golang/glog"
	"learn101/elevator/reply"
	"learn101/elevator/session"
	"time"

	//"encoding/json"
	"fmt"
	"io"
	"learn101/elevator"
	"learn101/elevator/packet"
	"log"
	"net"
	"strings"
)




func NewTcpService() {
	//1.建立监听端口
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	fmt.Println("listen Start...:")
	//els := NewTestEls()
	fmt.Println("初始化电梯数据...")
	for {
		//2.接收客户端的链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		//3.开启一个Goroutine，处理链接
		go connSt(conn)
	}
}

func connSt(c net.Conn)  {
	in := make(chan []byte, 16)
	sess := session.NewSession(c,in)
	defer func() {
		glog.Info("disconnect:" + c.RemoteAddr().String())
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
		go func() {
			for  {
				log.Println("打印电梯信息:")
				for _,j:=range elevator.Els {
					log.Printf("ElevatorId:%s,Floor:%d,State:%s,CurrentState:%s,IsInFloor:%t\n",j.ElevatorId,j.Floor,j.State,j.CurrentState,j.IsInFloor)
				}
				time.Sleep(time.Second*10)
			}
		}()
	for {
		//此处应该先 解包识别byte[0:2]的code 然后去传入 不同的方法。
		head := make([]byte, packet.HEADER_LEN)
		_, err := io.ReadFull(c, head) //读取头部的2个字节
		if err != nil {
			log.Println(err)
		}
		code := binary.BigEndian.Uint16(head)
		ParseCode(code,sess)
	}
}
//func coonSt(c net.Conn) {
//	//in := make(chan []byte, 16)
//	//sess := session.NewSession(in)
//	defer func() {
//		log.Println("disconnect",c.RemoteAddr().String())
//		c.Close()
//	}()
//	//打印电梯信息
//	go func() {
//		for  {
//			log.Println(".................")
//			for _,j:=range elevator.Els {
//				log.Printf("ElevatorId:%s,Floor:%d,State:%s,CurrentState:%s,IsInFloor:%s\n",j.ElevatorId,j.Floor,j.State,j.CurrentState,j.IsInFloor)
//			}
//			time.Sleep(time.Second*10)
//		}
//	}()
//	for  {
//		//此处应该先 解包识别byte[0:2]的code 然后去传入 不同的方法。
//		head := make([]byte, packet.HEADER_LEN)
//		_, err := io.ReadFull(c, head) //读取头部的2个字节
//		if err != nil {
//			log.Println(err)
//		}
//		code := binary.BigEndian.Uint16(head)
//		ParseCode(code,c)
//	}
//
//}


func ParseCode(code uint16,s *session.Session)  {
	switch code {
	//
	case reply.UPDATE_ELE:
		reply.ReplyUpdateElevator(s,elevator.Els)
	case reply.CHOOSE_ELE:
		reply.ReplyRightElevator(s,elevator.Els)
	case reply.ARRIVED_START:
		reply.ReplyElevatorArriveStart(s)
	case reply.ARRIVED_END:
		reply.ReplyElevatorArriveEnd(s)
	case reply.ROBOT_In_Floor:
		reply.ReplyRobotInFloor(s,elevator.Els)
	case reply.ROBOT_OUT_Floor:
		reply.ReplyRobotOutFloor(s,elevator.Els)
	default:
		reply.ReplyError(s)
	}
}


func NewTestEls() elevator.Elevators {
	els := elevator.NewElevators()
	el1 := &elevator.Elevator{
		ElevatorId:   "1",
		Floor:        3,
		State:        "0",
		CurrentState: "1",
	}
	el2 := &elevator.Elevator{
		ElevatorId:   "2",
		Floor:        2,
		State:        "0",
		CurrentState: "1",
	}
	el3 := &elevator.Elevator{
		ElevatorId:   "3",
		Floor:        1,
		State:        "0",
		CurrentState: "0",
	}
	el4 := &elevator.Elevator{
		ElevatorId:   "4",
		Floor:        -1,
		State:        "0",
		CurrentState: "0",
	}
	els.Update(el1)
	els.Update(el2)
	els.Update(el3)
	els.Update(el4)
	return els
}

func connHandler(c net.Conn) {
	//1.conn是否有效
	if c == nil {
		log.Panic("无效的 socket 连接")
	}

	//2.新建网络数据流存储结构
	buf := make([]byte, 4096)
	//3.循环读取网络数据流
	for {
		//3.1 网络数据流读入 buffer
		cnt, err := c.Read(buf)
		//3.2 数据读尽、读取错误 关闭 socket 连接
		if cnt == 0 || err != nil {
			c.Close()
			break
		}

		//3.3 根据输入流进行逻辑处理
		//buf数据 -> 去两端空格的string
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		//去除 string 内部空格
		cInputs := strings.Split(inStr, " ")
		//获取 客户端输入第一条命令
		fCommand := cInputs[0]

		fmt.Println("客户端传输->" + fCommand)

		switch fCommand {
		case "ping":
			c.Write([]byte("服务器端回复-> pong\n"))
		case "hello":
			c.Write([]byte("服务器端回复-> world\n"))
		default:
			c.Write([]byte("服务器端回复" + fCommand + "\n"))
		}

		//c.Close() //关闭client端的连接，telnet 被强制关闭

		fmt.Printf("来自 %v 的连接关闭\n", c.RemoteAddr())
	}
}
