package simlate

import (
	"encoding/json"
	"learn101/elevator/packet"
	"learn101/elevator/reply"
	"log"
	"net"
	"time"
)

func ParseCodeScheduling(code uint16, c net.Conn) {
	switch code {
	case reply.ROBOT_START:
		ReqRobotInStart(c)
	case reply.ROBOT_END:
		ReqRobotOutEnd(c)
	default:
		reply.ReplyError(c)
	}
}

func ReqRobotInStart(c net.Conn) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task reply.Task
	_ = json.Unmarshal(q.Content, &task)
	log.Printf("收到任务进入电梯：taskID:%s,ElevatorID:%s,Start:%d,End:%d",task.TaskID,task.ElevatorID,task.Start,task.End)
	time.Sleep(time.Second * 5)
	log.Println("机器人已经进入电梯了")
	b := packet.Packet(task, reply.ROBOT_In_Floor)
	c.Write(b)
}


func ReqRobotOutEnd(c net.Conn) {
	q, err := packet.UnPacket(c)
	if err != nil {
		log.Println(err)
	}
	var task reply.Task
	_ = json.Unmarshal(q.Content, &task)
	log.Printf("收到任务驶出电梯：taskID:%s,ElevatorID:%s,Start:%d,End:%d",task.TaskID,task.ElevatorID,task.Start,task.End)
	time.Sleep(time.Second * 5)
	log.Println("机器人已经出电梯了")
	b := packet.Packet(task, reply.ROBOT_OUT_Floor)
	c.Write(b)
}
