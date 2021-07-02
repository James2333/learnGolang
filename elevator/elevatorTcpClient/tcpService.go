package elevatorTcpClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"learn101/elevator"
	"log"
	"net"
	"strings"
)

var els elevator.Elevators

type ReqEl struct {
	//鉴别更新还是获取的操作由operation识别
	Operation string //传值0为更新 1为请求电梯
	ReqEle
}

type ReqEle struct {
	TaskId    string //taskId
	Start     int64  //strat
	End       int64
	EleId     string //eleId
	State     string
	Floor     int64
	IsInFloor bool //电梯内是否有机器人 true 为在
}

type Res struct {
	Result bool
	EleId  string
	Error  error
}

func NewTcpService() {
	//1.建立监听端口
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	fmt.Println("listen Start...:")
	els = NewTestEls()
	fmt.Println("初始化电梯数据...")
	for {
		//2.接收客户端的链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		//3.开启一个Goroutine，处理链接
		//go connHandler(conn)
		go coonSt(conn)
	}
}

func coonSt(c net.Conn) {
	//1.conn是否有效
	if c == nil {
		log.Panic("conn无效")
	}
	for {
		//encode 写入数据，decode读取数据
		d := json.NewDecoder(c)
		e := json.NewEncoder(c)
		var msg ReqEl
		_ = d.Decode(&msg)

		switch msg.Operation {
		case "0":
			//更新电梯信息
			midel := &elevator.Elevator{
				ElevatorId:   msg.EleId,
				Floor:        msg.Floor,
				State:        msg.State,
				CurrentState: "0",
			}
			fmt.Println("-------------------------")
			for _, v := range els {
				fmt.Print(v)
			}
			els.Update(midel)
			fmt.Printf("电梯%s被状态被更新了\n", midel.ElevatorId)
			for _, v := range els {
				fmt.Print(v)
			}
			fmt.Println("\n-------------------------")
			resEl := &Res{
				Result: true,
				EleId:  msg.EleId,
				Error:  nil,
			}
			//回传的信息
			e.Encode(&resEl)
		case "1":
			//获取可用电梯
			el, err := els.RightElevator(msg.Start)
			if err != nil {
				resEl := &Res{
					Result: false,
					EleId:  "",
					Error:  errors.New("当前无可用电梯"),
				}
				//传回错误信息
				e.Encode(&resEl)
				fmt.Println("当前无电梯可用")
				for _, v := range els {
					fmt.Print(v)
				}
				fmt.Println()
				continue
			}

			fmt.Printf("本次选用的是电梯%s,taskID为%s", el, msg.TaskId)
			for _, v := range els {
				fmt.Print(v)
			}
			fmt.Println()

			resEl := &Res{
				Result: true,
				EleId:  el,
				Error:  nil,
			}
			e.Encode(&resEl)
		case "2":
			//电梯更新是否抵达起点|终点楼层
			midel := &elevator.Elevator{
				ElevatorId:   msg.EleId,
				Floor:        msg.Floor,
				State:        msg.State,
				CurrentState: "1",
			}
			fmt.Println("-------------------------")
			for _, v := range els {
				fmt.Print(v)
			}
			els.Update(midel)
			fmt.Printf("电梯%s被状态被更新了\n", midel.ElevatorId)
			for _, v := range els {
				fmt.Print(v)
			}
			fmt.Println("\n-------------------------")
			//此时这里检测 机器人是否进|出电梯
			//检测是否进电梯  start楼层等于电梯现在在的楼层，并且isInFloor为true
			if msg.Start == els[msg.EleId].Floor && els[msg.EleId].IsInFloor {
				resEl := &Res{
					Result: true,
					EleId:  msg.EleId,
					Error:  nil,
				}
				//回传的信息 告知电梯可以启动了
				e.Encode(&resEl)
			} else if msg.End == els[msg.EleId].Floor && !els[msg.EleId].IsInFloor {
				//检测是否已经出电梯，end==floor 并且isInFloor为false
				resEl := &Res{
					Result: true,
					EleId:  msg.EleId,
					Error:  nil,
				}
				//回传的信息 告知电梯可以变为空闲状态了，变更为发送0操作
				e.Encode(&resEl)
			}
			resEl := &Res{
				Result: false,
				EleId:  msg.EleId,
				Error:  errors.New("再等等吧"),
			}
			//回传的信息
			e.Encode(&resEl)
		case "3":
			//调度询问电梯是否抵达起点
			//传入的start==目前电梯的floor则表示电梯已经抵达起点
			//接收到返回后 调度把isInfloor的值改为true
			if msg.EleId == els[msg.EleId].ElevatorId && msg.Start == els[msg.EleId].Floor {
				resEl := &Res{
					Result: true,
					EleId:  msg.EleId,
					Error:  nil,
				}
				e.Encode(&resEl)
				return
			}
			resEl := &Res{
				Result: false,
				EleId:  msg.EleId,
				Error:  errors.New("电梯还未抵达起点！"),
			}
			e.Encode(&resEl)
		case "4":
			//调度询问电梯是否抵达终点
			//传入的end==目前电梯的floor则表示电梯已经抵达终点
			//接到true的回复 调度才能让机器人驶出电梯
			if msg.EleId == els[msg.EleId].ElevatorId && msg.End == els[msg.EleId].Floor {
				resEl := &Res{
					Result: true,
					EleId:  msg.EleId,
					Error:  nil,
				}
				e.Encode(&resEl)
				return
			}
			resEl := &Res{
				Result: false,
				EleId:  msg.EleId,
				Error:  errors.New("电梯还未抵达终点！"),
			}
			e.Encode(&resEl)
		case "5":
			//调度发送已经驶入电梯的请求
			//如果接到机器人已驶入电梯，并且楼层正确的，则把电梯信息更新
			if msg.IsInFloor &&msg.Start==els[msg.EleId].Floor{
				els[msg.EleId].IsInFloor=true
				resEl := &Res{
					Result: true,
					EleId:  msg.EleId,
					Error:  nil,
				}
				e.Encode(&resEl)
				return
			}
			resEl := &Res{
				Result: false,
				EleId:  msg.EleId,
				Error:  errors.New("更新机器人驶入状态失败"),
			}
			e.Encode(&resEl)
			case "6":
			//调度发送已经驶出电梯的请求
			//如果接到机器人已驶出电梯，并且楼层正确的，则把电梯信息更新
			if !msg.IsInFloor &&msg.End==els[msg.EleId].Floor{
				els[msg.EleId].IsInFloor=false
				resEl := &Res{
					Result: true,
					EleId:  msg.EleId,
					Error:  nil,
				}
				e.Encode(&resEl)
				return
			}
			resEl := &Res{
				Result: false,
				EleId:  msg.EleId,
				Error:  errors.New("更新机器人驶出状态失败"),
			}
			e.Encode(&resEl)
		default:
			c.Write([]byte("请求参数错误！"))
		}

	}

}

//func NewNullEls()  elevator.Elevators{
//	els := elevator.NewElevators()
//	return *els
//}
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

//func GetRightElevator(w http.ResponseWriter,r *http.Request)  {
//	//els:=NewTestEls()
//	var result struct {
//		eleID string
//	}
//	for _,v:=range els{
//		log.Printf("改变值之前%v",v)
//	}
//	result.eleID,_=els.RightElevator(-1,3)
//	for _,v:=range els{
//		log.Printf("改变值之后%v",v)
//	}
//	resJson,_:=json.Marshal(result)
//	w.Write(resJson)
//}
//处理请求，类型就是net.Conn
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
