package main

import (
	"google.golang.org/grpc"
	"io"
	pb "learn101/grpcService"
	"log"
	"net"
)

type SimpleService struct {

}
const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)
func (s *SimpleService)Router(srv pb.Simple_RouterServer) (error) {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res)
	}
	//res:=&pb.SimpleResponse{
	//	Code:                 200,
	//	Value:                "hello"+req.Data,
	//}
}


func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	grpcS:=grpc.NewServer()

	pb.RegisterSimpleServer(grpcS,&SimpleService{})
	err = grpcS.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
