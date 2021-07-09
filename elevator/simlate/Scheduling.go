package simlate

import (
	"encoding/json"
	"learn101/elevator/packet"
	"learn101/elevator/reply"
	"learn101/elevator/session"
	"log"
	"time"
)

func ParseCodeScheduling(code uint16, s *session.Session) {
	switch code {
	case reply.ROBOT_START:
		ReqRobotInStart(s)
	case reply.ROBOT_END:
		ReqRobotOutEnd(s)
	default:
		reply.ReplyError(s)
	}
}

func ReqRobotInStart(s *session.Session) {
	q, err := packet.UnPacket(s.C)
	if err != nil {
		log.Println(err)
	}
	var task reply.Task
	_ = json.Unmarshal(q.Content, &task)
	log.Printf("收到任务进入电梯：taskID:%s,ElevatorID:%s,Start:%d,End:%d",task.TaskID,task.ElevatorID,task.Start,task.End)
	time.Sleep(time.Second * 5)
	log.Println("机器人已经进入电梯了")
	b := packet.Packet(task, reply.ROBOT_In_Floor)
	s.Ch<-b
}


func ReqRobotOutEnd(s *session.Session) {
	q, err := packet.UnPacket(s.C)
	if err != nil {
		log.Println(err)
	}
	var task reply.Task
	_ = json.Unmarshal(q.Content, &task)
	log.Printf("收到任务驶出电梯：taskID:%s,ElevatorID:%s,Start:%d,End:%d",task.TaskID,task.ElevatorID,task.Start,task.End)
	time.Sleep(time.Second * 5)
	log.Println("机器人已经出电梯了")
	b := packet.Packet(task, reply.ROBOT_OUT_Floor)
	s.Ch<-b
}
