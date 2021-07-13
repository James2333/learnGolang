package simlate

import (
	"encoding/json"
	"learn101/elevator"
	"learn101/elevator/packet"
	"learn101/elevator/reply"
	"learn101/elevator/session"
	"log"
	"time"
)

func ParseCodeElevator(code uint16, s *session.Session) {
	log.Printf("收到%s的请求，请求头%d.",s.C.RemoteAddr().String(),code)
	switch code {
	case reply.ELE_TO_START:
		ReqElevatorArriveStart(s)
	case reply.ELE_TO_END:
		ReqElevatorArriveEnd(s)
	case reply.UPDATE_ELE:
		log.Printf("更新电梯信息成功！")
	default:
		reply.ReplyError(s)
	}
}
//更新电梯信息
func ReqUpdateElevator(s *session.Session)  {
	ele:=elevator.Elevator{
		ElevatorId:   "111",
		Floor:        1,
		State:        "0",
	}
	b:=packet.Packet(ele,reply.UPDATE_ELE)
	s.Ch<-b
}
//电梯去到起点楼层
func ReqElevatorArriveStart(s *session.Session){
	q, err := packet.UnPacket(s.C)
	if err != nil {
		log.Println(err)
	}
	var task reply.Task
	_ = json.Unmarshal(q.Content, &task)
	log.Printf("收到任务驶向起点楼层：taskID:%s,ElevatorID:%s,Start:%d,End:%d",task.TaskID,task.ElevatorID,task.Start,task.End)
	log.Println("电梯驶向起点楼层:",task.Start)
	time.Sleep(time.Second*5)
	log.Println("电梯抵达起点楼层:",task.Start)
	b:=packet.Packet(task,reply.ARRIVED_START)
	s.Ch<-b
}
//电梯去到终点楼层
func ReqElevatorArriveEnd(s *session.Session){
	q, err := packet.UnPacket(s.C)
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
	s.Ch<-b
}