package main

import (
	"context"
	"google.golang.org/grpc"
	pb "learn101/grpcService"
	"log"
)
const (
	// Address 连接地址
	Address string = ":8000"
)
var grpcClient pb.SimpleClient
//func listRouter()  {
//	req:=&pb.SimpleRequest{
//		Data:                 "fuckman",
//	}
//	stream,err:=grpcClient.Router(context.Background(),req)
//	if err != nil {
//
//	}
//
//	for  {
//		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
//		res, err := stream.Recv()
//		// 判断消息流是否已经结束
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			log.Fatalf("ListStr get stream err: %v", err)
//		}
//		// 打印返回值
//		log.Println(res)
//	}
//}
// routeList 调用服务端RouteList方法
func routeList() {
	//调用服务端RouteList方法，获流
	stream, err := grpcClient.Router(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}
	for n := 0; n < 5; n++ {
		//向流中发送消息
		err := stream.Send(&pb.SimpleRequest{Data: "fuckman"+string(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient=pb.NewSimpleClient(conn)
	routeList()
	//req:=&pb.SimpleRequest{
	//	Data:                 "fuckman",
	//}
	//res,err:=grpcClient.Router(context.Background(),req)
	//if err != nil {
	//
	//}
	//log.Println(res)
}
