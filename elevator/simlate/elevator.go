package simlate

import (
	"encoding/json"
	"learn101/elevator"
	"learn101/elevator/packet"
	"learn101/elevator/reply"
	"log"
	"net"
	"time"
)

func ParseCodeElevator(code uint16, c net.Conn) {
	switch code {
	case reply.ELE_TO_START:
		ReqElevatorArriveStart(c)
	case reply.ELE_TO_END:
		ReqElevatorArriveEnd(c)
	default:
		reply.ReplyError(c)
	}
}
//更新电梯信息
func ReqUpdateElevator(c net.Conn)  {
	ele:=elevator.Elevator{
		ElevatorId:   "111",
		Floor:        1,
		State:        "0",
	}
	b:=packet.Packet(ele,reply.UPDATE_ELE)
	c.Write(b)
}
//电梯去到起点楼层
func ReqElevatorArriveStart(c net.Conn){
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task reply.Task
	_ = json.Unmarshal(q.Content, &task)
	log.Printf("收到任务驶向起点楼层：taskID:%s,ElevatorID:%s,Start:%d,End:%d",task.TaskID,task.ElevatorID,task.Start,task.End)
	log.Println("电梯驶向起点楼层",task.Start)
	time.Sleep(time.Second*5)
	log.Println("电梯抵达起点楼层",task.Start)
	b:=packet.Packet(task,reply.ARRIVED_START)
	c.Write(b)
}
//电梯去到终点楼层
func ReqElevatorArriveEnd(c net.Conn){
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task reply.Task
	_ = json.Unmarshal(q.Content, &task)
	log.Printf("收到任务驶向终点楼层：taskID:%s,ElevatorID:%s,Start:%d,End:%d",task.TaskID,task.ElevatorID,task.Start,task.End)
	log.Println("电梯驶向终点楼层",task.End)
	time.Sleep(time.Second*5)
	log.Println("电梯抵达终点楼层",task.End)
	b:=packet.Packet(task,reply.ARRIVED_END)
	c.Write(b)
}