package reply

import (
	"encoding/json"
	"errors"
	"learn101/elevator"
	"learn101/elevator/packet"
	"log"
	"net"
	"sync"
)

//初次解包先解开byte[0:2]前两个字节，转换成code，之后switch 这个code 跳转进不同的方法。
//电梯管理需要向调度发送的信息:
//向调度发送电梯已经抵达起点楼层 ROBOT_START=2004
//向调度发送电梯已经抵达终点楼层 ROBOT_END=2005
//电梯管理需要向调度和电梯发送的信息:
//请求电梯到起点楼层 ELE_TO_START=2001
//请求电剃到终点楼层 ELE_TO_END=2002
//向电梯发送任务结束，进入空闲状态 ELE_TO_FREE=2003
//电梯管理需要向调度接受的信息:
//CHOOSE_ELE    = 1002	调度请求最优电梯
//ROBOT_In_Floor     = 1005 机器人是否进电梯里面
//ROBOT_OUT_Floor     = 1006	机器人是否已经出电梯里面
//电梯管理需要向电梯接受的信息:
//UPDATE_ELE    = 1001	更新电梯信息
//ARRIVED_START = 1003	电梯抵达起点楼层
//ARRIVED_END   = 1004	电梯抵达终点楼层
//请求一个电梯到结束任务的流程
//调度请求最优电梯>请求电梯到起点楼层>电梯抵达起点楼层>向调度发送电梯已经抵达起点楼层>机器人是否在电梯里面\n
//>请求电剃到终点楼层>电梯抵达终点楼层>向调度发送电梯已经抵达终点楼层>机器人是否在电梯里面>向电梯发送任务结束，进入空闲状态
const (
	UPDATE_ELE      = 1001
	CHOOSE_ELE      = 1002
	ARRIVED_START   = 1003
	ARRIVED_END     = 1004
	ROBOT_In_Floor  = 1005
	ROBOT_OUT_Floor = 1006
	ERROR           = 1024

	ELE_TO_START = 2001
	ELE_TO_END   = 2002
	ELE_TO_FREE  = 2003
	ROBOT_START  = 2004
	ROBOT_END    = 2005
)

type Message struct {
	bool
	error
}
type Task struct {
	net.Conn
	ElevatorID string `json:"elevator_id"`
	TaskID     string `json:"task_id"`
	Start      int64  `json:"start"`
	End        int64  `json:"end"`
}
type Tasks map[string]*Task

var m sync.Mutex
var tasks Tasks



//更新电梯信息
func ReplyUpdateElevator(c net.Conn, els elevator.Elevators) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var ele elevator.Elevator
	_ = json.Unmarshal(q.Content, &ele)
	ele.Conn = &c
	err=els.Update(&ele)
	if err != nil {
		b:=packet.Packet("",ERROR)
		c.Write(b)
		return
	}
	//msg := Message{
	//	bool:  true,
	//	error: nil,
	//}
	b := packet.Packet("", UPDATE_ELE)
	c.Write(b)
}

//调度请求最优电梯
func ReplyRightElevator(c net.Conn, els elevator.Elevators) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task Task
	_ = json.Unmarshal(q.Content, &task)
	elID, err := els.RightElevator(task.Start)
	if err != nil {
		b:=packet.Packet("",ERROR)
		c.Write(b)
		return
	}
	task.ElevatorID = elID      //把电梯id加到这个task里面去
	els[elID].CurrentState = "1" //电梯变为繁忙状态 这个后面任务结束才能更新成空闲。
	tasks[task.TaskID] = &task   //新增一个任务
	tasks[task.TaskID].Conn = c  //记录调度此次的连接信息，方便之后请求
	ReqElevatorToStart(task.TaskID, els)
	buffer := packet.Packet(task, CHOOSE_ELE)
	c.Write(buffer)
}

//请求电梯到起点楼层
func ReqElevatorToStart(taskid string, els elevator.Elevators) {
	c := *els[tasks[taskid].ElevatorID].Conn
	//把整个task发出去
	buffer := packet.Packet(tasks[taskid], ELE_TO_START)
	c.Write(buffer)
}

//电梯抵达起点楼层
func ReplyElevatorArriveStart(c net.Conn) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task Task
	_ = json.Unmarshal(q.Content, &task)
	task = *tasks[task.TaskID]
	ReqElevatorArriveStart(task.TaskID)
	return
}

//向调度发送电梯已经抵达起点楼层
func ReqElevatorArriveStart(taskid string) {
	c := tasks[taskid].Conn
	//把整个task发出去
	buffer := packet.Packet(tasks[taskid], ROBOT_START)
	c.Write(buffer)
	return
}

//向调度发送电梯已经抵达终点楼层
func ReqElevatorArriveEnd(taskid string) {
	c := tasks[taskid].Conn
	//把整个task发出去
	buffer := packet.Packet(tasks[taskid], ROBOT_END)
	c.Write(buffer)
	return
}

//机器人是否进电梯里面
func ReplyRobotInFloor(c net.Conn, els elevator.Elevators) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task Task
	_ = json.Unmarshal(q.Content, &task)
	task = *tasks[task.TaskID]

	els[task.ElevatorID].IsInFloor = true
	ReqElevatorToEnd(task.TaskID, els) //请求电梯到终点楼层
	return
}

//机器人是否已经出电梯里面
func ReplyRobotOutFloor(c net.Conn, els elevator.Elevators) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task Task
	_ = json.Unmarshal(q.Content, &task)
	task = *tasks[task.TaskID]

	els[task.ElevatorID].IsInFloor = false
	ReqElevatorTaskEnd(task.TaskID, els) //请求电梯到终点楼层
	return
}

//请求电梯到终点楼层
func ReqElevatorToEnd(taskid string, els elevator.Elevators) {
	c := *els[tasks[taskid].ElevatorID].Conn
	//把整个task发出去
	buffer := packet.Packet(tasks[taskid], ELE_TO_END)
	c.Write(buffer)
	return
}

//电梯抵达终点楼层
func ReplyElevatorArriveEnd(c net.Conn) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task Task
	_ = json.Unmarshal(q.Content, &task)
	task = *tasks[task.TaskID]
	ReqElevatorArriveEnd(task.TaskID) //向调度发送电梯已经抵达终点楼层
	return
}

//向电梯发送任务结束，进入空闲状态
func ReqElevatorTaskEnd(taskid string, els elevator.Elevators) {
	//更新电梯信息，防止心跳不及时。
	els[tasks[taskid].ElevatorID].IsInFloor = false
	els[tasks[taskid].ElevatorID].CurrentState = "0"
	els[tasks[taskid].ElevatorID].Floor=tasks[taskid].End
	c := *els[tasks[taskid].ElevatorID].Conn
	m.Lock()
	delete(tasks, taskid) //删除任务
	m.Unlock()
	buffer := packet.Packet("任务结束，进入空闲状态", ELE_TO_FREE)
	c.Write(buffer)
	return
}

//错误返回
func ReplyError(c net.Conn) {
	buffer := packet.Packet(errors.New("404 not found"), ERROR)
	c.Write(buffer)
	return
}